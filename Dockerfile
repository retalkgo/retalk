# Build stage
FROM golang:1.24-alpine AS builder

WORKDIR /app

# Install build dependencies
RUN apk add --no-cache git make

# Copy go mod files
COPY go.mod go.sum ./
RUN go mod download

# Copy source code
COPY . .

# Build the application
ARG VERSION=docker
ARG COMMIT=unknown
RUN go build -o bin/retalk -v \
    -ldflags "-w -s -X 'github.com/retalkgo/retalk/internal/version.Version=${VERSION}' \
    -X 'github.com/retalkgo/retalk/internal/version.Commit=${COMMIT}'"

# Runtime stage
FROM alpine:latest

WORKDIR /app

# Install runtime dependencies
RUN apk add --no-cache ca-certificates tzdata

# Copy binary from builder
COPY --from=builder /app/bin/retalk /app/retalk

# Copy example config (user should mount their own config)
COPY retalk.example.yaml /app/retalk.example.yaml

# Create non-root user
RUN addgroup -g 1000 retalk && \
    adduser -D -u 1000 -G retalk retalk && \
    chown -R retalk:retalk /app

USER retalk

# Expose default port
EXPOSE 2716

# Run the application
CMD ["/app/retalk"]
