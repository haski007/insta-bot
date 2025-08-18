from fastapi import FastAPI, Query, HTTPException
from .downloader import get_post_info, create_instaloader
import logging

# Emit via Uvicorn's logger so messages appear under its configured handlers
logger = logging.getLogger("uvicorn.error")
logger.propagate = False

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
    if isinstance(result, dict) and result.get("error"):
        msg = str(result.get("error"))
        # Map known rate-limit/unauthorized messages to 429; otherwise 502
        if "Please wait a few minutes" in msg or "429" in msg or "401" in msg or "403" in msg:
            raise HTTPException(status_code=429, detail=msg)
        raise HTTPException(status_code=502, detail=msg)
    return result