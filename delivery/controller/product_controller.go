package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/rickyhidayatt/usecase"
)

type ProductController struct {
	productUseCase usecase.ProductUseCase
}

func (pc *ProductController) GetAllProduct(c *gin.Context) {
	products, err := pc.productUseCase.GetAllProduct()
	if err != nil {
		c.JSON(500, gin.H{
			"message": err.Error(),
		})
	} else {
		c.JSON(200, gin.H{
			"message": "OK",
			"data":    products,
		})
	}
}

func NewProductController(router *gin.Engine, productUc usecase.ProductUseCase) *ProductController {
	newProductController := ProductController{
		productUseCase: productUc,
	}
	router.GET("/products", newProductController.GetAllProduct)
	return &newProductController
}
