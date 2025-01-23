package router

import (
	g "github.com/gin-gonic/gin"
)

func Initialize() {
	router := g.Default()

	initializeRoutes(router)

	router.Run(":8080")
}
