package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"sakaba.link/api/src/repository"
)

// DishController is a controller for Dish API.
type DishController struct {
	Repository repository.DishRepository
}

// GetAllDishes returns all the dishes.
func (c DishController) GetAllDishes(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"statusCode": 200,
		"body":       c.Repository.GetAllDishes(),
	})
}

// GetDishByID returns the specified dish.
func (c DishController) GetDishByID(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"statusCode": 200,
		"body":       c.Repository.GetDishByID(ctx.Param("id")),
	})
}
