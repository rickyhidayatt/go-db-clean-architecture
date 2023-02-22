package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/rickyhidayatt/model"
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

func (pc *ProductController) CreateNewProduct(c *gin.Context) {
	var newProduct *model.Product
	err := c.ShouldBind(&newProduct) // Bind nge  pharse request dan memasukan otomatis ke struct jikaa di json sama dengan struct maka akan masuk otomatis
	if err != nil {
		c.JSON(400, gin.H{
			"message": err.Error(),
		})
	} else {
		err := pc.productUseCase.CreateNewProduct(newProduct) //kalo oke tambah inputan dari user ke newProduct. yg nantinya akan di pasing otomatis input ke DB
		if err != nil {
			c.JSON(500, gin.H{
				"message": err.Error(),
			})
		} else {
			c.JSON(200, gin.H{
				"message": "OK",
				"data":    newProduct,
			})
		}
	}
}

func NewProductController(router *gin.Engine, productUc usecase.ProductUseCase) *ProductController {
	newProductController := ProductController{
		productUseCase: productUc,
	}
	router.GET("/products", newProductController.GetAllProduct)
	router.POST("/product", newProductController.CreateNewProduct)
	return &newProductController
}
