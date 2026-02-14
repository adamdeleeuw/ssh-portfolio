# Stage 1: Build
FROM golang:1.26-alpine AS builder

# Install build dependencies
RUN apk add --no-cache git

# Set working directory
WORKDIR /build

# Copy go mod files
COPY go.mod go.sum ./
RUN go mod download

# Copy source code
COPY . .

# Build binary (static binary for alpine)
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o ssh-portfolio ./cmd/server

# Stage 2: Runtime
FROM alpine:latest

# Install ca-certificates for HTTPS
RUN apk --no-cache add ca-certificates

# Create non-root user
RUN addgroup -S appgroup && adduser -S appuser -G appgroup

# Create data directory for host keys and database
RUN mkdir -p /data /app/content && \
    chown -R appuser:appgroup /data /app

# Set working directory
WORKDIR /app

# Copy binary from builder
COPY --from=builder /build/ssh-portfolio .

# Copy content files
COPY --chown=appuser:appgroup content /app/content

# Switch to non-root user
USER appuser

# Expose SSH port (standard port 22)
EXPOSE 22

# Set environment variables
ENV PORT=22 \
    HOST_KEY_PATH=/data/ssh_host_ed25519_key \
    CONTENT_DIR=/app/content

# Health check
HEALTHCHECK --interval=30s --timeout=3s --start-period=5s --retries=3 \
    CMD nc -z localhost 22 || exit 1

# Run the binary
ENTRYPOINT ["./ssh-portfolio"]
