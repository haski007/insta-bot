import os
import logging
import requests
import threading
from urllib.parse import unquote
import re
import time
import random
import json
from requests.adapters import HTTPAdapter
from urllib3.util.retry import Retry
from instaloader import Instaloader, Post
from instaloader.exceptions import TwoFactorAuthRequiredException

# Configure logging to emit via Uvicorn's logger so INFO-level lines are visible in container logs
logger = logging.getLogger("uvicorn.error")

# Global Instaloader instance (the "primary"/first account — kept for backwards
# compatibility with callers that only know about a single account, e.g. the
# startup health check in server.py).
_L = None

MIN_SESSIONID_LEN = 10
HOST_SESSION_IMPORT = "/secrets/session.json"
MIN_SESSION_FILE_BYTES = 32

# Multi-account support: account 1 is configured via the plain env vars
# (INSTAGRAM_USERNAME, INSTAGRAM_PASSWORD, ...) for backwards compatibility.
# Additional accounts are configured with a numeric suffix, e.g.
# INSTAGRAM_USERNAME_2 / INSTAGRAM_PASSWORD_2 / INSTAGRAM_SESSIONID_2 /
# INSTAGRAM_PROXY_2 / INSTAGRAM_SESSION_FILE_2, INSTAGRAM_USERNAME_3, ...
# Requests rotate across all accounts with a valid session, and automatically
# fail over to the next account when one is rate-limited/blocked.
MAX_ACCOUNTS = 10


def _account_env(name: str, suffix: str) -> str | None:
    """Read an env var for a given account suffix ("" for account 1, "_2", "_3", ...)."""
    return os.getenv(f"{name}{suffix}")


def _account_label(suffix: str) -> str:
    username = _account_env("INSTAGRAM_USERNAME", suffix)
    return username or f"account{suffix or '#1'}"


def _host_session_import_path(suffix: str) -> str:
    return HOST_SESSION_IMPORT if not suffix else f"/secrets/session{suffix}.json"


def _default_session_save_path(suffix: str) -> str:
    return "/data/session.json" if not suffix else f"/data/session{suffix}.json"


def discover_account_suffixes() -> list[str]:
    """Return account suffixes ("" for account 1, then "_2", "_3", ...) that have config."""
    suffixes = [""]
    for i in range(2, MAX_ACCOUNTS + 1):
        suffix = f"_{i}"
        if _account_env("INSTAGRAM_USERNAME", suffix) or _account_env("INSTAGRAM_SESSIONID", suffix):
            suffixes.append(suffix)
    return suffixes


def resolve_session_paths(suffix: str = "") -> tuple[str | None, str]:
    """Return (path_to_load or None, path_to_save). Missing session triggers login + save."""
    save_path = _account_env("INSTAGRAM_SESSION_FILE", suffix) or _default_session_save_path(suffix)
    host_import = _host_session_import_path(suffix)
    if is_usable_session_file(host_import):
        logger.info(f"Using host session import from {host_import}")
        return host_import, save_path
    if is_usable_session_file(save_path):
        return save_path, save_path
    return None, save_path


def get_sessionid(loader: Instaloader) -> str:
    """Return Instagram sessionid cookie value, or empty string if missing."""
    try:
        return (loader.context._session.cookies.get("sessionid") or "").strip()
    except Exception:
        return ""


def read_sessionid_from_file(session_file: str) -> str:
    """Read sessionid from an Instaloader session pickle on disk."""
    import pickle

    try:
        with open(session_file, "rb") as f:
            data = pickle.load(f)
    except EOFError:
        return ""
    if isinstance(data, dict):
        return (data.get("sessionid") or "").strip()
    return ""


def is_usable_session_file(path: str | None) -> bool:
    """True when file exists, is non-empty, and contains a real sessionid."""
    if not path or not os.path.isfile(path):
        return False
    try:
        if os.path.getsize(path) < MIN_SESSION_FILE_BYTES:
            return False
    except OSError:
        return False
    try:
        return len(read_sessionid_from_file(path)) >= MIN_SESSIONID_LEN
    except Exception:
        return False


def has_valid_session(loader: Instaloader) -> bool:
    """True when the active Instaloader cookies include a real sessionid."""
    return len(get_sessionid(loader)) >= MIN_SESSIONID_LEN


def remove_invalid_session_file(path: str | None) -> None:
    """Delete corrupt or empty session files so login can write a fresh one."""
    if not path or not os.path.isfile(path):
        return
    if is_usable_session_file(path):
        return
    try:
        os.remove(path)
        logger.warning(f"Removed invalid session file {path!r} (will create a new one on login)")
    except OSError as e:
        logger.warning(f"Could not remove invalid session {path!r}: {e}")


def try_load_session(loader: Instaloader, username: str, session_path: str) -> bool:
    """Load session from disk; return True only if cookies have a valid sessionid."""
    if not is_usable_session_file(session_path):
        return False
    try:
        session_dir = os.path.dirname(session_path)
        if session_dir:
            os.makedirs(session_dir, exist_ok=True)
        loader.load_session_from_file(username, session_path)
        sid = get_sessionid(loader)
        logger.info(
            f"Loaded Instagram session from {session_path} "
            f"(logged_in={loader.context.is_logged_in}, sessionid_len={len(sid)})"
        )
        return has_valid_session(loader)
    except Exception as e:
        logger.warning(f"Could not load session from {session_path!r}: {e}")
        return False


