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
