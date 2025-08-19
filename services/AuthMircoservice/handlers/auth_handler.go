package handlers

import (
	"auth-micro/models"
	"auth-micro/services"
	"net/http"

	"github.com/kataras/iris/v12"
	"go.uber.org/zap"
)

type AuthHandler struct {
	Service *services.AuthService
	Logger  *zap.Logger
}

// Register new user
func (h *AuthHandler) Register(ctx iris.Context) {
	var user models.User
	if err := ctx.ReadJSON(&user); err != nil {
		ctx.StatusCode(http.StatusBadRequest)
		ctx.JSON(iris.Map{"message": "Invalid request"})
		return
	}

	token, err := h.Service.Register(&user)
	if err != nil {
		h.Logger.Warn("Failed to register user", zap.Error(err))
		ctx.StatusCode(http.StatusBadRequest)
		ctx.JSON(iris.Map{"message": err.Error()})
		return
	}

	ctx.JSON(iris.Map{"message": "User created successfully", "token": token})
}

// Authenticate user login
func (h *AuthHandler) Login(ctx iris.Context) {
	var user models.User
	if err := ctx.ReadJSON(&user); err != nil {
		ctx.StatusCode(http.StatusBadRequest)
		ctx.JSON(iris.Map{"message": "Invalid request"})
		return
	}

	token, err := h.Service.Authenticate(user.Email, user.Password)
	if err != nil {
		h.Logger.Warn("Authentication failed", zap.Error(err))
		ctx.StatusCode(http.StatusUnauthorized)
		ctx.JSON(iris.Map{"message": err.Error()})
		return
	}

	ctx.JSON(iris.Map{"message": "Login successful", "token": token})
}
