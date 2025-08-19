package services

import (
	"auth-micro/jwt"
	"auth-micro/models"
	"auth-micro/repository"
	"errors"

	"golang.org/x/crypto/bcrypt"
)

type AuthService struct {
	Repo       repository.UserRepository
	JWTManager *jwt.JWTManager
}

// Register a new user
func (s *AuthService) Register(user *models.User) (string, error) {
	// Check if user exists
	existing, _ := s.Repo.Authenticate(user.Email)
	if existing != nil {
		return "", errors.New("user already exists")
	}

	// Hash password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	user.Password = string(hashedPassword)

	// Save user
	if err := s.Repo.Create(user); err != nil {
		return "", err
	}

	// Generate token
	return s.JWTManager.GeneratingToken(user)
}

// Authenticate user and generate token
func (s *AuthService) Authenticate(email, password string) (string, error) {
	user, err := s.Repo.Authenticate(email)
	if err != nil {
		return "", errors.New("user not found")
	}

	// Compare password
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
		return "", errors.New("invalid credentials")
	}

	// Generate token
	return s.JWTManager.GeneratingToken(user)
}
