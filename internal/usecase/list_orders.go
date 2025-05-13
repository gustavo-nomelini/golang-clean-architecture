package usecase

import (
	"github.com/gustavo-nomelini/golang-clean-architecture/internal/domain/entity"
)

type OrderRepository interface {
	ListOrders() ([]*entity.Order, error)
	Save(order *entity.Order) error
}

type ListOrdersUseCase struct {
	OrderRepository OrderRepository
}

func NewListOrdersUseCase(orderRepository OrderRepository) *ListOrdersUseCase {
	return &ListOrdersUseCase{
		OrderRepository: orderRepository,
	}
}

func (u *ListOrdersUseCase) Execute() ([]*entity.Order, error) {
	orders, err := u.OrderRepository.ListOrders()
	if err != nil {
		return nil, err
	}
	return orders, nil
}

// CreateOrderInput represents the input data for order creation
type CreateOrderInput struct {
	ID    string  `json:"id"`
	Price float64 `json:"price"`
	Tax   float64 `json:"tax"`
}

// CreateOrderUseCase handles order creation
type CreateOrderUseCase struct {
	OrderRepository OrderRepository
}

// NewCreateOrderUseCase creates a new use case for order creation
func NewCreateOrderUseCase(orderRepository OrderRepository) *CreateOrderUseCase {
	return &CreateOrderUseCase{
		OrderRepository: orderRepository,
	}
}

// Execute creates a new order
func (u *CreateOrderUseCase) Execute(input CreateOrderInput) (*entity.Order, error) {
	order := entity.NewOrder(input.ID, input.Price, input.Tax)
	err := u.OrderRepository.Save(order)
	if err != nil {
		return nil, err
	}
	return order, nil
}
