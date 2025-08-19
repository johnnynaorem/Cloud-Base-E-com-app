package main

import (
	"fmt"
	"log"
	"paymentMicroservice/config"
	"paymentMicroservice/handlers"
	"paymentMicroservice/repository"
	"paymentMicroservice/router"
	"paymentMicroservice/services"

	"github.com/kataras/iris/v12"
)

func main() {
	// ✅ Load secrets before starting the service
	config.LoadSecretsFromGCP()

	// ✅ DB connection
	db, err := config.ConnectToDB()
	if err != nil {
		log.Fatalf("❌ Connection Lost: %v", err)
	}

	// ✅ Dependency Injection
	repo := &repository.PaymentRepoImpl{DB: db}

	// Run subscriber in a goroutine
	go config.SubscribeToOrderEvents(repo)

	service := &services.PaymentService{Repo: repo}
	paymentHandler := &handlers.PaymentHandler{Service: *service}

	app := iris.New()

	// ✅ Setup routes
	router.Setup(app, paymentHandler)

	// ✅ Start server
	fmt.Println("🚀 Payment Service running on http://localhost:8081")
	app.Listen(":8081")
}
