package services

import (
	"productMicro/models"
	"productMicro/repository"
)

type ProductService struct {
	Repo repository.ProductRepository
}

func (ps *ProductService) GetAllProducts() ([]models.Product, error) {
	return ps.Repo.GetProducts()
}

func (ps *ProductService) GetProductByID(id string) (*models.Product, error) {
	return ps.Repo.GetProductById(id)
}

func (ps *ProductService) CreateProduct(product *models.Product) (*models.Product, error) {
	return ps.Repo.CreateProduct(product)
}

func (ps *ProductService) UpdateProduct(id string, product *models.Product) (*models.Product, error) {
	return ps.Repo.UpdateProduct(id, product)
}
