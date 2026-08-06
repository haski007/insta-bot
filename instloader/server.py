from fastapi import FastAPI, Query, HTTPException
from .downloader import (
    get_post_info,
    get_story_info,
    create_instaloader,
    get_account_pool,
    describe_accounts,
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


def _raise_from_result(result: dict):
    msg = str(result.get("error"))
    # Map known rate-limit/unauthorized messages to 429; otherwise 502
    if "Please wait a few minutes" in msg or "429" in msg or "401" in msg or "403" in msg:
        raise HTTPException(status_code=429, detail=msg)
    raise HTTPException(status_code=502, detail=msg)


app = FastAPI()

# Initialize Instaloader on startup
@app.on_event("startup")
async def startup_event():
    """Initialize Instaloader when the server starts"""
    try:
        logger.info("=== INSTALOADER STARTUP ===")
        logger.info("Initializing Instaloader...")
        # Also builds/logs any additional accounts (INSTAGRAM_USERNAME_2, ...).
        get_account_pool()
        accounts = describe_accounts()
        valid_accounts = [a for a in accounts if a["valid"]]

        for a in accounts:
            status = "✅ valid" if a["valid"] else "❌ invalid"
            logger.info(
                f"  Account {a['label']!r}: {status} "
                f"(logged_in={a['logged_in']}, sessionid_len={a['sessionid_len']})"
            )

        if valid_accounts:
            logger.info(
                f"✅ {len(valid_accounts)}/{len(accounts)} Instagram account(s) have a valid session"
            )
            logger.info("=== INSTALOADER STARTUP COMPLETE ===")
            return

        # No account has a usable session — fall back to detailed diagnostics
        # for the primary account (account 1), which is the common case.
        loader = create_instaloader()
        session_file = os.getenv("INSTAGRAM_SESSION_FILE")
        username = os.getenv("INSTAGRAM_USERNAME")
        reason = session_validation_error(
            loader, username=username, session_file=session_file
        )
        logger.error(f"❌ No Instagram account has a valid session — {reason}")
        logger.error(
            "Service will not start. Password login on VPS is often blocked by Instagram. "
            "Import session from your PC: secrets/session.json or INSTAGRAM_SESSIONID in .env. "
            "For multiple accounts, use INSTAGRAM_USERNAME_2/_3/... suffixes. "
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
        _raise_from_result(result)
    return result


@app.get("/story")
def story(
    media_id: str = Query(..., description="Numeric Instagram story media id"),
    username: str = Query("", description="Optional story owner username from the URL"),
):
    result = get_story_info(media_id, username=username)
    if isinstance(result, dict) and result.get("error"):
        _raise_from_result(result)
    return result
