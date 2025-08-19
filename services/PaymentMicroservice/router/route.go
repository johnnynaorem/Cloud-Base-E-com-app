package router

import (
	"paymentMicroservice/handlers"

	"github.com/iris-contrib/middleware/cors"
	"github.com/kataras/iris/v12"
)

func Setup(app *iris.Application, paymentHandler *handlers.PaymentHandler) {
	// ✅ CORS
	crs := cors.New(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowCredentials: true,
	})
	app.UseRouter(crs)

	// ✅ Health check
	app.Get("/health", func(ctx iris.Context) {
		ctx.StatusCode(iris.StatusOK)
		ctx.WriteString("OK")
	})

	// ✅ Routes
	app.Get("/payment", func(ctx iris.Context) {
		ctx.WriteString("Hello, Payment Service!")
	})
	app.Get("/payment/:id", paymentHandler.GetPaymentById)
}
