package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/yumaeda/sakaba-api/src/model"
	"github.com/yumaeda/sakaba-api/src/repository"
)

// RestaurantDrinkController is a controller for Drink API.
type RestaurantDrinkController struct {
	Repository repository.RestaurantDrinkRepository
}

// AddRestaurantDrink adds the specified genre to the specified restaurant.
func (c RestaurantDrinkController) AddRestaurantDrink(ctx *gin.Context) {
	var json model.RestaurantDrink
	if err := ctx.ShouldBindJSON(&json); err == nil {
		dbError := c.Repository.AddRestaurantDrink(json.RestaurantID, json.DrinkID)
		if dbError == nil {
			ctx.JSON(http.StatusOK, gin.H{
				"statusCode": 200,
				"body":       "New restaurant drink is inserted.",
			})
			return
		}
	}

	ctx.JSON(http.StatusBadRequest, gin.H{
		"statusCode": 400,
		"error":      "Restaurant drink insertion failed.",
	})
}
