from fastapi import FastAPI, Query, HTTPException
from .downloader import (
    get_post_info,
    create_instaloader,
    has_valid_session,
    get_sessionid,
    session_validation_error,
)
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
        session_file = os.getenv("INSTAGRAM_SESSION_FILE")
        username = os.getenv("INSTAGRAM_USERNAME")

        if has_valid_session(loader):
            logger.info(
                f"✅ Valid Instagram session (sessionid length={len(get_sessionid(loader))})"
            )
            logger.info("=== INSTALOADER STARTUP COMPLETE ===")
            return

        reason = session_validation_error(
            loader, username=username, session_file=session_file
        )
        logger.error(f"❌ Invalid Instagram session — {reason}")
        logger.error(
            "Service will not start. Password login on VPS is often blocked by Instagram. "
            "Import session from your PC: secrets/session.json or INSTAGRAM_SESSIONID in .env. "
            "See secrets/README.md and scripts/export_instagram_session.py"
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
        if _env_bool("INSTLOADER_ALLOW_ANONYMOUS", default=False):
            logger.warning("INSTLOADER_ALLOW_ANONYMOUS=true — continuing despite init failure")
            return
        sys.exit(1)


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
