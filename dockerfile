# 🛠 STAGE 1: Build the Go binary
FROM golang:1.24-alpine AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

# Build a statically linked binary
RUN go build -o server ./cmd/main.go

# 📦 STAGE 2: Run the binary in a minimal container
FROM alpine:latest

WORKDIR /app

# Copy only the built binary from the previous stage
COPY --from=builder /app/server .

# Set port (optional, for docs)
EXPOSE 8080

CMD ["./server"]
