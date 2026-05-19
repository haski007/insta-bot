FROM golang:latest

# Install ffmpeg for rendering anglicism card video replies.
RUN apt-get update \
    && apt-get install -y --no-install-recommends ffmpeg \
    && rm -rf /var/lib/apt/lists/*

# Copy the source code for the bot into the image
COPY . /insta-bot

# Set the working directory to the app directory
WORKDIR /insta-bot

# Build the bot
RUN make build

# Run the bot
CMD ["./build/insta-bot"]
