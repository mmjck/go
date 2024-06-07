package routers

import (
	"authentication/handlers"
	"authentication/models"
	"log"

	"github.com/gin-gonic/gin"
)

type Login struct {
	logger       *log.Logger
	loginHandler *handlers.Login
	flags        *models.Flags
}

func (r *Login) RegisterRoutes() *gin.Engine {
	ginEngine := gin.Default()
	ginEngine.POST("/login", r.loginHandler.Login)

}
