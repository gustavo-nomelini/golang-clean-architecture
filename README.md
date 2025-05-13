# Golang Clean Architecture

This project demonstrates a clean architecture implementation in Go, featuring Order listing functionality through multiple interfaces:

- REST API endpoint
- gRPC service
- GraphQL query

## Project Structure

```
├── cmd
│   └── server
│       └── main.go
├── internal
│   ├── domain
│   │   └── entity
│   │       └── order.go
│   ├── infra
│   │   ├── database
│   │   │   ├── order_repository.go
│   │   │   └── db.go
│   │   ├── web
│   │   │   ├── graphql
│   │   │   │   └── handler.go
│   │   │   ├── grpc
│   │   │   │   └── service.go
│   │   │   └── rest
│   │   │       └── handler.go
│   ├── usecase
│   │   └── list_orders.go
├── migrations
│   └── 000001_create_orders_table.up.sql
│   └── 000001_create_orders_table.down.sql
├── api.http
├── Dockerfile
├── docker-compose.yaml
├── go.mod
└── go.sum
```

## Running the Application

1. Clone the repository
2. Start the application and database:
   ```
   docker compose up -d
   ```
3. The application will be available at:
   - REST API: http://localhost:8080/order
   - gRPC: localhost:50051
   - GraphQL: http://localhost:8080/graphql

## Testing with api.http

You can use the provided `api.http` file to test both creating and listing orders:

1. First create an order using the POST request
2. Then list orders using the GET request

## Technologies

- Go 1.20+
- PostgreSQL
- GORM (ORM)
- Gin (REST framework)
- gRPC
- GraphQL (gqlgen)
- Docker
