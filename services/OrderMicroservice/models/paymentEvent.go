package models

type PaymentEvent struct {
	OrderID string      `json:"orderId"`
	Status  OrderStatus `json:"status"`
}
