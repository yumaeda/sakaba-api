package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"sakaba.link/api/src/repository"
)

type CategoryController struct{}

func (c *CategoryController) GetCategoriesByRestaurantId(ctx *gin.Context) {
	categoyRepository := repository.CategoryRepository{}
	categories := categoyRepository.GetCategories(ctx.Param("id"))

	ctx.JSON(http.StatusOK, gin.H{
		"statusCode": 200,
		"body":       categories,
	})
}
