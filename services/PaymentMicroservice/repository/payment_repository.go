package repository

import (
	"context"
	"paymentMicroservice/models"

	"cloud.google.com/go/pubsub"
)

type PaymentRepo interface {
	MakePayment(client *pubsub.Client, ctx context.Context, paymentCred *models.PaymentCreate) error
	GetPaymentById(id string) (*models.Payment, error)
	UpdatePaymentStatus(id string, newStatus string) (*models.Payment, error)
}