def format_login_error(exc: BaseException) -> str:
    """Turn Instaloader/login failures into actionable log messages."""
    msg = str(exc).strip()
    hints: list[str] = []
    lower = msg.lower()
    exc_name = type(exc).__name__

    if "unexpected null login result" in lower:
        hints.append(
            "Instagram rejected password login from this VPS/proxy (very common). "
            "Do not rely on auto-login on the server."
        )
        hints.append(
            "On your PC: log in once, run save_session_to_file, then "
            "copy secrets/session.json to the VPS OR set INSTAGRAM_SESSIONID from browser cookies."
        )
    if "challenge" in lower or "checkpoint" in lower:
        hints.append("Account needs a manual security challenge in the Instagram app/browser first.")
    if "two-factor" in lower or exc_name == "TwoFactorAuthRequiredException":
        hints.append("Set INSTAGRAM_2FA_CODE (one-time SMS/app code) and restart, or import a session file.")
    if "bad credentials" in lower or exc_name == "BadCredentialsException":
        hints.append("Wrong INSTAGRAM_USERNAME or INSTAGRAM_PASSWORD.")
    if "429" in msg or "wait a few minutes" in lower:
        hints.append("Rate-limited — wait 15–30 minutes or switch proxy/IP.")

    if hints:
        return f"{msg} — {' '.join(hints)}"
    return msg


def _perform_password_login(loader: Instaloader, username: str, password: str) -> None:
    """Login with optional 2FA code from INSTAGRAM_2FA_CODE."""
    try:
        loader.login(username, password)
    except TwoFactorAuthRequiredException:
        code = (os.getenv("INSTAGRAM_2FA_CODE") or "").strip()
        if not code:
            raise RuntimeError(
                "Instagram requires 2FA. Set INSTAGRAM_2FA_CODE for this container start, "
                "or import a session file from your browser/local machine."
            ) from None
        logger.info("2FA required — using INSTAGRAM_2FA_CODE")
        loader.two_factor_login(code)


def save_session_file(loader: Instaloader, session_save_path: str) -> None:
    session_dir = os.path.dirname(session_save_path)
    if session_dir:
        os.makedirs(session_dir, exist_ok=True)
    loader.save_session_to_file(session_save_path)
    logger.info(f"Saved session to {session_save_path!r}")


def build_session_data_from_sessionid(sessionid: str, suffix: str = "") -> dict:
    """Instaloader 4.x expects a cookie dict for load_session(), not a plain string."""
    ds_user_id = (_account_env("INSTAGRAM_DS_USER_ID", suffix) or "").strip()
    if not ds_user_id:
        head = sessionid.split(":", 1)[0]
        if head.isdigit():
            ds_user_id = head
    return {
        "sessionid": sessionid,
        "ds_user_id": ds_user_id,
        "csrftoken": (_account_env("INSTAGRAM_CSRFTOKEN", suffix) or "imported").strip(),
        "mid": (_account_env("INSTAGRAM_MID", suffix) or "").strip(),
        "ig_pr": "1",
        "ig_vw": "1920",
        "ig_cb": "1",
        "s_network": "",
        "ig_did": (_account_env("INSTAGRAM_IG_DID", suffix) or "").strip(),
    }


def try_load_sessionid_from_env(
    loader: Instaloader,
    username: str,
    session_save_path: str,
    suffix: str = "",
) -> bool:
    """Load sessionid from INSTAGRAM_SESSIONID env (browser export) and persist to disk."""
    raw = (_account_env("INSTAGRAM_SESSIONID", suffix) or "").strip().strip('"').strip("'")
    if not raw:
        return False
    sessionid = unquote(raw)
    try:
        logger.info("Loading session from INSTAGRAM_SESSIONID env var...")
        loader.load_session(username, build_session_data_from_sessionid(sessionid, suffix))
        if not has_valid_session(loader):
            logger.error("INSTAGRAM_SESSIONID was set but session is still invalid after load_session")
            return False
        save_session_file(loader, session_save_path)
        return True
    except Exception as e:
        logger.error(f"Failed to load INSTAGRAM_SESSIONID: {format_login_error(e)}")
        return False


def login_and_save_session(
    loader: Instaloader,
    username: str,
    password: str,
    session_save_path: str,
) -> bool:
    """Log in with password and persist session file. Returns True on success."""
    remove_invalid_session_file(session_save_path)
    try:
        logger.info(f"No valid session file — logging in as {username} to create {session_save_path!r}...")
        _perform_password_login(loader, username, password)
        sid_len = len(get_sessionid(loader))
        logger.info(f"Login status: logged_in={loader.context.is_logged_in}, sessionid_len={sid_len}")
        if not has_valid_session(loader):
            logger.error("Login finished but sessionid cookie is missing or too short")
            return False
        save_session_file(loader, session_save_path)
        logger.info(f"Created new session file at {session_save_path!r}")
        return True
    except Exception as e:
        logger.error(f"Login failed: {format_login_error(e)}")
        return False


