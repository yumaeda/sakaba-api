package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"sakaba.link/api/src/model"
	"sakaba.link/api/src/repository"
)

// RestaurantController is a controller for Restaurant API.
type RestaurantController struct {
	Repository repository.RestaurantRepository
}

// GetOpenRestaurants returns open restaurants.
func (c RestaurantController) GetOpenRestaurants(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"statusCode": 200,
		"body":       c.Repository.GetOpenRestaurants(),
	})
}

// GetOpenRestaurantsByGenreID returns open restaurants for the specified genre.
func (c RestaurantController) GetOpenRestaurantsByGenreID(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"statusCode": 200,
		"body":       c.Repository.GetOpenRestaurantsByGenreID(ctx.Param("id")),
	})
}

// GetOpenRestaurantsByDrinkID returns open restaurants for the specified drink.
func (c RestaurantController) GetOpenRestaurantsByDrinkID(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"statusCode": 200,
		"body":       c.Repository.GetOpenRestaurantsByDrinkID(ctx.Param("id")),
	})
}

// GetOpenRestaurantsByDishID returns open restaurants for the specified dish.
func (c RestaurantController) GetOpenRestaurantsByDishID(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"statusCode": 200,
		"body":       c.Repository.GetOpenRestaurantsByDishID(ctx.Param("id")),
	})
}

// GetOpenRestaurantCount returns number of open restaurants.
func (c RestaurantController) GetOpenRestaurantCount(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"statusCode": 200,
		"body":       c.Repository.GetOpenRestaurantCount(),
	})
}

// AddRestaurant adds a new resgaurant.
func (c RestaurantController) AddRestaurant(ctx *gin.Context) {
	var json model.Restaurant
	if err := ctx.ShouldBindJSON(&json); err == nil {
		dbError := c.Repository.AddRestaurant(json.URL, json.Name, json.Genre, json.Tel, json.BusinessDayInfo, json.Address, json.Latitude, json.Longitude, json.Area)
		if dbError == nil {
			ctx.JSON(http.StatusOK, gin.H{
				"statusCode": 200,
				"body":       "New restaurant is inserted.",
			})
			return
		}
	}

	ctx.JSON(http.StatusBadRequest, gin.H{
		"statusCode": 400,
		"error":      "Restaurant insertion failed.",
	})
}
