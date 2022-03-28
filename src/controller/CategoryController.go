package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"sakaba.link/api/src/repository"
)

// CategoryController is a controller for Category API.
type CategoryController struct{}

// GetCategoriesByRestaurantID returns categories for the specified restaurant.
func (c *CategoryController) GetCategoriesByRestaurantID(ctx *gin.Context) {
	categoyRepository := repository.CategoryRepository{}
	categories := categoyRepository.GetCategories(ctx.Param("id"))

	ctx.JSON(http.StatusOK, gin.H{
		"statusCode": 200,
		"body":       categories,
	})
}