def session_validation_error(
    loader: Instaloader,
    *,
    username: str | None,
    session_file: str | None,
) -> str:
    """Human-readable reason why the current session is not usable."""
    sid = get_sessionid(loader)
    file_sid = ""
    if session_file and os.path.isfile(session_file) and not is_usable_session_file(session_file):
        return (
            f"Session file {session_file!r} is missing, empty, or corrupt. "
            "Set INSTAGRAM_USERNAME + INSTAGRAM_PASSWORD to auto-create it on startup."
        )

    if session_file and is_usable_session_file(session_file):
        try:
            file_sid = read_sessionid_from_file(session_file)
        except Exception as e:
            return (
                f"INSTAGRAM_SESSION_FILE={session_file!r} exists but cannot be read ({e}). "
                "Delete the file and restart, or set INSTAGRAM_PASSWORD for auto re-login."
            )

    if not sid and not file_sid:
        if not session_file or not os.path.isfile(session_file):
            return (
                f"No valid Instagram session for {username!r}. "
                "Password login from a VPS/proxy is often blocked by Instagram "
                '("Unexpected null login result"). Import a session from your PC: '
                "put session.json in secrets/, or set INSTAGRAM_SESSIONID from browser cookies."
            )
        return (
            f"Instagram session for {username!r} has an empty sessionid in {session_file!r}. "
            "Delete the file and import a fresh session from a browser or local machine."
        )
    if len(sid) < MIN_SESSIONID_LEN and len(file_sid) < MIN_SESSIONID_LEN:
        return (
            f"sessionid too short (cookie={len(sid)} chars, file={len(file_sid)} chars). "
            "Session is corrupt or expired — regenerate on the VPS."
        )
    if not sid and file_sid:
        return (
            f"session file contains sessionid but cookies do not — session was not applied. "
            f"Check INSTAGRAM_USERNAME matches the session file ({username!r})."
        )
    return "Unknown session validation failure."


def _create_instaloader_account(suffix: str = "") -> Instaloader:
    """Create and optionally login to Instaloader for one account.

    `suffix` selects which set of `INSTAGRAM_*` env vars to use: "" for the
    primary account (plain names, e.g. INSTAGRAM_USERNAME), or "_2", "_3", ...
    for additional accounts (e.g. INSTAGRAM_USERNAME_2).

    Enhancements:
    - Supports HTTP(S) proxy via `INSTAGRAM_PROXY[_N]` or standard env vars
    - Allows custom User-Agent via `INSTAGRAM_USER_AGENT[_N]`
    - Persists session to file `INSTAGRAM_SESSION_FILE[_N]` (if provided)
    - Adds connection retries for transient 403/5xx
    """
    label = _account_label(suffix)
    logger.info(f"Creating new Instaloader instance for {label!r}...")
    # The iPhone-private-API lookup Instaloader uses for "high-quality" video/image
    # versions requires a genuine mobile-app session; sessionid-only imports (the
    # normal case here) always get 403 "login_required" from i.instagram.com after
    # 3 slow retries. Disabling it avoids ~3-8s of wasted latency per video request
    # while falling back to the perfectly usable GraphQL/web-API video URL.
    iphone_support = (_account_env("INSTAGRAM_IPHONE_SUPPORT", suffix) or "").strip().lower() in (
        "1", "true", "yes", "y", "on",
    )
    if not iphone_support:
        logger.info(f"[{label}] iPhone high-quality media lookup disabled (INSTAGRAM_IPHONE_SUPPORT=false)")
    loader = Instaloader(iphone_support=iphone_support)

    # Configure proxy if provided (falls back to the shared INSTAGRAM_PROXY for
    # extra accounts if no dedicated proxy was set; ideally each account should
    # use its own proxy/IP to avoid correlating them).
    proxy = _account_env("INSTAGRAM_PROXY", suffix) or (os.getenv("INSTAGRAM_PROXY") if suffix else None)
    try:
        if proxy:
            logger.info(f"[{label}] Configuring proxy")
            loader.context._session.proxies = {
                'http': proxy,
                'https': proxy
            }
        else:
            # Allow standard HTTP(S)_PROXY env variables to be used
            loader.context._session.trust_env = True
        logger.info(f"[{label}] Proxy configuration applied")
    except Exception as e:
        logger.error(f"[{label}] Failed to configure proxy: {e}")

    # Configure User-Agent if provided (helps avoid bot detection on DC IPs)
    user_agent = _account_env("INSTAGRAM_USER_AGENT", suffix) or (os.getenv("INSTAGRAM_USER_AGENT") if suffix else None)
    if user_agent:
        try:
            loader.context._session.headers['User-Agent'] = user_agent
            # Common real-browser headers improve legitimacy a bit
            loader.context._session.headers.setdefault('Accept', 'text/html,application/xhtml+xml,application/xml;q=0.9,image/avif,image/webp,*/*;q=0.8')
            loader.context._session.headers.setdefault('Accept-Language', 'en-US,en;q=0.9')
            loader.context._session.headers.setdefault('Upgrade-Insecure-Requests', '1')
            logger.info(f"[{label}] Custom User-Agent header applied")
        except Exception as e:
            logger.error(f"[{label}] Failed to set custom User-Agent: {e}")

    # Add basic retries for transient HTTP errors
    try:
        retry = Retry(
            total=3,
            backoff_factor=1.0,
            status_forcelist=[403, 429, 500, 502, 503, 504],
            allowed_methods=["HEAD", "GET", "OPTIONS"]
        )
        adapter = HTTPAdapter(max_retries=retry)
        loader.context._session.mount('http://', adapter)
        loader.context._session.mount('https://', adapter)
        logger.info(f"[{label}] HTTP retries configured")
    except Exception as e:
        logger.error(f"[{label}] Failed to configure HTTP retries: {e}")

    # Get credentials and session persistence settings from environment variables
    username = _account_env("INSTAGRAM_USERNAME", suffix)
    password = _account_env("INSTAGRAM_PASSWORD", suffix)
    session_load_path, session_save_path = resolve_session_paths(suffix)
    session_dir = os.path.dirname(session_save_path)
    if session_dir:
        os.makedirs(session_dir, exist_ok=True)

    if username and session_load_path:
        if not try_load_session(loader, username, session_load_path):
            logger.info(f"[{label}] Existing session file not loaded or invalid")
    elif username:
        logger.info(
            f"[{label}] No session file yet at {session_save_path!r} "
            f"(optional import: {_host_session_import_path(suffix)})"
        )

    if has_valid_session(loader):
        logger.info(f"[{label}] Valid sessionid present; skipping login")
    elif username and try_load_sessionid_from_env(loader, username, session_save_path, suffix):
        logger.info(f"[{label}] Session loaded from INSTAGRAM_SESSIONID{suffix}")
    elif username and password:
        login_and_save_session(loader, username, password, session_save_path)
    else:
        logger.warning(
            f"[{label}] No valid session. Set INSTAGRAM_PASSWORD{suffix} (often blocked on VPS), "
            f"INSTAGRAM_SESSIONID{suffix}, or {_host_session_import_path(suffix)} from your local machine."
        )

    sid_len = len(get_sessionid(loader))
    valid = has_valid_session(loader)
    logger.info(
        f"[{label}] Instaloader ready; logged_in={loader.context.is_logged_in}, "
        f"valid_session={valid}, sessionid_len={sid_len}"
    )
    return loader


