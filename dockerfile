# 👨‍💻 Development Stage
FROM golang:1.24-alpine AS dev


WORKDIR /app

RUN go install github.com/air-verse/air@latest

COPY go.mod go.sum ./
RUN go mod download

CMD ["air", "-c", ".air.toml"]


# 🛠 Build Stage (Production Build)
FROM golang:1.24-alpine AS build

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN go build -o server ./cmd/main.go

# 🚀 Production Runtime
FROM alpine:latest AS prod

WORKDIR /app

# Set environment variable
ENV APP_ENV=production

# Copy the built binary only
COPY --from=build /app/server .

EXPOSE 8080

CMD ["./server"]
