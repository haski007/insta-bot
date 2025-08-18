import os
import logging
import requests
import re
import time
import random
from requests.adapters import HTTPAdapter
from urllib3.util.retry import Retry
from instaloader import Instaloader, Post

# Configure logging
logger = logging.getLogger(__name__)

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
        except Exception as e:
            logger.warning(f"No valid session file loaded ({e}); will try fresh login if password is provided")

    # Login if credentials are provided and not already logged in
    if username and password and not _L.context.is_logged_in:
        try:
            logger.info(f"Attempting to login as {username}...")
            _L.login(username, password)
            logger.info(f"Successfully logged in as {username}")
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
    
    return _L

def get_instaloader():
    """Get the global Instaloader instance, creating it if necessary"""
    global _L
    if _L is None:
        logger.info("Creating new Instaloader instance")
        _L = create_instaloader()
    return _L

def get_post_info(shortcode: str):
    try:
        # Use the global instance
        loader = get_instaloader()
        logger.info(f"Instaloader logged_in={loader.context.is_logged_in}")
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

        # Fallback: try scraping public page meta tags (og:video / og:image)
        try:
            # Gentle backoff before fallback
            time.sleep(random.uniform(0.8, 1.6))

            loader = get_instaloader()
            session = loader.context._session
            headers = dict(session.headers)
            headers.setdefault('User-Agent', os.getenv('INSTAGRAM_USER_AGENT', headers.get('User-Agent', 'Mozilla/5.0')))
            headers.setdefault('Accept', 'text/html,application/xhtml+xml,application/xml;q=0.9,image/avif,image/webp,*/*;q=0.8')
            headers.setdefault('Accept-Language', 'en-US,en;q=0.9')

            candidate_urls = [
                f"https://www.instagram.com/p/{shortcode}/",
                f"https://www.instagram.com/reel/{shortcode}/",
            ]

            html = None
            for u in candidate_urls:
                try:
                    resp = session.get(u, headers=headers, timeout=20)
                    if resp.status_code == 200 and 'og:' in resp.text:
                        html = resp.text
                        break
                except Exception as e2:
                    logger.warning(f"Fallback fetch failed for {u}: {e2}")

            if not html:
                return {"error": err_msg}

            # Extract metadata
            def meta_content(prop: str) -> str:
                # Simple regex for meta tags
                m = re.search(rf'<meta[^>]+property=["\']{re.escape(prop)}["\'][^>]+content=["\']([^"\']+)["\']', html, re.IGNORECASE)
                return m.group(1) if m else ""

            og_video = meta_content('og:video') or meta_content('og:video:secure_url')
            og_image = meta_content('og:image')
            og_title = meta_content('og:title')
            og_desc = meta_content('og:description')
            author = meta_content('instagram:owner_user_name') or meta_content('og:site_name')

            is_video = bool(og_video)
            return {
                "shortcode": shortcode,
                "is_video": is_video,
                "url": og_image,
                "video_url": og_video if is_video else None,
                "caption": og_title or og_desc,
                "owner": author,
                "likes": 0,
                "comments": 0,
                "timestamp": "",
            }
        except Exception as e2:
            logger.error(f"Fallback scrape failed for {shortcode}: {e2}")
            return {"error": err_msg}