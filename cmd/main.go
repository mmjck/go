package main

import (
	"products/controller"
	"products/db"
	"products/repository"
	"products/usecase"

	"github.com/gin-gonic/gin"
)

func main() {
	server := gin.Default()
	dbConnection, err := db.ConnectDb()

	if err != nil {
		panic(err)
	}

	ProductRepository := repository.NewProductRepository(dbConnection)

	ProductUseCase := usecase.NewProdutUseCase(ProductRepository)
	ProductController := controller.NewProdutController(ProductUseCase)

	server.GET("/ping", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{
			"message": "pong",
		})
	})

	server.GET("products", ProductController.GetProducts)
	server.GET("products/:id", ProductController.GetProductById)

	server.POST("products", ProductController.CreateProducts)

	server.Run(":8000")
}
