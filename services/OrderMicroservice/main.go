package main

import (
	"context"
	"fmt"
	"log"
	"orderPaymentMicroservice/config"
	"orderPaymentMicroservice/handlers"
	"orderPaymentMicroservice/jwt"
	"orderPaymentMicroservice/repository"
	"orderPaymentMicroservice/services"
	"os"

	"cloud.google.com/go/pubsub"
	"github.com/iris-contrib/middleware/cors"
	"github.com/joho/godotenv"
	"github.com/kataras/iris/v12"
)

func init() {
	envPath := os.Getenv("DOTENV_CONFIG_PATH")
	if envPath == "" {
		envPath = ".env" // fallback for local dev
	}

	if err := godotenv.Load(envPath); err != nil {
		log.Println("⚠️  No .env file found, using environment variables")
	}
}
func main() {

	secret := os.Getenv("SECRET_KEY")
	app := iris.New()
	db, error := config.ConnectToDB()
	if error != nil {
		fmt.Printf("Connection Lost: %s", error)
		return
	}

	// ? Dependency Injection
	ctx := context.Background()
	pubsubClient, err := pubsub.NewClient(ctx, "johnny-projectt")
	if err != nil {
		fmt.Errorf("Something went wrong: %s", err)
	}

	repo := &repository.OrderRepoImpl{DB: db}
	config.ListenForPayments(repo)
	service := &services.OrderService{Repo: repo}
	orderHandler := &handlers.OrderHandler{Service: service, PubSubClient: pubsubClient,
		Ctx: ctx}

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

	app.Get("/order", func(ctx iris.Context) {
		ctx.WriteString("Hello, Order Microservice!")
	})
	app.Use(jwt.AuthMiddleware(secret))
	app.Post("/order/create", orderHandler.GenerateOrder)
	app.Get("/order/getorders", orderHandler.GetProducts)
	app.Get("/order/:id", orderHandler.GetOrderById)
	app.Patch("/order", orderHandler.UpdateOrderStatus)

	app.Listen(":8080")
}
