package handler

import (
	"github.com/gin-gonic/gin"
)

type errorResponse struct {
	Message string `json:"message"`
}

type statusResponce struct {
	Status string `json:"status"`
}

func newErrorResponse(c *gin.Context, statuscode int, message string) {
	c.AbortWithStatusJSON(statuscode, errorResponse{Message: message})
}
