package services

import (
	"paymentMicroservice/models"
	"paymentMicroservice/repository"
)

type PaymentService struct {
	Repo repository.PaymentRepo
}

// func (ps *PaymentService) MakePayment(p *models.PaymentCreate) error {
// 	return ps.Repo.MakePayment(p)
// }

func (ps *PaymentService) GetPaymentById(id string) (*models.Payment, error) {
	return ps.Repo.GetPaymentById(id)
}

func (ps *PaymentService) UpdatePaymentStatus(id string, status string) (*models.Payment, error) {
	return ps.Repo.UpdatePaymentStatus(id, status)
}
