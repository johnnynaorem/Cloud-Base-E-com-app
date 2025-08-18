package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"productMicro/config"
	"productMicro/handlers"
	"productMicro/jwt"
	"productMicro/repository"
	"productMicro/services"
	"strings"

	"github.com/iris-contrib/middleware/cors"
	"github.com/kataras/iris/v12"

	secretmanager "cloud.google.com/go/secretmanager/apiv1"
	secretmanagerpb "google.golang.org/genproto/googleapis/cloud/secretmanager/v1"
)

// ✅ Load secrets from GCP Secret Manager into environment variables
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

	// ✅ Parse secrets into env variables
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
	// ✅ Load secrets first
	loadSecretsFromGCP()

	secret := os.Getenv("SECRET_KEY")
	app := iris.New()

	// ✅ DB connection
	db, err := config.ConnectToDB()
	if err != nil {
		log.Fatalf("❌ Connection Lost: %v", err)
	}

	// ✅ Dependency Injection
	repo := &repository.ProductRepoImpl{DB: db}
	service := &services.ProductService{Repo: repo}
	productHandler := &handlers.ProductHandler{Service: service}

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

	// ✅ Routes
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