def create_instaloader():
    """Create (or return the cached) primary Instaloader instance (account 1).

    Kept for backwards compatibility with callers that only know about a
    single account (e.g. the startup health check). For multi-account
    request handling, use `get_account_pool()` instead.
    """
    global _L

    if _L is not None:
        logger.info("Instaloader already initialized, returning existing instance")
        return _L

    _L = _create_instaloader_account("")
    return _L


def get_instaloader():
    """Get the global (primary) Instaloader instance, creating it if necessary"""
    global _L
    if _L is None:
        _L = create_instaloader()
    return _L


# Cache of additional accounts (suffix -> Instaloader), lazily built.
_EXTRA_ACCOUNTS: dict[str, Instaloader] = {}
# Every configured account that was successfully built, valid or not — used
# for startup diagnostics (see describe_accounts()).
_ALL_ACCOUNTS: list[tuple[str, Instaloader]] = []
# Cache of the final rotation pool, built once from all configured accounts.
_ACCOUNT_POOL: list[Instaloader] | None = None
_rotation_lock = threading.Lock()
_rotation_counter = 0


def get_account_pool() -> list[Instaloader]:
    """Return all configured Instagram accounts usable for requests.

    Builds the primary account plus any INSTAGRAM_USERNAME_2/_3/... accounts on
    first use. Prefers accounts with a valid session; if none are valid (e.g.
    anonymous/dev mode), falls back to returning all configured accounts so
    unauthenticated fetching still works as before.
    """
    global _ACCOUNT_POOL
    if _ACCOUNT_POOL is not None:
        return _ACCOUNT_POOL

    if not _ALL_ACCOUNTS:
        _ALL_ACCOUNTS.append(("", get_instaloader()))
        for suffix in discover_account_suffixes()[1:]:
            loader = _EXTRA_ACCOUNTS.get(suffix)
            if loader is None:
                try:
                    loader = _create_instaloader_account(suffix)
                    _EXTRA_ACCOUNTS[suffix] = loader
                except Exception as e:
                    logger.error(f"Failed to initialize Instagram account {_account_label(suffix)!r}: {e}")
                    continue
            _ALL_ACCOUNTS.append((suffix, loader))

    accounts = [loader for _, loader in _ALL_ACCOUNTS]
    valid = [a for a in accounts if has_valid_session(a)]
    if valid:
        if len(valid) < len(accounts):
            logger.warning(
                f"{len(accounts) - len(valid)} of {len(accounts)} configured Instagram "
                "account(s) have no valid session and will be excluded from rotation"
            )
        _ACCOUNT_POOL = valid
    else:
        _ACCOUNT_POOL = accounts

    if len(_ACCOUNT_POOL) > 1:
        logger.info(f"Instagram account pool ready: {len(_ACCOUNT_POOL)} accounts in rotation")
    return _ACCOUNT_POOL


