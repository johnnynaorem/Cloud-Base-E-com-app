package handlers

import (
	"context"
	"fmt"
	"orderPaymentMicroservice/config"
	"orderPaymentMicroservice/models"
	"orderPaymentMicroservice/services"

	"cloud.google.com/go/pubsub"
	"github.com/kataras/iris/v12"
)

type OrderHandler struct {
	Service      services.OrderServiceInterface
	PubSubClient *pubsub.Client
	Ctx          context.Context
}

func (h *OrderHandler) GenerateOrder(ctx iris.Context) {
	if h.Service == nil {
		ctx.StatusCode(iris.StatusInternalServerError)
		ctx.WriteString("Service is nil in handler")
		return
	}

	var orderCreateDTO models.OrderCreateDTO
	if err := ctx.ReadJSON(&orderCreateDTO); err != nil {
		ctx.StatusCode(iris.StatusBadRequest)
		ctx.WriteString("Invalid input")
		return
	}
	result, err := h.Service.GenerateOrder(&orderCreateDTO)
	if err != nil {
		ctx.StatusCode(iris.StatusInternalServerError)
		ctx.JSON(iris.Map{"error": "Internal server error"})
		return
	}

	event := &models.OrderEvent{
		OrderID:     result.OrderId,
		TotalAmount: result.TotalAmount,
	}

	config.PublishOrderCreated(h.Ctx, h.PubSubClient, *event)
	ctx.StatusCode(iris.StatusCreated)
	ctx.JSON(iris.Map{"message": "Order created successfully", "result": result})
}

func (o *OrderHandler) GetProducts(ctx iris.Context) {
	result, err := o.Service.GetOrders()
	if err != nil {
		ctx.StatusCode(iris.StatusInternalServerError)
		ctx.JSON(iris.Map{"error": "Internal server error"})
	}

	if result == nil {
		ctx.StatusCode(iris.StatusNotFound)
		ctx.JSON(iris.Map{"message": "No orders found"})
		return
	}
	ctx.StatusCode(iris.StatusOK)
	ctx.JSON(iris.Map{"message": "Orders retrieved successfully", "orders": result})
}

func (o *OrderHandler) GetOrderById(ctx iris.Context) {
	id := ctx.Params().Get("id")

	result, err := o.Service.GetOrderById(id)
	if err != nil {
		ctx.StatusCode(iris.StatusInternalServerError)
		ctx.JSON(iris.Map{"error": "Internal server error"})
		return
	}

	if result.OrderId == "" {
		ctx.StatusCode(iris.StatusNotFound)
		ctx.JSON(iris.Map{"message": fmt.Sprintf("Order not found with ORDER ID: %s", id)})
		return
	}
	ctx.StatusCode(iris.StatusOK)
	ctx.JSON(iris.Map{"message": "Order retrieved successfully", "order": result})
}

func (o *OrderHandler) UpdateOrderStatus(ctx iris.Context) {
	var orderStatusUpdateDTO models.OrderStatusUpdateDTO
	if err := ctx.ReadJSON(&orderStatusUpdateDTO); err != nil {
		ctx.StatusCode(iris.StatusBadRequest)
		ctx.JSON(iris.Map{"error": "Invalid input"})
		return
	}
	result, err := o.Service.UpdateOrderStatus(orderStatusUpdateDTO.OrderId, string(orderStatusUpdateDTO.Status))
	if err != nil {
		ctx.StatusCode(iris.StatusInternalServerError)
		ctx.JSON(iris.Map{"error": "Internal server error"})
		return
	}

	if result == nil {
		ctx.StatusCode(iris.StatusNotFound)
		ctx.JSON(iris.Map{"message": "Order not found"})
		return
	}
	ctx.StatusCode(iris.StatusOK)
	ctx.JSON(iris.Map{"message": "Order status updated successfully to " + result.Status})
}
