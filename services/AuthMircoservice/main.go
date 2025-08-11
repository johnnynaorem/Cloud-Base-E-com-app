package main

import (
	"auth-micro/config"
	"auth-micro/jwt"
	"log"
	"os"
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
	envPath := os.Getenv("DOTENV_CONFIG_PATH")
	if envPath == "" {
		envPath = ".env" // fallback for local dev
	}

	if err := godotenv.Load(envPath); err != nil {
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
	jwtManager = jwt.NewJWTManager(os.Getenv("SECRET KEY"), 5*time.Hour)

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
