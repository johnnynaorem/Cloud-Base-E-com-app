package main

import (
	"log"
	"os"
	"productMicro/config"
	"productMicro/handlers"
	"productMicro/repository"
	"productMicro/router"
	"productMicro/services"

	"github.com/kataras/iris/v12"
)

func main() {
	// ✅ Load secrets first
	config.LoadSecretsFromGCP()

	secret := os.Getenv("SECRET_KEY")

	// ✅ DB connection
	db, err := config.ConnectToDB()
	if err != nil {
		log.Fatalf("❌ Connection Lost: %v", err)
	}

	// ✅ Dependency Injection
	repo := &repository.ProductRepoImpl{DB: db}
	service := &services.ProductService{Repo: repo}
	productHandler := &handlers.ProductHandler{Service: service}

	// ✅ Setup Iris app
	app := iris.New()
	router.Setup(app, productHandler, secret)
}
