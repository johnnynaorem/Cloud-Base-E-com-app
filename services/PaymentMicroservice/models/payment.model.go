package models

type PaymentStatus string

const (
	PaymentPending PaymentStatus = "PENDING"
	PaymentSuccess PaymentStatus = "SUCCESS"
	PaymentFailed  PaymentStatus = "FAILED"
)

type PaymentMethod string

const (
	CreditCard PaymentMethod = "CREDIT_CARD"
	DebitCard  PaymentMethod = "DEBIT_CARD"
	PayPal     PaymentMethod = "PAYPAL"
	UPI        PaymentMethod = "UPI"
	NetBanking PaymentMethod = "NET_BANKING"
)

type Payment struct {
	ID            string        `gorm:"type:char(50);primaryKey"`
	OrderID       string        `gorm:"not null;index"`
	Amount        float64       `gorm:"not null"`
	Status        PaymentStatus `gorm:"type:varchar(20);default:'PENDING'"`
	Method        PaymentMethod `gorm:"type:varchar(30);not null"`
	TransactionID string        `gorm:"type:varchar(100);uniqueIndex"`
	PaidAt        string
}

type PaymentCreate struct {
	OrderID string        `gorm:"not null;index"`
	Amount  float64       `gorm:"not null"`
	Method  PaymentMethod `gorm:"type:varchar(30);not null"`
}

type PaymentStatusUpdateDTO struct {
	OrderID string
	Status  string
}

type OrderCreatedEvent struct {
	OrderID     string  `json:"orderId"`
	TotalAmount float64 `json:"totalAmount"`
}

type OrderStatus string

const (
	Pending    OrderStatus = "PENDING"
	Processing OrderStatus = "PROCESSING"
	Shipped    OrderStatus = "SHIPPED"
	Delivered  OrderStatus = "DELIVERED"
	Canceled   OrderStatus = "CANCELED"
)

type PaymentEvent struct {
	OrderID string      `json:"orderId"`
	Status  OrderStatus `json:"status"`
}
