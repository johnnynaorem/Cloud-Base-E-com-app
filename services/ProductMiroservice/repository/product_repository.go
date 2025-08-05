package repository

import "productMicro/models"

type ProductRepository interface {
	CreateProduct(p *models.Product) (*models.Product, error)
	GetProductById(id string) (*models.Product, error)
	GetProducts() ([]models.Product, error)
	UpdateProduct(id string, p *models.Product) (*models.Product, error)
	DeleteProduct(id string) error
}
