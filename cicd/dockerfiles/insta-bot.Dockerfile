FROM golang:1.21 as builder

# Set the working dixrectory in the container
WORKDIR /app

# Copy the Go modules and sum files
COPY go.mod go.sum ./

# Download Go module dependencies
RUN go mod download

# Copy the rest of the source code
COPY . .

# Build the application
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o insta-bot ./cmd/app

# Use a small image to run the application
FROM alpine:latest
RUN apk --no-cache add ca-certificates

WORKDIR /root/

# Copy the pre-built binary file from the previous stage
COPY --from=builder /app/insta-bot .


# Command to run the executable
CMD ["./insta-bot"]