def describe_accounts() -> list[dict]:
    """Diagnostic summary of every configured account (valid or not).

    Builds the account pool if not already built. Intended for startup health
    checks / logging, not for the hot request path.
    """
    get_account_pool()
    return [
        {
            "label": _account_label(suffix),
            "valid": has_valid_session(loader),
            "sessionid_len": len(get_sessionid(loader)),
            "logged_in": bool(loader.context.is_logged_in),
        }
        for suffix, loader in _ALL_ACCOUNTS
    ]


def _rotated_accounts() -> list[Instaloader]:
    """Return the account pool reordered so consecutive calls start with a
    different account (round-robin), with the rest kept as failover order."""
    pool = get_account_pool()
    if len(pool) <= 1:
        return pool
    global _rotation_counter
    with _rotation_lock:
        start = _rotation_counter % len(pool)
        _rotation_counter += 1
    return pool[start:] + pool[:start]

# Public Instagram web app id; required by the private web API endpoints.
IG_WEB_APP_ID = "936619743392459"

# Shortcode alphabet used by Instagram (URL-safe base64).
_SHORTCODE_ALPHABET = (
    "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789-_"
)


def shortcode_to_mediaid(shortcode: str) -> int:
    """Decode an Instagram shortcode (URL-safe base64) into its numeric media id."""
    mediaid = 0
    for char in shortcode:
        mediaid = mediaid * 64 + _SHORTCODE_ALPHABET.index(char)
    return mediaid


def _has_media(info: dict) -> bool:
    """True when the result dict carries a usable image or video URL."""
    media = info.get("media") or []
    if media:
        return any(bool(m.get("video_url")) or bool(m.get("url")) for m in media)
    return bool(info.get("video_url")) or bool(info.get("url"))


def _media_entry_from_web_item(item: dict) -> dict | None:
    """Build a single media entry from a web-API media item (image/video slide)."""
    video_versions = item.get("video_versions") or []
    video_url = video_versions[0]["url"] if video_versions else ""

    image_url = ""
    candidates = (item.get("image_versions2") or {}).get("candidates") or []
    if candidates:
        image_url = candidates[0]["url"]

    if not video_url and not image_url:
        return None

    is_video = bool(video_url)
    return {
        "is_video": is_video,
        "url": image_url or "",
        "video_url": video_url if is_video else None,
    }


def _build_post_result(
    shortcode: str,
    media_items: list[dict],
    *,
    caption: str = "",
    owner: str = "",
    likes: int = 0,
    comments: int = 0,
    timestamp: str = "",
) -> dict:
    """Normalize post payload; top-level url fields mirror the first media item."""
    first = media_items[0] if media_items else {
        "is_video": False,
        "url": "",
        "video_url": None,
    }
    return {
        "shortcode": shortcode,
        "is_video": bool(first.get("is_video")),
        "url": first.get("url") or "",
        "video_url": first.get("video_url"),
        "caption": caption or "",
        "owner": owner or "",
        "likes": likes or 0,
        "comments": comments or 0,
        "timestamp": timestamp or "",
        "media": media_items,
    }


def _parse_web_api_item(item: dict, *, label: str) -> dict:
    """Turn a web-API media item into our normalized post/story payload."""
    # media_type: 1=image, 2=video, 8=carousel (sidecar)
    media_items: list[dict] = []
    if item.get("media_type") == 8:
        for slide in item.get("carousel_media") or []:
            entry = _media_entry_from_web_item(slide)
            if entry:
                media_items.append(entry)
        logger.info(f"Web API carousel for {label}: {len(media_items)} slides")
    else:
        entry = _media_entry_from_web_item(item)
        if entry:
            media_items.append(entry)

    if not media_items:
        raise RuntimeError("web api returned no usable media urls")

    # Caption/owner live on the parent item for carousels, not on each slide.
    caption_node = item.get("caption") or {}
    caption = caption_node.get("text", "") if isinstance(caption_node, dict) else ""
    owner = (item.get("user") or {}).get("username", "")

    taken_at = item.get("taken_at")
    timestamp = ""
    if isinstance(taken_at, (int, float)):
        from datetime import datetime, timezone

        timestamp = datetime.fromtimestamp(taken_at, tz=timezone.utc).isoformat()

    code = item.get("code") or label
    return _build_post_result(
        str(code),
        media_items,
        caption=caption,
        owner=owner,
        likes=item.get("like_count", 0) or 0,
        comments=item.get("comment_count", 0) or 0,
        timestamp=timestamp,
    )


def _fetch_via_web_api_mediaid(
    mediaid: int | str,
    session: requests.Session,
    *,
    referer: str,
    label: str,
) -> dict:
    """Fetch media (post or story) by numeric id via Instagram's private web API."""
    api_url = f"https://www.instagram.com/api/v1/media/{mediaid}/info/"

    headers = dict(session.headers)
    headers["X-IG-App-ID"] = IG_WEB_APP_ID
    headers.setdefault("Accept", "*/*")
    headers.setdefault("Referer", referer)

    # Don't follow redirects: an unauthenticated/rejected session on this endpoint
    # causes Instagram to bounce between login/challenge redirects until requests'
    # 30-redirect cap trips (~5s wasted per call). A bare redirect here reliably
    # means the session isn't accepted for this endpoint, so fail fast instead.
    resp = session.get(api_url, headers=headers, timeout=20, allow_redirects=False)
    logger.info(f"Web API GET {api_url} -> {resp.status_code}")
    if resp.is_redirect or resp.status_code in (301, 302, 303, 307, 308):
        location = resp.headers.get("Location", "")
        raise RuntimeError(f"web api redirected (likely login_required) -> {location!r}")
    if resp.status_code != 200:
        raise RuntimeError(f"web api http {resp.status_code}")

    data = resp.json()
    items = data.get("items") or []
    if not items:
        raise RuntimeError("web api returned no items")
    return _parse_web_api_item(items[0], label=label)


