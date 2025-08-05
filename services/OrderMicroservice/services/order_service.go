package services

import "orderPaymentMicroservice/models"

type OrderServiceInterface interface {
	GenerateOrder(o *models.OrderCreateDTO) (*models.Order, error)
	GetOrderById(id string) (*models.Order, error)
	GetOrders() ([]models.Order, error)
	UpdateOrderStatus(id string, status string) (*models.Order, error)
}
