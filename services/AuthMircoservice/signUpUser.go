package main

import (
	"auth-micro/models"
	"fmt"
	"net/http"
	"regexp"

	"github.com/kataras/iris/v12"
	"go.uber.org/zap"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

func validateUser(user models.User) error {
	// Email format validation
	if !isValidEmail(user.Email) {
		return fmt.Errorf("invalid email format")
	}

	// Password length validation
	if len(user.Password) < 8 {
		return fmt.Errorf("password must be at least 8 characters long")
	}

	// Phone number format validation (assuming it should be numeric)
	if !isValidPhone(user.Phone) {
		return fmt.Errorf("invalid phone number format")
	}

	// Address and City check
	// if user.Address == "" || user.City == "" {
	// 	return fmt.Errorf("Address and City are required")
	// }

	return nil
}

// Helper function to validate email format using regex
func isValidEmail(email string) bool {
	re := regexp.MustCompile(`^[a-z0-9._%+-]+@[a-z0-9.-]+\.[a-z]{2,}$`)
	return re.MatchString(email)
}

// Helper function to validate phone number format
func isValidPhone(phone string) bool {
	re := regexp.MustCompile(`^\d{10}$`)
	return re.MatchString(phone)
}

func AddUser(ctx iris.Context) {
	var user models.User
	ctx.ReadJSON(&user)

	logger.Info("Recieved User Request", zap.String("useremail", user.Email), zap.String("username", user.Name))

	// ?1. Write the validation Logic

	if err := validateUser(user); err != nil {
		ctx.StatusCode(http.StatusBadRequest)
		ctx.JSON(iris.Map{"message": err.Error()})
		return
	}

	// ?2. Check for Existing user.
	var existingUser models.User

	userNotFoundError := dbConnector.Where("email = ?", user.Email).First(&existingUser).Error

	if userNotFoundError == gorm.ErrRecordNotFound {
		// ?3. hash password
		hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)

		if err != nil {
			logger.Error("Fail to Hashing Password", zap.String("message", err.Error()))
		}

		newUser := &models.User{Name: user.Name, Email: user.Email, Password: string(hashedPassword), Phone: user.Phone}

		primaryKey := dbConnector.Create(newUser)

		if primaryKey.Error != nil {
			logger.Error("Failed to Create user", zap.String("userPhone ", user.Phone), zap.Error(primaryKey.Error))
			ctx.StatusCode(http.StatusConflict)
			ctx.JSON(iris.Map{"message": "The Phone is already registered"})
			return
		}
		logger.Info(fmt.Sprintf("User %s created successfully", user.Name))
		token, err := jwtManager.GeneratingToken(newUser)
		if err != nil {
			ctx.StatusCode(http.StatusInternalServerError)
			ctx.JSON(iris.Map{"token failure": "Couldn't generate the token"})
		}
		ctx.JSON(iris.Map{"message": "User created successfully", "token": token})

	} else {
		logger.Warn("User Email Already Exist", zap.String("usermail", user.Email))
		ctx.StatusCode(http.StatusConflict)
		ctx.JSON(iris.Map{"message": "User Email Already Exist"})
	}

}