def _fetch_via_web_api(shortcode: str, session: requests.Session) -> dict:
    """Fetch post/reel media info via Instagram's private web API."""
    mediaid = shortcode_to_mediaid(shortcode)
    return _fetch_via_web_api_mediaid(
        mediaid,
        session,
        referer=f"https://www.instagram.com/p/{shortcode}/",
        label=shortcode,
    )


def _fetch_story_via_instaloader(loader: Instaloader, media_id: int) -> dict:
    """Fallback: Instaloader StoryItem by media id (requires login)."""
    from instaloader import StoryItem

    item = StoryItem.from_mediaid(loader.context, media_id)
    media_items = [{
        "is_video": item.is_video,
        "url": item.url or "",
        "video_url": item.video_url if item.is_video else None,
    }]
    owner = ""
    try:
        owner = item.owner_username or ""
    except Exception:
        pass
    timestamp = ""
    try:
        if item.date_utc:
            timestamp = item.date_utc.isoformat()
    except Exception:
        pass
    return _build_post_result(
        str(media_id),
        media_items,
        owner=owner,
        timestamp=timestamp,
    )


def _rate_limited_error(msg: str) -> bool:
    """True when an error looks like a per-account block (worth trying another account)."""
    return any(tok in msg for tok in ("429", "401", "403", "login_required", "wait a few minutes"))


def _get_story_info_with_loader(loader: Instaloader, media_id: str, mediaid_int: int, username: str) -> dict:
    logger.info(
        f"Fetching story media_id={media_id} username={username!r} "
        f"logged_in={loader.context.is_logged_in}"
    )
    session = loader.context._session

    if username:
        referer = f"https://www.instagram.com/stories/{username}/{media_id}/"
    else:
        referer = "https://www.instagram.com/"

    attempts: list[str] = []

    try:
        result = _fetch_via_web_api_mediaid(
            mediaid_int,
            session,
            referer=referer,
            label=media_id,
        )
        if username and not result.get("owner"):
            result["owner"] = username
        if _has_media(result):
            return result
        logger.warning(f"Web API returned no media for story {media_id}")
    except Exception as e:
        attempts.append(str(e))
        logger.error(f"Web API failed for story {media_id}: {attempts[-1]}")

    try:
        result = _fetch_story_via_instaloader(loader, mediaid_int)
        if username and not result.get("owner"):
            result["owner"] = username
        if _has_media(result):
            return result
        logger.warning(f"Instaloader StoryItem returned no media for {media_id}")
    except Exception as e:
        attempts.append(str(e))
        logger.error(f"Instaloader StoryItem failed for {media_id}: {attempts[-1]}")

    return {
        "error": (attempts[-1] if attempts else f"no media URL available for story {media_id}"),
        "_attempts": attempts,
    }


def get_story_info(media_id: str, username: str = "") -> dict:
    """Fetch a single Instagram story by numeric media id.

    Story URLs look like: https://www.instagram.com/stories/{username}/{media_id}/
    Requires a logged-in session that can view the story (public, or following
    private accounts). Stories expire after ~24h.

    When multiple Instagram accounts are configured (INSTAGRAM_USERNAME_2, ...),
    rotates across them and automatically retries with another account when one
    looks rate-limited or logged out.
    """
    media_id = (media_id or "").strip()
    if not media_id.isdigit():
        return {"error": f"invalid story media_id: {media_id!r}"}

    mediaid_int = int(media_id)
    username = (username or "").strip().lstrip("@")

    pool = _rotated_accounts()
    last_err = ""
    for loader in pool:
        result = _get_story_info_with_loader(loader, media_id, mediaid_int, username)
        if not result.get("error"):
            return result
        last_err = result["error"]
        attempts = result.pop("_attempts", None) or [last_err]
        if len(pool) > 1 and any(_rate_limited_error(a) for a in attempts):
            logger.warning(f"Story {media_id} looks rate-limited on this account; trying next account")
            continue
        break

    return {"error": last_err or f"no media URL available for story {media_id}"}


