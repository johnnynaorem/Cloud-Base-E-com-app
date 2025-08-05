package main

import (
	"auth-micro/config"
	"auth-micro/jwt"
	"log"
	"time"

	"github.com/iris-contrib/middleware/cors"
	"github.com/joho/godotenv"
	"github.com/kataras/iris/v12"

	"go.uber.org/zap"
	"gorm.io/gorm"
)

var logger *zap.Logger
var jwtManager *jwt.JWTManager
var dbConnector *gorm.DB
var err error

func init() {
	if err := godotenv.Load(); err != nil {
		log.Println("⚠️  No .env file found, using environment variables")
	}
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
		logger.Fatal("Failed to connect to the database", zap.Error(err))
	}

	// * Create a new jwt manager
	jwtManager = jwt.NewJWTManager("SECRET_KEY", 5*time.Hour)

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
