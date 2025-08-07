from fastapi import FastAPI, Query
from .downloader import get_post_info, create_instaloader
import logging

# Configure logging
logging.basicConfig(level=logging.INFO)
logger = logging.getLogger(__name__)

app = FastAPI()

# Initialize Instaloader on startup
@app.on_event("startup")
async def startup_event():
    """Initialize Instaloader when the server starts"""
    try:
        logger.info("=== INSTALOADER STARTUP ===")
        logger.info("Initializing Instaloader...")
        # This will trigger the login process
        loader = create_instaloader()
        
        # Check if we're logged in
        if loader.context.is_logged_in:
            logger.info("✅ Instaloader is LOGGED IN and ready to use")
        else:
            logger.warning("⚠️  Instaloader is NOT logged in - will have rate limits")
        
        logger.info("=== INSTALOADER STARTUP COMPLETE ===")
    except Exception as e:
        logger.error(f"❌ Failed to initialize Instaloader: {e}")
        # Continue running the server even if Instaloader fails to initialize

@app.get("/")
def health_check():
    return {"status": "healthy", "service": "instloader"}

@app.get("/media")
def media(shortcode: str = Query(...)):
    result = get_post_info(shortcode)
    return result