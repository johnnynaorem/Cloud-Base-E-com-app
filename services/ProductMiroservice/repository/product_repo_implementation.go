package repository

import (
	"productMicro/models"

	"gorm.io/gorm"
)

type ProductRepoImpl struct {
	DB *gorm.DB
}

func NewOrderRepoImpl(db *gorm.DB) ProductRepository {
	return &ProductRepoImpl{
		DB: db,
	}
}

func (p *ProductRepoImpl) CreateProduct(product *models.Product) (*models.Product, error) {
	if err := p.DB.Create(product).Error; err != nil {
		return nil, err
	}
	return product, nil
}

func (p *ProductRepoImpl) GetProductById(id string) (*models.Product, error) {
	var product models.Product
	err := p.DB.Where("id = ?", id).Find(&product).Error
	if err != nil {
		return nil, err
	}
	return &product, nil
}

func (p *ProductRepoImpl) GetProducts() ([]models.Product, error) {
	var products []models.Product
	err := p.DB.Find(&products).Error
	if err != nil {
		return nil, err
	}
	return products, nil
}

func (p *ProductRepoImpl) UpdateProduct(id string, product *models.Product) (*models.Product, error) {
	var updateProduct models.Product
	err := p.DB.Model(&updateProduct).Where("id = ?", id).Updates(product).Error
	if err != nil {
		return nil, err
	}
	return &updateProduct, nil
}

func (p *ProductRepoImpl) DeleteProduct(id string) error {
	var product models.Product
	err := p.DB.Where("id = ?", id).Delete(&product).Error
	if err != nil {
		return err
	}
	return nil
}
