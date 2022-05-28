package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"sakaba.link/api/src/repository"
)

// DrinkController is a controller for Drink API.
type DrinkController struct{}

// GetAllDrinks returns all the drinks.
func (c *DrinkController) GetAllDrinks(ctx *gin.Context) {
	drinkRepository := repository.DrinkRepository{}
	allDrinks := drinkRepository.GetAllDrinks()

	ctx.JSON(http.StatusOK, gin.H{
		"statusCode": 200,
		"body":       allDrinks,
	})
}

// GetDrinkByID returns the specified drink.
func (c *DrinkController) GetDrinkByID(ctx *gin.Context) {
	id := ctx.Param("id")
	drinkRepository := repository.DrinkRepository{}
	drink := drinkRepository.GetDrinkByID(id)

	ctx.JSON(http.StatusOK, gin.H{
		"statusCode": 200,
		"body":       drink,
	})
}
