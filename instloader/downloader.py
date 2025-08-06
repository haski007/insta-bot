from instaloader import Instaloader, Post

# Load session only once - without authentication
L = Instaloader()

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