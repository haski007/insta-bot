import os
import logging
import requests
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

# Global Instaloader instance
_L = None

MIN_SESSIONID_LEN = 10
HOST_SESSION_IMPORT = "/secrets/session.json"
MIN_SESSION_FILE_BYTES = 32


def resolve_session_paths() -> tuple[str | None, str]:
    """Return (path_to_load or None, path_to_save). Missing session triggers login + save."""
    save_path = os.getenv("INSTAGRAM_SESSION_FILE") or "/data/session.json"
    if is_usable_session_file(HOST_SESSION_IMPORT):
        logger.info(f"Using host session import from {HOST_SESSION_IMPORT}")
        return HOST_SESSION_IMPORT, save_path
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


def build_session_data_from_sessionid(sessionid: str) -> dict:
    """Instaloader 4.x expects a cookie dict for load_session(), not a plain string."""
    ds_user_id = (os.getenv("INSTAGRAM_DS_USER_ID") or "").strip()
    if not ds_user_id:
        head = sessionid.split(":", 1)[0]
        if head.isdigit():
            ds_user_id = head
    return {
        "sessionid": sessionid,
        "ds_user_id": ds_user_id,
        "csrftoken": (os.getenv("INSTAGRAM_CSRFTOKEN") or "imported").strip(),
        "mid": (os.getenv("INSTAGRAM_MID") or "").strip(),
        "ig_pr": "1",
        "ig_vw": "1920",
        "ig_cb": "1",
        "s_network": "",
        "ig_did": (os.getenv("INSTAGRAM_IG_DID") or "").strip(),
    }


def try_load_sessionid_from_env(
    loader: Instaloader,
    username: str,
    session_save_path: str,
) -> bool:
    """Load sessionid from INSTAGRAM_SESSIONID env (browser export) and persist to disk."""
    raw = (os.getenv("INSTAGRAM_SESSIONID") or "").strip().strip('"').strip("'")
    if not raw:
        return False
    sessionid = unquote(raw)
    try:
        logger.info("Loading session from INSTAGRAM_SESSIONID env var...")
        loader.load_session(username, build_session_data_from_sessionid(sessionid))
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


