import os
import logging
from instaloader import Instaloader, Post

# Configure logging
logger = logging.getLogger(__name__)

# Global Instaloader instance
_L = None

def create_instaloader():
    """Create and optionally login to Instaloader"""
    global _L
    
    if _L is not None:
        logger.info("Instaloader already initialized, returning existing instance")
        return _L
    
    logger.info("Creating new Instaloader instance...")
    _L = Instaloader()
    
    # Configure proxy if provided
    proxy = os.getenv('INSTAGRAM_PROXY')
    if proxy:
        try:
            logger.info(f"Configuring proxy: {proxy}")
            _L.context._session.proxies = {
                'http': proxy,
                'https': proxy
            }
            logger.info("Proxy configured successfully")
        except Exception as e:
            logger.error(f"Failed to configure proxy: {e}")
    
    # Get credentials from environment variables
    username = os.getenv('INSTAGRAM_USERNAME')
    password = os.getenv('INSTAGRAM_PASSWORD')
    
    # Login if credentials are provided
    if username and password:
        try:
            logger.info(f"Attempting to login as {username}...")
            _L.login(username, password)
            logger.info(f"Successfully logged in as {username}")
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
        logger.info("|||||||||||||||||||||||||||||||||||||")
        loader = get_instaloader()
        logger.info(loader)
        # check if loader is logged in
        logger.info(loader.context.is_logged_in)
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
        logger.error(f"Error getting post info for {shortcode}: {e}")
        return {"error": str(e)}