FROM golang:latest
# Copy the source code for the bot into the image
COPY . /insta-bot

# Set the working directory to the app directory
WORKDIR /insta-bot

# Build the bot
RUN make build


# Run the bot
CMD ["./build/insta-bot", "-config", "config/local.yaml", "-log_level", "DEBUG"]
