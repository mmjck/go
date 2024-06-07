package handlers

import (
	"authentication/models"

	"github.com/gin-gonic/gin"
)

func ok(context *gin.Context, status int, message string, data interface{}) {
	context.AbortWithStatusJSON(status, models.Response{
		Data:    data,
		Status:  status,
		Message: message,
	})
}

func badRequest(context *gin.Context, status int, message string, erros []models.ErrorDetail) {
	context.AbortWithStatusJSON(status, models.Response{
		Error:   erros,
		Status:  status,
		Message: message,
	})
}

//
