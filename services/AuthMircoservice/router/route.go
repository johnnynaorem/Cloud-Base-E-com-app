package router

import (
	"auth-micro/handlers"

	"github.com/iris-contrib/middleware/cors"
	"github.com/kataras/iris/v12"
)

func Setup(app *iris.Application, authHandler *handlers.AuthHandler) {
	// ✅ CORS
	crs := cors.New(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowCredentials: true,
	})
	app.UseRouter(crs)

	// ✅ Health
	app.Get("/health", func(ctx iris.Context) {
		ctx.StatusCode(iris.StatusOK)
		ctx.WriteString("OK")
	})

	// ✅ Auth routes
	app.Post("/auth/save-user", authHandler.Register)
	app.Post("/auth/login-user", authHandler.Login)
	app.Get("/auth", func(ctx iris.Context) {
		ctx.JSON(iris.Map{"message": "Welcome to the Auth Microservice!"})
	})
}
