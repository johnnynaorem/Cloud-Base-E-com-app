package main

import (
	"auth-micro/models"
	"net/http"
	"strings"

	"github.com/kataras/iris/v12"
	"go.uber.org/zap"
	"golang.org/x/crypto/bcrypt"
)

func AuthenticateUser(ctx iris.Context) {
	var user models.User
	ctx.ReadJSON(&user)

	userEmail := user.Email
	userPassword := user.Password

	logger.Info("Recieved Authenticate User Request", zap.String("user email", userEmail))

	// ? validation logic
	if userEmail == "" || userPassword == "" || !strings.Contains(userEmail, "@") || !strings.Contains(userEmail, ".") || len(userPassword) < 6 {
		logger.Warn("Invalid Request", zap.String("user email", userEmail))
		ctx.StatusCode(http.StatusBadRequest)
		ctx.JSON(iris.Map{"message": "The request contains missing or invalid fields"})
		return
	}

	// ? After fields are validated
	var existingUser *models.User

	userNotFoundError := dbConnector.Where("email = ?", user.Email).First(&existingUser).Error
	// ? If the user is already exist -> userNotFoundError = nil
	// ? If the user does not exist -> userNotFoundError = error

	if userNotFoundError != nil {
		logger.Warn("Authentication failed", zap.Error(userNotFoundError))
		ctx.StatusCode(http.StatusNotFound)
		ctx.JSON(iris.Map{"message": "User not found"})
		return
	}

	// ? config.ComparePassword(existingUser.Password, userPassword) is nil success passwords are matching.
	// ? config.ComparePassword(existingUser.Password, userPassword) is not nil failed passwords are not matching.
	//? compare the passwords
	err := bcrypt.CompareHashAndPassword([]byte(existingUser.Password), []byte(userPassword))

	if err != nil {
		logger.Warn("Authentication failed due to wrong password")
		ctx.StatusCode(http.StatusBadRequest)
		ctx.JSON(iris.Map{"message": "Authentication failed"})
		return
	}

	logger.Info("User Authenticated Successfully", zap.String("username", existingUser.Name), zap.String("useremail", userEmail))
	token, err := jwtManager.GeneratingToken(existingUser)
	if err != nil {
		ctx.StatusCode(http.StatusInternalServerError)
		ctx.JSON(iris.Map{"token failure": "Couldn't generate the token"})
		return
	}
	ctx.StatusCode(http.StatusOK)
	ctx.JSON(iris.Map{"message": "User login successfully", "token": token})
}
