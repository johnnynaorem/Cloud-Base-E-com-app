package main

import (
	"auth-micro/config"
	"auth-micro/jwt"
	"context"
	"log"
	"os"
	"strings"
	"time"

	"github.com/iris-contrib/middleware/cors"
	"github.com/kataras/iris/v12"

	"go.uber.org/zap"
	"gorm.io/gorm"

	secretmanager "cloud.google.com/go/secretmanager/apiv1"
	secretmanagerpb "google.golang.org/genproto/googleapis/cloud/secretmanager/v1"
)

var (
	logger      *zap.Logger
	jwtManager  *jwt.JWTManager
	dbConnector *gorm.DB
	err         error
)

func loadSecretsFromGCP() {
	ctx := context.Background()

	secretName := os.Getenv("SECRET_CREDENTIALS")
	if secretName == "" {
		log.Fatal("❌ SECRET_NAME environment variable not set")
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

func init() {
	// Load secrets directly from GCP Secret Manager
	loadSecretsFromGCP()

	// Setup logger
	var err error
	logger, err = zap.NewDevelopment()
	if err != nil {
		panic(err)
	}
	defer logger.Sync()
}

func main() {
	dbConnector, err = config.ConnectToDB()
	if err != nil {
		logger.Fatal("❌ Failed to connect to the database", zap.Error(err))
	}

	jwtManager = jwt.NewJWTManager(os.Getenv("SECRET_KEY"), 5*time.Hour)

	// Iris setup
	app := iris.New()

	crs := cors.New(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowCredentials: true,
	})

	app.UseRouter(crs)

	app.Get("/health", func(ctx iris.Context) {
		ctx.StatusCode(iris.StatusOK)
		ctx.WriteString("OK")
	})

	app.Post("/auth/save-user", AddUser)
	app.Post("/auth/login-user", AuthenticateUser)
	app.Get("/auth", func(ctx iris.Context) {
		ctx.StatusCode(iris.StatusOK)
		ctx.JSON(iris.Map{"message": "Welcome to the Auth Microservice!"})
	})

	app.Listen(":8083")
}
