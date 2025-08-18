package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"paymentMicroservice/config"
	"paymentMicroservice/handlers"
	"paymentMicroservice/repository"
	"paymentMicroservice/services"
	"strings"

	"github.com/iris-contrib/middleware/cors"
	"github.com/kataras/iris/v12"

	secretmanager "cloud.google.com/go/secretmanager/apiv1"
	secretmanagerpb "google.golang.org/genproto/googleapis/cloud/secretmanager/v1"
)

// ✅ Load secrets from Secret Manager into environment variables
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

	// ✅ Parse secrets into environment variables
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
	// ✅ Load secrets before starting the service
	loadSecretsFromGCP()

	app := iris.New()

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

	fmt.Println("🚀 Payment Service running on http://localhost:8081")
	app.Listen(":8081")
}
