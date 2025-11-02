# ---------------------------
# 1️⃣ Build stage
# ---------------------------
FROM golang:1.25.3-alpine AS builder

# Install git for fetching dependencies
RUN apk add --no-cache git

# Set working directory
WORKDIR /app

# Copy go.mod and go.sum
COPY go.mod go.sum ./

# Download dependencies
RUN go mod download

# Copy the rest of the project
COPY . .

# Build the binary
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o gin-app ./...

# ---------------------------
# 2️⃣ Run stage
# ---------------------------
FROM alpine:latest

# Set working directory
WORKDIR /root/

# Copy binary from builder
COPY --from=builder /app/gin-app .

# Expose HTTP port
EXPOSE 8000

# Run the application
CMD ["./gin-app"]