def _fetch_public_page_metadata(shortcode: str, session: requests.Session) -> dict:
    """Fetch media info by scraping public page HTML (OG tags and JSON-LD)."""
    headers = dict(session.headers)
    headers.setdefault('User-Agent', os.getenv('INSTAGRAM_USER_AGENT', headers.get('User-Agent', 'Mozilla/5.0')))
    headers.setdefault('Accept', 'text/html,application/xhtml+xml,application/xml;q=0.9,image/avif,image/webp,*/*;q=0.8')
    headers.setdefault('Accept-Language', 'en-US,en;q=0.9')
    headers.setdefault('Referer', 'https://www.instagram.com/')

    candidate_urls = [
        f"https://www.instagram.com/p/{shortcode}/",
        f"https://www.instagram.com/reel/{shortcode}/",
    ]

    html = None
    for u in candidate_urls:
        try:
            resp = session.get(u, headers=headers, timeout=20)
            logger.info(f"Fallback GET {u} -> {resp.status_code}")
            if resp.status_code == 200:
                html = resp.text
                break
        except Exception as e2:
            logger.warning(f"Fallback fetch failed for {u}: {e2}")

    if not html:
        raise RuntimeError("fallback html not available")

    # Extract basic OG meta
    def meta_content(prop: str) -> str:
        m = re.search(rf'<meta[^>]+property=["\']{re.escape(prop)}["\'][^>]+content=["\']([^"\']+)["\']', html, re.IGNORECASE)
        return m.group(1) if m else ""

    og_video = meta_content('og:video') or meta_content('og:video:secure_url')
    og_image = meta_content('og:image')
    og_title = meta_content('og:title')
    og_desc = meta_content('og:description')
    author = meta_content('instagram:owner_user_name') or meta_content('og:site_name')

    is_video = bool(og_video)
    video_url = og_video
    image_url = og_image
    caption = og_title or og_desc

    # JSON-LD enhancement (often contains contentUrl/thumbnailUrl)
    try:
        for m in re.finditer(r'<script[^>]+type=["\']application/ld\+json["\'][^>]*>(.*?)</script>', html, re.IGNORECASE | re.DOTALL):
            data = json.loads(m.group(1).strip())
            if isinstance(data, dict):
                content_url = data.get('contentUrl') or data.get('video', {}).get('contentUrl') if isinstance(data.get('video'), dict) else None
                thumbnail_url = data.get('thumbnailUrl')
                if content_url and not video_url:
                    video_url = content_url
                    is_video = True
                if thumbnail_url and not image_url:
                    if isinstance(thumbnail_url, list) and thumbnail_url:
                        image_url = thumbnail_url[0]
                    elif isinstance(thumbnail_url, str):
                        image_url = thumbnail_url
                if not caption:
                    caption = data.get('description') or caption
            elif isinstance(data, list):
                for item in data:
                    if not isinstance(item, dict):
                        continue
                    content_url = item.get('contentUrl')
                    if content_url and not video_url:
                        video_url = content_url
                        is_video = True
                    thumbnail_url = item.get('thumbnailUrl')
                    if thumbnail_url and not image_url:
                        if isinstance(thumbnail_url, list) and thumbnail_url:
                            image_url = thumbnail_url[0]
                        elif isinstance(thumbnail_url, str):
                            image_url = thumbnail_url
                    if not caption:
                        caption = item.get('description') or caption
    except Exception as e:
        logger.warning(f"JSON-LD parse failed: {e}")

    # Inline JSON fallback: extract video_url/display_url/thumbnail_src from embedded JSON
    try:
        def _unescape(val: str) -> str:
            try:
                return json.loads(f'"{val}"')
            except Exception:
                return val.replace('\\u0026', '&').replace('\\/', '/')

        if not video_url:
            m = re.search(r'"video_url":"([^"\\]+(?:\\.[^"\\]+)*)"', html)
            if m:
                video_url = _unescape(m.group(1))
                is_video = True
        if not image_url:
            m = re.search(r'"display_url":"([^"\\]+(?:\\.[^"\\]+)*)"', html)
            if m:
                image_url = _unescape(m.group(1))
        if not image_url:
            m = re.search(r'"thumbnail_src":"([^"\\]+(?:\\.[^"\\]+)*)"', html)
            if m:
                image_url = _unescape(m.group(1))
        # Additional patterns frequently present in reels
        if not video_url:
            # video_versions: [{"url":"...mp4"}]
            m = re.search(r'"video_versions"\s*:\s*\[\s*\{[^\}]*?"url"\s*:\s*"([^"\\]+(?:\\.[^"\\]+)*)"', html, re.DOTALL)
            if m:
                video_url = _unescape(m.group(1))
                is_video = True
        if not video_url:
            # playback_url may point to m3u8; prefer mp4 if found later
            m = re.search(r'"playback_url"\s*:\s*"([^"\\]+)"', html)
            if m:
                candidate = _unescape(m.group(1))
                # only accept if it looks like mp4; otherwise keep as last resort
                if candidate.endswith('.mp4'):
                    video_url = candidate
                    is_video = True
        if not video_url:
            # Generic .mp4 in HTML (e.g., data URLs or sources)
            m = re.search(r'https?://[^"\s<>]+\.mp4', html)
            if m:
                video_url = _unescape(m.group(0))
                is_video = True
        if not author:
            m = re.search(r'"username":"([^"]+)"', html)
            if m:
                author = _unescape(m.group(1))
        if not caption:
            m = re.search(r'"edge_media_to_caption"\s*:\s*\{\s*"edges"\s*:\s*\[\{\s*"node"\s*:\s*\{\s*"text"\s*:\s*"(.*?)"', html, re.DOTALL)
            if m:
                caption = _unescape(m.group(1))
    except Exception as e:
        logger.warning(f"Inline JSON extraction failed: {e}")

    # oEmbed fallback for thumbnail/author when still missing
    try:
        if not (video_url or image_url):
            for base in ("https://www.instagram.com/p/", "https://www.instagram.com/reel/"):
                oembed_url = f"https://www.instagram.com/oembed/?url={base}{shortcode}/"
                r = session.get(oembed_url, headers=headers, timeout=15)
                logger.info(f"oEmbed GET {oembed_url} -> {r.status_code}")
                if r.status_code == 200:
                    data = r.json()
                    image_url = image_url or data.get("thumbnail_url") or ""
                    author = author or data.get("author_name") or author
                    caption = caption or data.get("title") or caption
                    if image_url:
                        break
    except Exception as e:
        logger.warning(f"oEmbed fallback failed: {e}")

    media_items = []
    if video_url or image_url:
        media_items.append({
            "is_video": is_video,
            "url": image_url or "",
            "video_url": video_url if is_video else None,
        })

    return _build_post_result(
        shortcode,
        media_items,
        caption=caption or "",
        owner=author or "",
    )


