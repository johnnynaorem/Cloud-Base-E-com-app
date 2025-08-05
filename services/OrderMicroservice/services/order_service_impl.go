package services

import (
	"orderPaymentMicroservice/models"
	"orderPaymentMicroservice/repository"
)

type OrderService struct {
	Repo repository.OrderRepository
}

func (os *OrderService) GenerateOrder(o *models.OrderCreateDTO) (*models.Order, error) {
	return os.Repo.GenerateOrder(o)
}

func (os *OrderService) GetOrderById(id string) (*models.Order, error) {
	return os.Repo.GetOrderById(id)
}

func (os *OrderService) GetOrders() ([]models.Order, error) {
	return os.Repo.GetOrders()
}
func (os *OrderService) UpdateOrderStatus(id string, status string) (*models.Order, error) {
	return os.Repo.UpdateOrderStatus(id, status)
}
