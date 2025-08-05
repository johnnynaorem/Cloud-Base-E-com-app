package services

import "productMicro/models"

type ProductServiceInterface interface {
	GetAllProducts() ([]models.Product, error)
	GetProductByID(id string) (*models.Product, error)
	CreateProduct(product *models.Product) (*models.Product, error)
	UpdateProduct(id string, product *models.Product) (*models.Product, error)
}
