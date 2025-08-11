package main

import (
	"fmt"
	"log"
	"os"
	"paymentMicroservice/config"
	"paymentMicroservice/handlers"
	"paymentMicroservice/repository"
	"paymentMicroservice/services"

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
	app := iris.New()

	db, err := config.ConnectToDB()
	if err != nil {
		fmt.Println("Connection Lost", err)
		return
	}

	// Dependency Injection
	repo := &repository.PaymentRepoImpl{DB: db}

	// ✅ Run subscriber in a goroutine
	go config.SubscribeToOrderEvents(repo)

	service := &services.PaymentService{Repo: repo}
	paymentHandler := &handlers.PaymentHandler{Service: *service}

	crs := cors.New(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowCredentials: true,
	})

	app.UseRouter(crs)

	app.Get("/health", func(ctx iris.Context) {
		ctx.StatusCode(iris.StatusOK)
		ctx.WriteString("OK")
	})

	app.Get("/payment", func(ctx iris.Context) {
		ctx.WriteString("Hello, Payment Service!")
	})

	app.Get("payment/:id", paymentHandler.GetPaymentById)

	fmt.Println("🚀 Payment Service running on http://localhost:8081")
	app.Listen(":8081")
}
