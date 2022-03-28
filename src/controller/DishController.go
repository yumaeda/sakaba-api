package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"sakaba.link/api/src/repository"
)

// DishController is a controller for Dish API.
type DishController struct{}

// GetAllDishes returns all the dishes.
func (c *DishController) GetAllDishes(ctx *gin.Context) {
	dishRepository := repository.DishRepository{}
	allDishes := dishRepository.GetAllDishes()

	ctx.JSON(http.StatusOK, gin.H{
		"statusCode": 200,
		"body":       allDishes,
	})
}

// GetDishByID returns the specified dish.
func (c *DishController) GetDishByID(ctx *gin.Context) {
	dishRepository := repository.DishRepository{}
	dish := dishRepository.GetDishByID(ctx.Param("id"))

	ctx.JSON(http.StatusOK, gin.H{
		"statusCode": 200,
		"body":       dish,
	})
}
