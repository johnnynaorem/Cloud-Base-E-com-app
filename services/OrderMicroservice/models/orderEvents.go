package models

type OrderEvent struct {
	OrderID     string  `json:"orderId"`
	TotalAmount float64 `json:"totalAmount"`
}
