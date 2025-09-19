# 1. Build Stage
# Uses the offical golang:1.22 image
FROM golang:1.25 AS builder
LABEL authors="simonaanchova"

# Set working directory inside container
WORKDIR /app

# Copy go.mod and go.sum first for dependency caching
COPY go.mod go.sum ./
RUN go mod download

# Copy the rest of the source code
COPY . .

# Build the Go app
RUN CGO_ENABLED=0  GOOS=linux go build -o document-service ./main.go

# 2. Runtime Stage

FROM alpine:latest

WORKDIR /root/

# Copy the compiled binary from builder
COPY --from=builder /app/document-service .

# Expose service port
EXPOSE 8080

# Run the binary
CMD ["./document-service"]

