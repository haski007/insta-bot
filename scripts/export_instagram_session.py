#!/usr/bin/env python3
"""Export Instagram session for insta-bot (run on your Mac/PC, not VPS).

Methods (first success wins):
  1. Browser cookies — log into instagram.com in Chrome/Firefox/Safari first
  2. INSTAGRAM_SESSIONID — paste sessionid cookie from DevTools
  3. Password login — often blocked; use only if 1–2 fail

Examples:
  pip3 install instaloader browser-cookie3

  # Best: already logged in in browser
  export INSTAGRAM_USERNAME=your_user
  python3 scripts/export_instagram_session.py --browser chrome

  # Or paste cookie from DevTools → Application → Cookies → sessionid
  export INSTAGRAM_USERNAME=your_user
  export INSTAGRAM_SESSIONID='...'
  python3 scripts/export_instagram_session.py

  scp session.json user@vps:~/projects/insta-bot/secrets/session.json
"""
from __future__ import annotations

import argparse
import os
import sys
from urllib.parse import unquote

from instaloader import Instaloader
from instaloader.exceptions import TwoFactorAuthRequiredException


def apply_proxy(loader: Instaloader) -> None:
    proxy = (os.environ.get("INSTAGRAM_PROXY") or "").strip()
    if proxy:
        loader.context._session.proxies = {"http": proxy, "https": proxy}
        print(f"Using proxy: {proxy.split('@')[-1] if '@' in proxy else proxy}")


def sessionid_len(loader: Instaloader) -> int:
    return len((loader.context._session.cookies.get("sessionid") or "").strip())


def save_if_ok(loader: Instaloader, username: str, out: str) -> bool:
    sid_len = sessionid_len(loader)
    if sid_len < 10:
        return False
    try:
        ok = loader.test_login()
        print(f"test_login()={ok}, sessionid length={sid_len}")
    except Exception as e:
        print(f"test_login() raised {e!r} (sessionid length={sid_len}) — saving anyway")
    loader.save_session_to_file(out)
    print(f"OK: wrote {out!r}")
    print(f"  scp {out} user@vps:~/projects/insta-bot/secrets/session.json")
    return True


def try_browser_cookies(loader: Instaloader, username: str, browser: str | None) -> bool:
    try:
        import browser_cookie3
    except ImportError:
        print("Install browser cookies support: pip3 install browser-cookie3", file=sys.stderr)
        return False

    browsers: list[tuple[str, object]] = []
    if browser:
        fn = getattr(browser_cookie3, browser, None)
        if fn is None:
            print(f"Unknown browser {browser!r}. Try: chrome, firefox, safari, edge", file=sys.stderr)
            return False
        browsers = [(browser, fn)]
    else:
        for name in ("chrome", "firefox", "safari", "edge", "brave"):
            fn = getattr(browser_cookie3, name, None)
            if fn is not None:
                browsers.append((name, fn))

    for name, fn in browsers:
        try:
            print(f"Trying cookies from {name}...")
            jar = fn(domain_name=".instagram.com")
            loader.context._session.cookies.update(jar)
            if sessionid_len(loader) >= 10:
                print(f"Got sessionid from {name}")
                loader.context.username = username
                return True
        except Exception as e:
            print(f"  {name}: {e}")
    return False


def build_session_data(sessionid: str) -> dict:
    """Instaloader 4.x load_session() expects a cookie dict, not a bare sessionid string."""
    ds_user_id = (os.environ.get("INSTAGRAM_DS_USER_ID") or "").strip()
    if not ds_user_id:
        head = sessionid.split(":", 1)[0]
        if head.isdigit():
            ds_user_id = head
    return {
        "sessionid": sessionid,
        "ds_user_id": ds_user_id,
        "csrftoken": (os.environ.get("INSTAGRAM_CSRFTOKEN") or "imported").strip(),
        "mid": (os.environ.get("INSTAGRAM_MID") or "").strip(),
        "ig_pr": "1",
        "ig_vw": "1920",
        "ig_cb": "1",
        "s_network": "",
        "ig_did": (os.environ.get("INSTAGRAM_IG_DID") or "").strip(),
    }


def try_sessionid_env(loader: Instaloader, username: str) -> bool:
    raw = (os.environ.get("INSTAGRAM_SESSIONID") or "").strip().strip('"').strip("'")
    if not raw:
        return False
    sessionid = unquote(raw)
    print("Loading INSTAGRAM_SESSIONID from env...")
    loader.load_session(username, build_session_data(sessionid))
    return sessionid_len(loader) >= 10


def try_password_login(loader: Instaloader, username: str, password: str) -> bool:
    print(f"Trying password login as {username}...")
    try:
        loader.login(username, password)
    except TwoFactorAuthRequiredException:
        code = (os.environ.get("INSTAGRAM_2FA_CODE") or "").strip() or input("2FA code: ").strip()
        loader.two_factor_login(code)
    except Exception as e:
        print(f"Login error: {e!r}", file=sys.stderr)
        return False
    return sessionid_len(loader) >= 10


def main() -> int:
    parser = argparse.ArgumentParser(description="Export Instagram session for insta-bot")
    parser.add_argument(
        "--browser",
        metavar="NAME",
        help="Import cookies from browser (chrome, firefox, safari, …). Log into instagram.com first.",
    )
    parser.add_argument(
        "--no-password",
        action="store_true",
        help="Skip password login (recommended when using --browser or INSTAGRAM_SESSIONID)",
    )
    args = parser.parse_args()

    username = (os.environ.get("INSTAGRAM_USERNAME") or "").strip().strip('"')
    password = (os.environ.get("INSTAGRAM_PASSWORD") or "").strip().strip('"')
    out = os.environ.get("SESSION_OUT", "session.json")

    if not username:
        print("Set INSTAGRAM_USERNAME", file=sys.stderr)
        return 1

    loader = Instaloader()
    apply_proxy(loader)

    # 1) Browser cookies
    if args.browser is not None or os.environ.get("INSTAGRAM_USE_BROWSER", "").lower() in ("1", "true", "yes"):
        if try_browser_cookies(loader, username, args.browser):
            return 0 if save_if_ok(loader, username, out) else 1

    # 2) sessionid from env
    if try_sessionid_env(loader, username):
        return 0 if save_if_ok(loader, username, out) else 1

    # 3) Password (unreliable)
    if not args.no_password and password:
        if try_password_login(loader, username, password):
            return 0 if save_if_ok(loader, username, out) else 1

    print(
        "\nAll methods failed.\n\n"
        "Do this:\n"
        "  1. Open https://www.instagram.com/ in Chrome and log in\n"
        "  2. pip3 install browser-cookie3\n"
        f"  3. python3 {sys.argv[0]} --browser chrome\n\n"
        "Or copy sessionid manually:\n"
        "  DevTools → Application → Cookies → instagram.com → sessionid\n"
        "  export INSTAGRAM_SESSIONID='paste_here'\n"
        f"  python3 {sys.argv[0]} --no-password\n",
        file=sys.stderr,
    )
    return 1


if __name__ == "__main__":
    raise SystemExit(main())
