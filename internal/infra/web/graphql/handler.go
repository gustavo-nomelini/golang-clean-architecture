package graphql

import (
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/gin-gonic/gin"
	"github.com/prodbygus/golang-clean-architecture/internal/usecase"
)

// OrderResolver is the resolver for order operations
type OrderResolver struct {
	ListOrdersUseCase  *usecase.ListOrdersUseCase
	CreateOrderUseCase *usecase.CreateOrderUseCase
}

// OrderType represents an order in GraphQL
type OrderType struct {
	ID         string  `json:"id"`
	Price      float64 `json:"price"`
	Tax        float64 `json:"tax"`
	FinalPrice float64 `json:"finalPrice"`
	CreatedAt  string  `json:"createdAt"`
}

// Query resolver for GraphQL
type QueryResolver struct {
	ListOrdersUseCase *usecase.ListOrdersUseCase
}

// Orders resolves the orders query
func (r *QueryResolver) Orders() ([]*OrderType, error) {
	orders, err := r.ListOrdersUseCase.Execute()
	if err != nil {
		return nil, err
	}

	var orderTypes []*OrderType
	for _, o := range orders {
		orderTypes = append(orderTypes, &OrderType{
			ID:         o.ID,
			Price:      o.Price,
			Tax:        o.Tax,
			FinalPrice: o.FinalPrice,
			CreatedAt:  o.CreatedAt.Format("2006-01-02T15:04:05Z07:00"),
		})
	}

	return orderTypes, nil
}

// MutationResolver for GraphQL
type MutationResolver struct {
	CreateOrderUseCase *usecase.CreateOrderUseCase
}

// CreateOrderInput represents the input for creating an order
type CreateOrderInput struct {
	ID    string  `json:"id"`
	Price float64 `json:"price"`
	Tax   float64 `json:"tax"`
}

// CreateOrder resolves the createOrder mutation
func (r *MutationResolver) CreateOrder(input CreateOrderInput) (*OrderType, error) {
	order, err := r.CreateOrderUseCase.Execute(usecase.CreateOrderInput{
		ID:    input.ID,
		Price: input.Price,
		Tax:   input.Tax,
	})
	if err != nil {
		return nil, err
	}

	return &OrderType{
		ID:         order.ID,
		Price:      order.Price,
		Tax:        order.Tax,
		FinalPrice: order.FinalPrice,
		CreatedAt:  order.CreatedAt.Format("2006-01-02T15:04:05Z07:00"),
	}, nil
}

// SetupGraphQL sets up the GraphQL handler
func SetupGraphQL(engine *gin.Engine, listOrdersUseCase *usecase.ListOrdersUseCase, createOrderUseCase *usecase.CreateOrderUseCase) {
	// In a real implementation, we would use gqlgen to generate resolvers
	// For this example, we'll create a basic handler
	engine.GET("/graphql", gin.WrapH(playground.Handler("GraphQL playground", "/query")))

	// In a real app, this would be properly implemented with gqlgen
	engine.POST("/query", func(c *gin.Context) {
		c.JSON(200, gin.H{"message": "GraphQL endpoint. In a real app, this would be implemented with gqlgen"})
	})
}
