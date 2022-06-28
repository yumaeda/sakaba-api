package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/yumaeda/sakaba-api/src/repository"
)

// DrinkController is a controller for Drink API.
type DrinkController struct {
	Repository repository.DrinkRepository
}

// GetAllDrinks returns all the drinks.
func (c DrinkController) GetAllDrinks(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"statusCode": 200,
		"body":       c.Repository.GetAllDrinks(),
	})
}

// GetDrinkByID returns the specified drink.
func (c DrinkController) GetDrinkByID(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"statusCode": 200,
		"body":       c.Repository.GetDrinkByID(ctx.Param("id")),
	})
}
