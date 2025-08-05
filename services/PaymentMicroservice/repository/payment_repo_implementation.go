package repository

import (
	"context"
	"math/rand"
	"fmt"

	"paymentMicroservice/models"
	"paymentMicroservice/pub"
	"time"

	"cloud.google.com/go/pubsub"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type PaymentRepoImpl struct {
	DB *gorm.DB
}

func NewPaymentRepoImpl(db *gorm.DB) PaymentRepo {
	return &PaymentRepoImpl{DB: db}
}

const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

func randomID(length int) string {
	seededRand := rand.New(rand.NewSource(time.Now().UnixNano()))
	b := make([]byte, length)
	for i := range b {
		b[i] = charset[seededRand.Intn(len(charset))]
	}
	return string(b)
}

func (p *PaymentRepoImpl) MakePayment(client *pubsub.Client, ctx context.Context, paymentDTO *models.PaymentCreate) error {
	paymentID := randomID(10)
	transactionId := "TRANS-" + uuid.New().String()
	newPayment := models.Payment{
		ID:            paymentID,
		OrderID:       paymentDTO.OrderID,
		TransactionID: transactionId,
		Method:        paymentDTO.Method,
		Status:        models.PaymentSuccess,
		Amount:        paymentDTO.Amount,
		PaidAt:        time.Now().Format("2006-01-02 15:04:05"),
	}
	if err := p.DB.Create(newPayment).Error; err != nil {
		return err
	}

	event := &models.PaymentEvent{
		OrderID: paymentDTO.OrderID,
		Status:  models.Processing,
	}
	fmt.Println("PublishOrderCreated")
	pub.PublishOrderCreated(ctx, client, *event)
	return nil
}

func (p *PaymentRepoImpl) GetPaymentById(id string) (*models.Payment, error) {
	var payment models.Payment
	err := p.DB.Where("order_id = ?", id).Find(&payment).Error
	if err != nil {
		return nil, err
	}
	return &payment, nil
}

func (p *PaymentRepoImpl) UpdatePaymentStatus(id string, newStatus string) (*models.Payment, error) {
	var updatePayment models.Payment
	err := p.DB.Model(&updatePayment).Where("order_id = ?", id).Update("status", newStatus).Error
	if err != nil {
		return nil, err
	}
	return &updatePayment, nil
}
