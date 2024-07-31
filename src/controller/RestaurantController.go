package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/yumaeda/sakaba-api/src/model"
	"github.com/yumaeda/sakaba-api/src/repository"
)

// RestaurantController is a controller for Restaurant API.
type RestaurantController struct {
	Repository repository.RestaurantRepository
}

// GetRestaurants returns all the restaurants.
func (c RestaurantController) GetRestaurants(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"statusCode": 200,
		"body":       c.Repository.GetRestaurants(),
	})
}

// GetRestaurantsByArea returns restaurants for the specified area.
func (c RestaurantController) GetRestaurantsByArea(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"statusCode": 200,
		"body":       c.Repository.GetRestaurantsByArea(ctx.Param("id")),
	})
}

// GetRestaurantsByGenreID returns open restaurants for the specified genre.
func (c RestaurantController) GetRestaurantsByGenreID(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"statusCode": 200,
		"body":       c.Repository.GetRestaurantsByGenreID(ctx.Param("id")),
	})
}

// GetRestaurantsByDrinkID returns open restaurants for the specified drink.
func (c RestaurantController) GetRestaurantsByDrinkID(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"statusCode": 200,
		"body":       c.Repository.GetRestaurantsByDrinkID(ctx.Param("id")),
	})
}

// GetRestaurantsByDishID returns open restaurants for the specified dish.
func (c RestaurantController) GetRestaurantsByDishID(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"statusCode": 200,
		"body":       c.Repository.GetRestaurantsByDishID(ctx.Param("id")),
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
		id, dbError := c.Repository.AddRestaurant(json.URL, json.Name, json.Genre, json.Tel, json.BusinessDayInfo, json.Address, json.Latitude, json.Longitude, json.Area)
		if dbError == nil {
			ctx.JSON(http.StatusOK, gin.H{
				"statusCode": 200,
				"id":         id,
			})
			return
		}
	}

	ctx.JSON(http.StatusBadRequest, gin.H{
		"statusCode": 400,
		"error":      "Restaurant insertion failed.",
	})
}
