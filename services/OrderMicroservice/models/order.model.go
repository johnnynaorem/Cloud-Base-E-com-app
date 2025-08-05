package models

type OrderStatus string

const (
	Pending    OrderStatus = "PENDING"
	Processing OrderStatus = "PROCESSING"
	Shipped    OrderStatus = "SHIPPED"
	Delivered  OrderStatus = "DELIVERED"
	Canceled   OrderStatus = "CANCELED"
)

type Order struct {
	OrderId     string      `json:"orderId"`
	CustomerId  string      `json:"customerId"`
	ProductId   string      `json:"productId"`
	Quantity    int         `json:"quantity"`
	TotalAmount float64     `json:"totalAmount"`
	OrderDate   string      `json:"orderDate"`
	Status      OrderStatus `json:"status"`
	// PaymentStatus string      `json:"paymentStatus"`
}

type OrderCreateDTO struct {
	CustomerId  string  `json:"customerId"`
	ProductId   string  `json:"productId"`
	Quantity    int     `json:"quantity"`
	TotalAmount float64 `json:"totalAmount"`
}

type OrderStatusUpdateDTO struct {
	OrderId string      `json:"orderId"`
	Status  OrderStatus `json:"status"`
}
