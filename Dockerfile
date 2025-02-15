# Build stage
FROM golang:1.21 AS builder

# Set working directory
WORKDIR /app

# Copy go.mod and go.sum first (Leverage Docker cache)
COPY go.mod go.sum ./
RUN go mod download

# Copy the entire project
COPY . .

# Build the binary
RUN go build -o main .
RUN go build -o cli .
# Final minimal image
FROM alpine:latest

# Set working directory in the final container
WORKDIR /root/

# Copy the compiled binary from the builder stage
COPY --from=builder /app/main .

# Run the application
CMD ["./main"]
