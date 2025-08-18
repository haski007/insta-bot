import os
import logging
import requests
import re
import time
import random
import json
from requests.adapters import HTTPAdapter
from urllib3.util.retry import Retry
from instaloader import Instaloader, Post

# Configure logging to emit via Uvicorn's logger so INFO-level lines are visible in container logs
logger = logging.getLogger("uvicorn.error")

# Global Instaloader instance
_L = None

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
    session_file = os.getenv('INSTAGRAM_SESSION_FILE')
    
    # Try loading session first (if path and username provided)
    if username and session_file:
        try:
            session_dir = os.path.dirname(session_file)
            if session_dir:
                os.makedirs(session_dir, exist_ok=True)
            _L.load_session_from_file(username, session_file)
            logger.info(f"Loaded Instagram session from file for user {username}")
            logger.info(f"Session load status: logged_in={_L.context.is_logged_in}")
        except Exception as e:
            logger.warning(f"No valid session file loaded ({e}); will try fresh login if password is provided")

    # Login only if not already logged in; otherwise skip misleading warnings
    if _L.context.is_logged_in:
        logger.info("Already logged in; skipping login")
    else:
        if username and password:
            try:
                logger.info(f"Attempting to login as {username}...")
                _L.login(username, password)
                logger.info(f"Successfully logged in as {username}")
                logger.info(f"Login status: logged_in={_L.context.is_logged_in}")
                # Save session for reuse if configured
                if session_file:
                    try:
                        session_dir = os.path.dirname(session_file)
                        if session_dir:
                            os.makedirs(session_dir, exist_ok=True)
                        _L.save_session_to_file(session_file)
                        logger.info("Saved Instagram session to file for reuse")
                    except Exception as e:
                        logger.error(f"Failed to save session to file: {e}")
            except Exception as e:
                logger.error(f"Login failed: {e}")
                logger.warning("Continuing without authentication (may have rate limits)")
        else:
            logger.warning("No Instagram credentials found in environment variables")
            logger.warning("Continuing without authentication (may have rate limits)")
    
    logger.info(f"Instaloader ready; logged_in={_L.context.is_logged_in}")
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
            if resp.status_code == 200 and ('og:' in resp.text or 'application/ld+json' in resp.text):
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