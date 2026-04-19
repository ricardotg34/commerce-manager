package dtos

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type SuccessResponse struct {
	Success bool `json:"success" example:"true"`
	Data    any  `json:"data,omitempty"`
}

type FailResponse struct {
	Success bool       `json:"success" example:"true"`
	Error   *ErrorInfo `json:"error,omitempty"`
}

type ErrorInfo struct {
	Code    string `json:"code" example:"404"`
	Message string `json:"message" example:"not found"`
}

// OK sends a success response.
func OK(c *gin.Context, data any) {
	c.JSON(http.StatusOK, SuccessResponse{
		Success: true,
		Data:    data,
	})
}

// Fail sends an error response.
func Fail(c *gin.Context, status int, code, message string) {
	c.JSON(status, FailResponse{
		Success: false,
		Error:   &ErrorInfo{Code: code, Message: message},
	})
}
