package routers

import (
	"authentication/handlers"
	"authentication/models"
	"authentication/token"
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
	ginEngine.POST("/validate-token", r.loginHandler.VerifyToken)

	return ginEngine
}

func NewRoute(l *log.Logger, f *models.Flags) *Login {
	loginHandler := handlers.NewLogin(l, f)

	token.Init()
	return &Login{
		logger:       l,
		loginHandler: loginHandler,
		flags:        f,
	}
}
