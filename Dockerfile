# Build stage
FROM golang:1.24-alpine AS builder

# Install build dependencies
RUN apk add --no-cache gcc musl-dev

# Set working directory
WORKDIR /app

# Copy go mod files
COPY go.mod go.sum ./

# Download dependencies
RUN go mod download

# Copy source code
COPY . .

# Build the application with optimizations
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -ldflags='-w -s' -o main .

# Final stage
FROM alpine:3.19

# Install runtime dependencies
RUN apk add --no-cache ca-certificates tzdata

# Create non-root user
RUN adduser -D -g '' appuser

# Set working directory
WORKDIR /app

# Copy binary from builder
COPY --from=builder /app/main .

# Copy config files
COPY --from=builder /app/config/model.conf ./config/
COPY --from=builder /app/config/policy.csv ./config/
COPY config.yml ./

# Set ownership
RUN chown -R appuser:appuser /app

# Use non-root user
USER appuser

# Expose port
EXPOSE 8080

# Set environment variables
ENV GIN_MODE=release

# Run the binary with a startup delay to ensure database is ready
CMD ["sh", "-c", "sleep 5 && ./main"]