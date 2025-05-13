package rest

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gustavo-nomelini/golang-clean-architecture/internal/usecase"
)

type OrderHandler struct {
	ListOrdersUseCase  *usecase.ListOrdersUseCase
	CreateOrderUseCase *usecase.CreateOrderUseCase
}

func NewOrderHandler(
	listOrdersUseCase *usecase.ListOrdersUseCase,
	createOrderUseCase *usecase.CreateOrderUseCase,
) *OrderHandler {
	return &OrderHandler{
		ListOrdersUseCase:  listOrdersUseCase,
		CreateOrderUseCase: createOrderUseCase,
	}
}

func (h *OrderHandler) ListOrders(c *gin.Context) {
	orders, err := h.ListOrdersUseCase.Execute()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, orders)
}

func (h *OrderHandler) CreateOrder(c *gin.Context) {
	var input usecase.CreateOrderInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	order, err := h.CreateOrderUseCase.Execute(input)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, order)
}

func SetupRoutes(engine *gin.Engine, orderHandler *OrderHandler) {
	engine.GET("/order", orderHandler.ListOrders)
	engine.POST("/order", orderHandler.CreateOrder)
}
