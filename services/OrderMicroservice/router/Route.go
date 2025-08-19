package router

import (
	"orderPaymentMicroservice/handlers"
	"orderPaymentMicroservice/jwt"

	"github.com/iris-contrib/middleware/cors"
	"github.com/kataras/iris/v12"
)

func Setup(app *iris.Application, orderHandler *handlers.OrderHandler, secret string) {
	// ✅ CORS
	crs := cors.New(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Authorization", "Content-Type"},
		AllowCredentials: true,
	})
	app.UseRouter(crs)

	// ✅ Preflight
	app.Options("/{any:path}", func(ctx iris.Context) {
		ctx.StatusCode(iris.StatusNoContent)
	})

	// ✅ Public Routes
	app.Get("/health", func(ctx iris.Context) {
		ctx.StatusCode(iris.StatusOK)
		ctx.WriteString("OK")
	})

	app.Get("/order", func(ctx iris.Context) {
		ctx.WriteString("Hello, Order Microservice!")
	})

	// ✅ Protected Routes
	app.Use(jwt.AuthMiddleware(secret))
	app.Post("/order/create", orderHandler.GenerateOrder)
	app.Get("/order/getorders", orderHandler.GetProducts)
	app.Get("/order/:id", orderHandler.GetOrderById)
	app.Patch("/order", orderHandler.UpdateOrderStatus)
}
