package service

import (
	"errors"
	"library-app/internal/core/domain"
	"library-app/internal/core/ports"
	"library-app/internal/utils"

	"golang.org/x/crypto/bcrypt"
)

type authService struct {
	repo ports.AuthRepository
}

func NewAuthService(repo ports.AuthRepository) ports.AuthService {
	return &authService{repo: repo}
}

func (s *authService) Register(name, email, password string) error {
	// Hash the password before storing it
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	user := &domain.User{
		Name:     name,
		Email:    email,
		Password: string(hashedPassword),
	}

	return s.repo.CreateUser(user)
}

func (s *authService) Login(email, password string) (string, error) {
	// Find the user by email
	user, err := s.repo.FindUserByEmail(email)
	if err != nil {
		return "", errors.New("user not found")
	}

	// Compare the provided password with the stored hash
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
		return "", errors.New("invalid credentials")
	}

	// Generate a JWT token (we'll implement this next)
	token, err := utils.GenerateJWT(user.ID)
	if err != nil {
		return "", errors.New("failed to generate token")
	}

	return token, nil
}
