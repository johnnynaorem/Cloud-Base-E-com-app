package models

import "gorm.io/gorm"

type Product struct {
	gorm.Model
	ProductName     string  `json:"product_name"`
	ProductDesc     string  `json:"product_desc"`
	ProductPrice    float64 `json:"product_price"`
	ProductStock    int     `json:"product_stock"`
	ProductImage    string  `json:"product_image"`
	ProductCategory string  `json:"product_category"`
}
