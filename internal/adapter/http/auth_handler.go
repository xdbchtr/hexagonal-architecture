package http

import (
	"library-app/internal/core/ports"
	"library-app/internal/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

type AuthHandler struct {
	authService ports.AuthService
}

func NewAuthHandler(authService ports.AuthService) *AuthHandler {
	return &AuthHandler{authService: authService}
}

func (h *AuthHandler) Register(c *gin.Context) {
	var request struct {
		Name     string `json:"name"`
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	if err := c.ShouldBindJSON(&request); err != nil {
		utils.RespondWithError(c, http.StatusBadRequest, "Invalid request payload", err)
		return
	}

	if err := h.authService.Register(request.Name, request.Email, request.Password); err != nil {
		utils.RespondWithError(c, http.StatusInternalServerError, "Failed to register user", err)
		return
	}

	utils.RespondWithSuccess(c, http.StatusCreated, "User registered successfully", nil)
}

func (h *AuthHandler) Login(c *gin.Context) {
	var request struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	if err := c.ShouldBindJSON(&request); err != nil {
		utils.RespondWithError(c, http.StatusBadRequest, "Invalid request payload", err)
		return
	}

	token, err := h.authService.Login(request.Email, request.Password)
	if err != nil {
		utils.RespondWithError(c, http.StatusUnauthorized, "Invalid credentials", err)
		return
	}

	utils.RespondWithSuccess(c, http.StatusOK, "Login successful", gin.H{"token": token})
}
