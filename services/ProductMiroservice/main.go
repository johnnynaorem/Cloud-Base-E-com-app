package main

import (
	"fmt"
	"log"
	"os"
	"productMicro/config"
	"productMicro/handlers"
	"productMicro/jwt"
	"productMicro/repository"
	"productMicro/services"

	"github.com/iris-contrib/middleware/cors"
	"github.com/joho/godotenv"
	"github.com/kataras/iris/v12"
)

func init() {
	if err := godotenv.Load(); err != nil {
		log.Println("⚠️  No .env file found, using environment variables")
	}
}

func main() {

	secret := os.Getenv("SECRET_KEY")
	app := iris.New()

	db, err := config.ConnectToDB()
	if err != nil {
		fmt.Println("Connection Lost", err)
		return
	}

	// Dependency Injection
	repo := &repository.ProductRepoImpl{DB: db}

	service := &services.ProductService{Repo: repo}
	productHandler := &handlers.ProductHandler{Service: service}

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

	app.Get("/health", func(ctx iris.Context) {
		ctx.StatusCode(iris.StatusOK)
		ctx.WriteString("OK")
	})

	app.Get("/product", func(ctx iris.Context) {
		ctx.WriteString("Hello, Product Service!")
	})

	app.Use(jwt.AuthMiddleware(secret))

	app.Post("/product/create", productHandler.CreateProduct)
	app.Get("/product/getproducts", productHandler.GetAllProducts)

	app.Get("/product/:id", productHandler.GetProductById)
	app.Put("/product/:id", productHandler.UpdateProduct)

	fmt.Println("🚀 Product Service running on http://localhost:8082")
	app.Listen(":8082")
}
