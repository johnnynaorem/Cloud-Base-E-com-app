package repository

import (
	"orderPaymentMicroservice/models"
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type OrderRepoImpl struct {
	DB *gorm.DB
}

func NewOrderRepoImpl(db *gorm.DB) OrderRepository {
	return &OrderRepoImpl{
		DB: db,
	}
}

func (o *OrderRepoImpl) GenerateOrder(orderCreateDTO *models.OrderCreateDTO) (*models.Order, error) {
	orderId := "ORD-" + uuid.New().String()
	newOrder := models.Order{
		OrderId:    orderId,
		CustomerId: orderCreateDTO.CustomerId,
		ProductId:  orderCreateDTO.ProductId,
		Quantity:   orderCreateDTO.Quantity,
		Status:     models.Pending,
		// PaymentStatus: "PENDING",
		TotalAmount: orderCreateDTO.TotalAmount,
		OrderDate:   time.Now().Format("2006-01-02 15:04:05"),
	}
	if err := o.DB.Create(newOrder).Error; err != nil {
		return nil, err
	}

	return &newOrder, nil
}

func (o *OrderRepoImpl) GetOrderById(id string) (*models.Order, error) {
	var order models.Order
	err := o.DB.Where("order_Id = ?", id).Find(&order).Error
	if err != nil {
		return nil, err
	}
	return &order, nil
}

func (o *OrderRepoImpl) GetOrders() ([]models.Order, error) {
	var orders []models.Order
	err := o.DB.Find(&orders).Error
	if err != nil {
		return nil, err
	}
	return orders, nil
}
func (o *OrderRepoImpl) UpdateOrderStatus(id string, status string) (*models.Order, error) {
	var updateOrder models.Order
	err := o.DB.Model(&updateOrder).Where("order_Id = ?", id).Update("status", status).Error
	if err != nil {
		return nil, err
	}
	return &updateOrder, nil
}
