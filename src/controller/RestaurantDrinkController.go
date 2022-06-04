package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"sakaba.link/api/src/model"
	"sakaba.link/api/src/repository"
)

// RestaurantDrinkController is a controller for Drink API.
type RestaurantDrinkController struct{}

// AddRestaurantDrink adds the specified genre to the specified restaurant.
func (c *RestaurantDrinkController) AddRestaurantDrink(ctx *gin.Context) {
	var json model.RestaurantDrink
	if err := ctx.ShouldBindJSON(&json); err == nil {
		restaurantDrinkRepository := repository.RestaurantDrinkRepository{}
		dbError := restaurantDrinkRepository.AddRestaurantDrink(json.RestaurantID, json.DrinkID)
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
