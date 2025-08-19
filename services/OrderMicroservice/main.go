package main

import (
	"context"
	"log"
	router "orderPaymentMicroservice/Router"
	"orderPaymentMicroservice/config"
	"orderPaymentMicroservice/handlers"
	"orderPaymentMicroservice/repository"
	"orderPaymentMicroservice/services"
	"os"

	"cloud.google.com/go/pubsub"
	"github.com/kataras/iris/v12"
)

func main() {
	// ✅ Load secrets before anything else
	config.LoadSecretsFromGCP()

	// ✅ DB connection
	db, err := config.ConnectToDB()
	if err != nil {
		log.Fatalf("❌ Connection Lost: %s", err)
	}

	// ✅ JWT secret
	secret := os.Getenv("SECRET_KEY")

	// ✅ Pub/Sub client
	ctx := context.Background()
	projectID := os.Getenv("GCP_PROJECT_ID") // keep project configurable
	if projectID == "" {
		projectID = "johnny-projectt" // fallback
	}
	pubsubClient, err := pubsub.NewClient(ctx, projectID)
	if err != nil {
		log.Fatalf("❌ Failed to create Pub/Sub client: %s", err)
	}

	// ✅ Dependencies
	repo := &repository.OrderRepoImpl{DB: db}
	config.ListenForPayments(repo)
	service := &services.OrderService{Repo: repo}
	orderHandler := &handlers.OrderHandler{
		Service:      service,
		PubSubClient: pubsubClient,
		Ctx:          ctx,
	}

	// ✅ Setup Iris
	app := iris.New()
	router.Setup(app, orderHandler, secret)

	// ✅ Start server
	app.Listen(":8080")
}
