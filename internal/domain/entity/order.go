package entity

import (
	"time"
)

type Order struct {
	ID         string
	Price      float64
	Tax        float64
	FinalPrice float64
	CreatedAt  time.Time
}

func NewOrder(id string, price float64, tax float64) *Order {
	return &Order{
		ID:         id,
		Price:      price,
		Tax:        tax,
		FinalPrice: price + tax,
		CreatedAt:  time.Now(),
	}
}
