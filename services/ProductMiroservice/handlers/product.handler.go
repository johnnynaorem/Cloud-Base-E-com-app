package handlers

import (
	"fmt"
	"productMicro/models"
	"productMicro/services"

	"github.com/kataras/iris/v12"
)

type ProductHandler struct {
	Service services.ProductServiceInterface
}

func (ph *ProductHandler) GetProductById(ctx iris.Context) {
	id := ctx.Params().Get("id")
	if id == "" {
		ctx.StatusCode(iris.StatusBadRequest)
		ctx.WriteString("Product ID is required")
		return
	}

	product, err := ph.Service.GetProductByID(id)
	if err != nil {
		ctx.StatusCode(iris.StatusInternalServerError)
		ctx.WriteString(fmt.Sprintf("Error retrieving product: %s", err.Error()))
		return
	}
	if product.ID == 0 {
		ctx.StatusCode(iris.StatusNotFound)
		ctx.WriteString("Product not found")
		return
	}
	ctx.StatusCode(iris.StatusOK)
	ctx.JSON(product)
}

func (ph *ProductHandler) GetAllProducts(ctx iris.Context) {
	products, err := ph.Service.GetAllProducts()
	if err != nil {
		ctx.StatusCode(iris.StatusInternalServerError)
		ctx.WriteString(fmt.Sprintf("Error retrieving products: %s", err.Error()))
		return
	}
	if len(products) == 0 {
		ctx.StatusCode(iris.StatusNotFound)
		ctx.WriteString("No products found")
		return
	}
	ctx.StatusCode(iris.StatusOK)
	ctx.JSON(products)
}

func (ph *ProductHandler) CreateProduct(ctx iris.Context) {
	var product models.Product
	if err := ctx.ReadJSON(&product); err != nil {
		ctx.StatusCode(iris.StatusBadRequest)
		ctx.WriteString("Invalid product data")
		return
	}

	createdProduct, err := ph.Service.CreateProduct(&product)
	if err != nil {
		ctx.StatusCode(iris.StatusInternalServerError)
		ctx.WriteString(fmt.Sprintf("Error creating product: %s", err.Error()))
		return
	}
	ctx.StatusCode(iris.StatusCreated)
	ctx.JSON(createdProduct)
}

func (ph *ProductHandler) UpdateProduct(ctx iris.Context) {
	var product models.Product
	id := ctx.Params().Get("id")
	if err := ctx.ReadJSON(&product); err != nil {
		ctx.StatusCode(iris.StatusBadRequest)
		ctx.WriteString("Invalid product data")
		return
	}
	updatedProduct, err := ph.Service.UpdateProduct(id, &product)
	if err != nil {
		ctx.StatusCode(iris.StatusInternalServerError)
		ctx.WriteString(fmt.Sprintf("Error updating product: %s", err.Error()))
		return
	}
	if updatedProduct == nil {
		ctx.StatusCode(iris.StatusNotFound)
		ctx.WriteString("Product not found")
		return
	}
	ctx.StatusCode(iris.StatusOK)
	ctx.JSON(updatedProduct)
}
