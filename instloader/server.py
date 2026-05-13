from fastapi import FastAPI, Query, HTTPException
from .downloader import get_post_info, create_instaloader
import logging
import os
import sys

# Emit via Uvicorn's logger so messages appear under its configured handlers
logger = logging.getLogger("uvicorn.error")


def _env_bool(name: str, default: bool = False) -> bool:
    v = (os.getenv(name) or "").strip().lower()
    if not v:
        return default
    return v in ("1", "true", "yes", "y", "on")


app = FastAPI()

# Initialize Instaloader on startup
@app.on_event("startup")
async def startup_event():
    """Initialize Instaloader when the server starts"""
    try:
        logger.info("=== INSTALOADER STARTUP ===")
        logger.info("Initializing Instaloader...")
        loader = create_instaloader()

        if loader.context.is_logged_in:
            logger.info("✅ Instaloader is LOGGED IN and ready to use")
            logger.info("=== INSTALOADER STARTUP COMPLETE ===")
            return

        logger.error(
            "❌ Instaloader is NOT logged in — service will not start without a valid session or login. "
            "Check: ./session.json mounted to INSTAGRAM_SESSION_FILE, INSTAGRAM_USERNAME matches the session, "
            "INSTAGRAM_PASSWORD if no session, proxy/IP if Instagram blocks the host."
        )
        # Emergency dev only: INSTLOADER_ALLOW_ANONYMOUS=true
        if _env_bool("INSTLOADER_ALLOW_ANONYMOUS", default=False):
            logger.warning(
                "INSTLOADER_ALLOW_ANONYMOUS=true — continuing without login (not for production)"
            )
            logger.info("=== INSTALOADER STARTUP COMPLETE (anonymous) ===")
            return

        sys.exit(1)
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
