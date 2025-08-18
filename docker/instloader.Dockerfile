FROM python:3.11-slim

WORKDIR /app

# Install runtime dependencies (include SOCKS support for proxies)
RUN pip install --no-cache-dir \
    fastapi \
    "uvicorn[standard]" \
    instaloader \
    "requests[socks]"

# Copy the instloader directory
COPY instloader/ /app/instloader/

EXPOSE 8003

CMD ["uvicorn", "instloader.server:app", "--host", "0.0.0.0", "--port", "8003"]
