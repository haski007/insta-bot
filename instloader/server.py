from fastapi import FastAPI, Query
from .downloader import get_post_info

app = FastAPI()

@app.get("/")
def health_check():
    return {"status": "healthy", "service": "instloader"}

@app.get("/media")
def media(shortcode: str = Query(...)):
    result = get_post_info(shortcode)
    return result