package router

import (
	g "github.com/gin-gonic/gin"
)

func Initialize() {
	router := g.Default()

	router.GET("/ping", func(c *g.Context) {
		c.JSON(200, g.H{
			"message": "pong",
		})
	})

	router.Run(":8080")
}