def _fetch_via_graphql(loader: Instaloader, shortcode: str) -> dict:
    """Fetch post via Instaloader GraphQL, including all sidecar slides."""
    post = Post.from_shortcode(loader.context, shortcode)
    media_items: list[dict] = []

    if post.typename == "GraphSidecar":
        for node in post.get_sidecar_nodes():
            media_items.append({
                "is_video": node.is_video,
                "url": node.display_url or "",
                "video_url": node.video_url if node.is_video else None,
            })
        logger.info(f"GraphQL carousel for {shortcode}: {len(media_items)} slides")
    else:
        media_items.append({
            "is_video": post.is_video,
            "url": post.url or "",
            "video_url": post.video_url if post.is_video else None,
        })

    return _build_post_result(
        shortcode,
        media_items,
        caption=post.caption or "",
        owner=post.owner_username or "",
        likes=post.likes or 0,
        comments=post.comments or 0,
        timestamp=post.date_utc.isoformat() if post.date_utc else "",
    )


def _get_post_info_with_loader(loader: Instaloader, shortcode: str) -> dict:
    logger.info(f"Instaloader logged_in={loader.context.is_logged_in}")
    session = loader.context._session
    attempts: list[str] = []

    # Optional: avoid the API/GraphQL entirely when forced
    if os.getenv('INSTAGRAM_FORCE_HTML_FALLBACK', 'false').lower() in ('1', 'true', 'yes'):
        logger.info("Force HTML fallback enabled; skipping web API and GraphQL")
        try:
            result = _fetch_public_page_metadata(shortcode, session)
            if _has_media(result):
                return result
            return {"error": f"no media URL available for {shortcode}", "_attempts": attempts}
        except Exception as e:
            logger.error(f"HTML fallback failed for {shortcode}: {e}")
            return {"error": str(e), "_attempts": [str(e)]}

    # 1) Private web API using the logged-in session — most reliable for reels.
    try:
        result = _fetch_via_web_api(shortcode, session)
        if _has_media(result):
            return result
        logger.warning(f"Web API returned no media for {shortcode}")
    except Exception as e:
        attempts.append(str(e))
        logger.error(f"Web API failed for {shortcode}: {attempts[-1]}")

    # 2) GraphQL via Instaloader (supports full carousels via sidecar nodes).
    try:
        result = _fetch_via_graphql(loader, shortcode)
        if _has_media(result):
            return result
        logger.warning(f"GraphQL returned no media for {shortcode}")
    except Exception as e:
        attempts.append(str(e))
        logger.error(f"GraphQL failed for {shortcode}: {attempts[-1]}")

    # 3) Public page HTML scraping (OG tags / JSON-LD / inline JSON).
    try:
        # Gentle backoff before the last-resort scrape
        time.sleep(random.uniform(0.8, 1.6))
        result = _fetch_public_page_metadata(shortcode, session)
        if _has_media(result):
            return result
        logger.warning(f"HTML fallback returned no media for {shortcode}")
    except Exception as e2:
        attempts.append(str(e2))
        logger.error(f"Fallback scrape failed for {shortcode}: {attempts[-1]}")

    return {
        "error": (attempts[-1] if attempts else f"no media URL available for {shortcode}"),
        "_attempts": attempts,
    }


def get_post_info(shortcode: str):
    """Fetch a post/reel's media info, rotating across configured Instagram
    accounts (INSTAGRAM_USERNAME_2, ...) and failing over to the next account
    when one looks rate-limited/logged-out."""
    pool = _rotated_accounts()
    last_err = ""
    for loader in pool:
        result = _get_post_info_with_loader(loader, shortcode)
        if not result.get("error"):
            return result
        last_err = result["error"]
        attempts = result.pop("_attempts", None) or [last_err]
        if len(pool) > 1 and any(_rate_limited_error(a) for a in attempts):
            logger.warning(f"{shortcode} looks rate-limited on this account; trying next account")
            continue
        break

    return {"error": last_err or f"no media URL available for {shortcode}"}