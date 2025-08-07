import os
from instaloader import Instaloader, Post

def create_instaloader():
    """Create and optionally login to Instaloader"""
    L = Instaloader()
    
    # Get credentials from environment variables
    username = os.getenv('INSTAGRAM_USERNAME')
    password = os.getenv('INSTAGRAM_PASSWORD')
    
    # Login if credentials are provided
    if username and password:
        try:
            L.login(username, password)
            print(f"Successfully logged in as {username}")
        except Exception as e:
            print(f"Login failed: {e}")
            print("Continuing without authentication (may have rate limits)")
    else:
        print("No Instagram credentials found in environment variables")
        print("Continuing without authentication (may have rate limits)")
    
    return L

# Initialize Instaloader with optional login
L = create_instaloader()

def get_post_info(shortcode: str):
    try:
        post = Post.from_shortcode(L.context, shortcode)
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
        return {"error": str(e)}