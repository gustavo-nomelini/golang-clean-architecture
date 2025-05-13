package database

import (
	"github.com/prodbygus/golang-clean-architecture/internal/domain/entity"
	"gorm.io/gorm"
)

type OrderRepository struct {
	DB *gorm.DB
}

func NewOrderRepository(db *gorm.DB) *OrderRepository {
	return &OrderRepository{DB: db}
}

func (r *OrderRepository) Save(order *entity.Order) error {
	orderModel := OrderModel{
		ID:         order.ID,
		Price:      order.Price,
		Tax:        order.Tax,
		FinalPrice: order.FinalPrice,
		CreatedAt:  order.CreatedAt,
	}
	return r.DB.Create(&orderModel).Error
}

func (r *OrderRepository) ListOrders() ([]*entity.Order, error) {
	var orderModels []OrderModel
	err := r.DB.Find(&orderModels).Error
	if err != nil {
		return nil, err
	}

	orders := make([]*entity.Order, len(orderModels))
	for i, model := range orderModels {
		orders[i] = &entity.Order{
			ID:         model.ID,
			Price:      model.Price,
			Tax:        model.Tax,
			FinalPrice: model.FinalPrice,
			CreatedAt:  model.CreatedAt,
		}
	}

	return orders, nil
}
