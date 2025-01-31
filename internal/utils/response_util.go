package utils

import (
	"library-app/internal/core/domain"

	"github.com/gin-gonic/gin"
)

func RespondWithSuccess(c *gin.Context, status int, message string, data interface{}) {
	response := domain.ApiResponse{
		Status:  status,
		Message: message,
		Data:    data,
	}
	c.JSON(status, response)
}

func RespondWithError(c *gin.Context, status int, message string, err error) {
	response := domain.ApiResponse{
		Status:  status,
		Message: message,
		Error:   err.Error(),
	}
	c.JSON(status, response)
}
