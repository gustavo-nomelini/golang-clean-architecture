FROM golang:1.24-alpine AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .
# Add go mod tidy before building
RUN go mod tidy && go build -o server ./cmd/server

FROM alpine:3.17

WORKDIR /app

COPY --from=builder /app/server .
COPY migrations ./migrations

EXPOSE 8080 50051

CMD ["./server"]