def create_instaloader():
    """Create and optionally login to Instaloader.

    Enhancements:
    - Supports HTTP(S) proxy via `INSTAGRAM_PROXY` or standard env vars
    - Allows custom User-Agent via `INSTAGRAM_USER_AGENT`
    - Persists session to file `INSTAGRAM_SESSION_FILE` (if provided)
    - Adds connection retries for transient 403/5xx
    """
    global _L
    
    if _L is not None:
        logger.info("Instaloader already initialized, returning existing instance")
        return _L
    
    logger.info("Creating new Instaloader instance...")
    _L = Instaloader()
    
    # Configure proxy if provided
    proxy = os.getenv('INSTAGRAM_PROXY')
    try:
        if proxy:
            logger.info(f"Configuring proxy (INSTAGRAM_PROXY)")
            _L.context._session.proxies = {
                'http': proxy,
                'https': proxy
            }
        else:
            # Allow standard HTTP(S)_PROXY env variables to be used
            _L.context._session.trust_env = True
        logger.info("Proxy configuration applied")
    except Exception as e:
        logger.error(f"Failed to configure proxy: {e}")

    # Configure User-Agent if provided (helps avoid bot detection on DC IPs)
    user_agent = os.getenv('INSTAGRAM_USER_AGENT')
    if user_agent:
        try:
            _L.context._session.headers['User-Agent'] = user_agent
            # Common real-browser headers improve legitimacy a bit
            _L.context._session.headers.setdefault('Accept', 'text/html,application/xhtml+xml,application/xml;q=0.9,image/avif,image/webp,*/*;q=0.8')
            _L.context._session.headers.setdefault('Accept-Language', 'en-US,en;q=0.9')
            _L.context._session.headers.setdefault('Upgrade-Insecure-Requests', '1')
            logger.info("Custom User-Agent header applied")
        except Exception as e:
            logger.error(f"Failed to set custom User-Agent: {e}")

    # Add basic retries for transient HTTP errors
    try:
        retry = Retry(
            total=3,
            backoff_factor=1.0,
            status_forcelist=[403, 429, 500, 502, 503, 504],
            allowed_methods=["HEAD", "GET", "OPTIONS"]
        )
        adapter = HTTPAdapter(max_retries=retry)
        _L.context._session.mount('http://', adapter)
        _L.context._session.mount('https://', adapter)
        logger.info("HTTP retries configured")
    except Exception as e:
        logger.error(f"Failed to configure HTTP retries: {e}")
    
    # Get credentials and session persistence settings from environment variables
    username = os.getenv('INSTAGRAM_USERNAME')
    password = os.getenv('INSTAGRAM_PASSWORD')
    session_load_path, session_save_path = resolve_session_paths()
    session_dir = os.path.dirname(session_save_path)
    if session_dir:
        os.makedirs(session_dir, exist_ok=True)

    if username and session_load_path:
        if not try_load_session(_L, username, session_load_path):
            logger.info("Existing session file not loaded or invalid")
    elif username:
        logger.info(
            f"No session file yet at {session_save_path!r} "
            f"(optional import: {HOST_SESSION_IMPORT})"
        )

    if has_valid_session(_L):
        logger.info("Valid sessionid present; skipping login")
    elif username and try_load_sessionid_from_env(_L, username, session_save_path):
        logger.info("Session loaded from INSTAGRAM_SESSIONID")
    elif username and password:
        login_and_save_session(_L, username, password, session_save_path)
    else:
        logger.warning(
            "No valid session. Set INSTAGRAM_PASSWORD (often blocked on VPS), "
            "INSTAGRAM_SESSIONID, or secrets/session.json from your local machine."
        )

    sid_len = len(get_sessionid(_L))
    valid = has_valid_session(_L)
    logger.info(
        f"Instaloader ready; logged_in={_L.context.is_logged_in}, "
        f"valid_session={valid}, sessionid_len={sid_len}"
    )
    return _L

def get_instaloader():
    """Get the global Instaloader instance, creating it if necessary"""
    global _L
    if _L is None:
        logger.info("Creating new Instaloader instance")
        _L = create_instaloader()
    return _L

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

    return {
        "shortcode": shortcode,
        "is_video": is_video,
        "url": image_url or "",
        "video_url": video_url if is_video else None,
        "caption": caption or "",
        "owner": author or "",
        "likes": 0,
        "comments": 0,
        "timestamp": "",
    }


def get_post_info(shortcode: str):
    try:
        # Use the global instance
        loader = get_instaloader()
        logger.info(f"Instaloader logged_in={loader.context.is_logged_in}")
        # Optional: avoid GraphQL entirely when forced
        if os.getenv('INSTAGRAM_FORCE_HTML_FALLBACK', 'false').lower() in ('1', 'true', 'yes'):
            logger.info("Force HTML fallback enabled; skipping GraphQL")
            result = _fetch_public_page_metadata(shortcode, loader.context._session)
            return result

        post = Post.from_shortcode(loader.context, shortcode)
        return {
            "shortcode": post.shortcode,
            "is_video": post.is_video,
            "url": post.url,
            "video_url": post.video_url if post.is_video else None,
            "caption": post.caption,
            "owner": post.owner_username,
            "likes": post.likes,
            "comments": post.comments,
            "timestamp": post.date_utc.isoformat(),
        }
    except Exception as e:
        err_msg = str(e)
        logger.error(f"Error getting post info for {shortcode}: {err_msg}")

        # Fallback: try scraping public page meta tags (og:video / og:image) and JSON-LD
        try:
            # Gentle backoff before fallback
            time.sleep(random.uniform(0.8, 1.6))
            loader = get_instaloader()
            session = loader.context._session
            return _fetch_public_page_metadata(shortcode, session)
        except Exception as e2:
            logger.error(f"Fallback scrape failed for {shortcode}: {e2}")
            return {"error": err_msg}