package http

import (
	"library-app/internal/core/domain"
	"library-app/internal/core/ports"
	"library-app/internal/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

type UserHandler struct {
	userService ports.UserService
}

func NewUserHandler(userService ports.UserService) *UserHandler {
	return &UserHandler{userService: userService}
}

func (h *UserHandler) CreateUser(c *gin.Context) {
	var user domain.User
	if err := c.ShouldBindJSON(&user); err != nil {
		utils.RespondWithError(c, http.StatusBadRequest, "Invalid request payload", err)
		return
	}
	if err := h.userService.CreateUser(&user); err != nil {
		utils.RespondWithError(c, http.StatusInternalServerError, "Failed to create user", err)
		return
	}
	utils.RespondWithSuccess(c, http.StatusCreated, "User created successfully", user)
}

func (h *UserHandler) GetUser(c *gin.Context) {
	id := c.Param("id")
	user, err := h.userService.GetUser(id)
	if err != nil {
		utils.RespondWithError(c, http.StatusNotFound, "User not found", err)
		return
	}
	utils.RespondWithSuccess(c, http.StatusOK, "User retrieved successfully", user)
}

func (h *UserHandler) GetAllUsers(c *gin.Context) {
	users, err := h.userService.GetAllUsers()
	if err != nil {
		utils.RespondWithError(c, http.StatusInternalServerError, "Failed to retrieve users", err)
		return
	}
	utils.RespondWithSuccess(c, http.StatusOK, "Users retrieved successfully", users)
}

func (h *UserHandler) UpdateUser(c *gin.Context) {
	id := c.Param("id")
	var user domain.User
	if err := c.ShouldBindJSON(&user); err != nil {
		utils.RespondWithError(c, http.StatusBadRequest, "Invalid request payload", err)
		return
	}
	if err := h.userService.UpdateUser(id, &user); err != nil {
		utils.RespondWithError(c, http.StatusInternalServerError, "Failed to update user", err)
		return
	}
	utils.RespondWithSuccess(c, http.StatusOK, "User updated successfully", user)
}

func (h *UserHandler) DeleteUser(c *gin.Context) {
	id := c.Param("id")
	if err := h.userService.DeleteUser(id); err != nil {
		utils.RespondWithError(c, http.StatusInternalServerError, "Failed to delete user", err)
		return
	}
	utils.RespondWithSuccess(c, http.StatusOK, "User deleted successfully", nil)
}
