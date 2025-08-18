package main

import (
	"context"
	"log"
	"orderPaymentMicroservice/config"
	"orderPaymentMicroservice/handlers"
	"orderPaymentMicroservice/jwt"
	"orderPaymentMicroservice/repository"
	"orderPaymentMicroservice/services"
	"os"
	"strings"

	"cloud.google.com/go/pubsub"
	"github.com/iris-contrib/middleware/cors"
	"github.com/kataras/iris/v12"

	secretmanager "cloud.google.com/go/secretmanager/apiv1"
	secretmanagerpb "google.golang.org/genproto/googleapis/cloud/secretmanager/v1"
)

func loadSecretsFromGCP() {
	ctx := context.Background()

	secretName := os.Getenv("SECRET_CREDENTIALS")
	if secretName == "" {
		log.Fatal("❌ SECRET_CREDENTIALS environment variable not set")
	}

	client, err := secretmanager.NewClient(ctx)
	if err != nil {
		log.Fatalf("❌ Failed to create secretmanager client: %v", err)
	}
	defer client.Close()

	req := &secretmanagerpb.AccessSecretVersionRequest{Name: secretName}
	result, err := client.AccessSecretVersion(ctx, req)
	if err != nil {
		log.Fatalf("❌ Failed to access secret version: %v", err)
	}

	secretData := string(result.Payload.Data)

	lines := strings.Split(secretData, "\n")
	for _, line := range lines {
		if strings.TrimSpace(line) == "" || strings.HasPrefix(line, "#") {
			continue
		}
		parts := strings.SplitN(line, "=", 2)
		if len(parts) == 2 {
			key := strings.TrimSpace(parts[0])
			value := strings.TrimSpace(parts[1])
			os.Setenv(key, value)
		}
	}
	log.Println("✅ Secrets loaded into environment variables")
}

func main() {
	// ✅ Load secrets before anything else
	loadSecretsFromGCP()

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

	// ✅ Routes
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

	// ✅ Start server
	app.Listen(":8080")
}
