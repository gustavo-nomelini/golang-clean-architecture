package main

import (
	"fmt"
	"log"
	"net"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gustavo-nomelini/golang-clean-architecture/internal/infra/database"
	"github.com/gustavo-nomelini/golang-clean-architecture/internal/infra/web/graphql"
	"github.com/gustavo-nomelini/golang-clean-architecture/internal/infra/web/grpc"
	"github.com/gustavo-nomelini/golang-clean-architecture/internal/infra/web/rest"
	"github.com/gustavo-nomelini/golang-clean-architecture/internal/usecase"
	ggrpc "google.golang.org/grpc"
)

func main() {
	// Connect to database
	db := database.Connect()

	// Setup repository
	orderRepository := database.NewOrderRepository(db)

	// Setup use cases
	listOrdersUseCase := usecase.NewListOrdersUseCase(orderRepository)
	createOrderUseCase := usecase.NewCreateOrderUseCase(orderRepository)

	// Setup REST API
	engine := gin.Default()
	orderHandler := rest.NewOrderHandler(listOrdersUseCase, createOrderUseCase)
	rest.SetupRoutes(engine, orderHandler)

	// Setup GraphQL
	graphql.SetupGraphQL(engine, listOrdersUseCase, createOrderUseCase)

	// Start gRPC server
	go startGRPCServer(listOrdersUseCase, createOrderUseCase)

	// Start HTTP server
	fmt.Println("Starting HTTP server on :8080")
	if err := http.ListenAndServe(":8080", engine); err != nil {
		log.Fatalf("Failed to start HTTP server: %v", err)
	}
}

func startGRPCServer(listOrdersUseCase *usecase.ListOrdersUseCase, createOrderUseCase *usecase.CreateOrderUseCase) {
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	s := ggrpc.NewServer()
	server := grpc.NewOrderServiceServer(listOrdersUseCase, createOrderUseCase)
	grpc.Register(s, server)

	fmt.Println("Starting gRPC server on :50051")
	if err := s.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
