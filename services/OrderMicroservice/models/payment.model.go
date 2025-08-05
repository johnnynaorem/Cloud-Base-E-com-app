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
	ID            string        `gorm:"type:char(36);primaryKey"`
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
