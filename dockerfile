# Stage 1: Build the Go binary
FROM golang:alpine AS builder

# Set the working directory inside the builder container
WORKDIR /app

# Copy the Go module files and source code
COPY go.mod go.sum ./
RUN go mod download
COPY . .

# Build the Go binary
RUN go build -o myapp ./cmd/main.go


# Stage 2: Final image with FFmpeg, youtube-dl in a virtual environment, and Go binary
FROM alpine:latest

# Install FFmpeg, Python, and other dependencies via apk
RUN apk add --no-cache \
    bash \
    ffmpeg \
    python3 \
    py3-pip \
    ca-certificates

# Create a virtual environment for youtube-dl
RUN python3 -m venv /opt/venv \
    && . /opt/venv/bin/activate \
    && pip install --no-cache-dir youtube-dl \
    && deactivate

# Ensure the virtual environment is available in PATH
ENV PATH="/opt/venv/bin:$PATH"

WORKDIR /app

COPY --from=builder /app/myapp .

CMD ["./myapp"]
