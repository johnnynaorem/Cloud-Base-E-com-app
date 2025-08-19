package router

import (
	"productMicro/handlers"
	"productMicro/jwt"

	"github.com/iris-contrib/middleware/cors"
	"github.com/kataras/iris/v12"
)

func Setup(app *iris.Application, productHandler *handlers.ProductHandler, secret string) {
	// ✅ CORS
	crs := cors.New(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Authorization", "Content-Type"},
		AllowCredentials: true,
	})
	app.UseRouter(crs)

	// ✅ Handle preflight
	app.Options("/{any:path}", func(ctx iris.Context) {
		ctx.StatusCode(iris.StatusNoContent)
	})

	// ✅ Health check
	app.Get("/health", func(ctx iris.Context) {
		ctx.StatusCode(iris.StatusOK)
		ctx.WriteString("OK")
	})

	// ✅ Public route
	app.Get("/product", func(ctx iris.Context) {
		ctx.WriteString("Hello, Product Service!")
	})

	// ✅ Protected routes
	app.Use(jwt.AuthMiddleware(secret))
	app.Post("/product/create", productHandler.CreateProduct)
	app.Get("/product/getproducts", productHandler.GetAllProducts)
	app.Get("/product/:id", productHandler.GetProductById)
	app.Put("/product/:id", productHandler.UpdateProduct)
}
