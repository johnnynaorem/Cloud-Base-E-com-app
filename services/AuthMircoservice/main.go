package main

import (
	"auth-micro/config"
	"auth-micro/handlers"
	"auth-micro/jwt"
	"auth-micro/repository"
	"auth-micro/router"
	"auth-micro/services"
	"os"
	"time"

	"github.com/kataras/iris/v12"
	"go.uber.org/zap"
)

func main() {
	// ✅ Load secrets
	config.LoadSecretsFromGCP()

	// ✅ Logger
	logger, _ := zap.NewDevelopment()
	defer logger.Sync()

	// ✅ DB connection
	db, err := config.ConnectToDB()
	if err != nil {
		logger.Fatal("❌ Failed to connect to DB", zap.Error(err))
	}

	// ✅ JWT manager
	jwtManager := jwt.NewJWTManager(os.Getenv("SECRET_KEY"), 5*time.Hour)

	// ✅ Setup dependencies
	repo := &repository.UserRepoImpl{DB: db}
	service := &services.AuthService{Repo: repo, JWTManager: jwtManager}
	authHandler := &handlers.AuthHandler{Service: service, Logger: logger}

	// ✅ Setup Iris
	app := iris.New()
	router.Setup(app, authHandler)

	// ✅ Start server
	app.Listen(":8083")
}
