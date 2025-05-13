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

## Troubleshooting

### Go Version Mismatch

If you encounter errors related to Go version mismatch when running `docker compose up -d`, ensure that:

1. The Go version in your `go.mod` file matches the version used in the Dockerfile
2. The current version in `go.mod` is `go 1.20` to match the Docker image `golang:1.20-alpine`
3. Alternatively, you can update the Dockerfile to use a newer Go version like `golang:1.24-alpine`

To fix version mismatch issues, either:

- Modify your go.mod file: `go 1.20` instead of a newer version
- Update the Dockerfile to use a newer Go version

## Implementações Recentes

### Script de Configuração

Foi adicionado um script `setup.sh` para facilitar a configuração e inicialização do projeto:

1. Executa `go mod tidy` para atualizar as dependências no arquivo go.mod
2. Inicia os containers Docker com `docker compose up -d`

Para utilizar o script:

```bash
chmod +x setup.sh
./setup.sh
```

### Compatibilidade de Versões Go

O projeto foi atualizado para usar Go 1.24 no Dockerfile, garantindo compatibilidade com códigos mais recentes. As seguintes mudanças foram implementadas:

1. Atualização do Dockerfile para utilizar a imagem `golang:1.24-alpine`
2. Adição de etapa `go mod tidy` no processo de build para resolver incompatibilidades de dependências

### Solução de Problemas Comuns

Agora existe uma seção de troubleshooting que ajuda a resolver problemas relacionados a:

- Incompatibilidade de versões do Go
- Dependências desatualizadas
- Erros durante o build Docker

Para garantir o funcionamento correto da aplicação, certifique-se de que:

- As dependências estão atualizadas usando `go mod tidy`
- O arquivo go.mod está compatível com a versão do Go utilizada no Dockerfile
- Os containers Docker estão sendo iniciados corretamente
