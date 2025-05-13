FROM golang:1.20-alpine AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .
RUN go build -o server ./cmd/server

FROM alpine:3.17

WORKDIR /app

COPY --from=builder /app/server .
COPY migrations ./migrations

EXPOSE 8080 50051

CMD ["./server"]
