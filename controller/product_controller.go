package controller

import (
	"net/http"
	"products/model"
	"products/usecase"
	"strconv"

	"github.com/gin-gonic/gin"
)

type productController struct {
	usecase usecase.ProductUseCase
}

func NewProdutController(u usecase.ProductUseCase) productController {
	return productController{
		usecase: u,
	}
}

func (p *productController) GetProducts(ctx *gin.Context) {
	products, err := p.usecase.GetProducts()

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err)
	}

	ctx.JSON(http.StatusOK, products)
}

func (p *productController) CreateProducts(ctx *gin.Context) {
	var product model.Product

	err := ctx.BindJSON(&product)

	if err != nil {
		ctx.JSON(http.StatusBadGateway, err)
		return
	}

	response, err := p.usecase.CreateProducts(product)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err)
		return
	}

	ctx.JSON(http.StatusCreated, response)
}

func (p *productController) GetProductById(ctx *gin.Context) {

	id := ctx.Param("id")

	if id == "" {
		r := model.Response{
			Message: "id not be null",
		}
		ctx.JSON(http.StatusBadRequest, r)
		return
	}

	productId, err := strconv.Atoi(id)

	if err != nil {
		r := model.Response{
			Message: "id must be a number",
		}
		ctx.JSON(http.StatusBadRequest, r)
		return
	}

	data, err := p.usecase.GetProductById(productId)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err)
		return
	}

	if data == nil {
		r := model.Response{
			Message: "product with id not found",
		}
		ctx.JSON(http.StatusNotFound, r)
		return
	}

	ctx.JSON(http.StatusCreated, data)

}
