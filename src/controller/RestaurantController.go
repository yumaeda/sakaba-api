package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"sakaba.link/api/src/repository"
)

// RestaurantController is a controller for Restaurant API.
type RestaurantController struct{}

// GetOpenRestaurants returns open restaurants.
func (c *RestaurantController) GetOpenRestaurants(ctx *gin.Context) {
	restaurantRepository := repository.RestaurantRepository{}
	restaurants := restaurantRepository.GetOpenRestaurants()

	ctx.JSON(http.StatusOK, gin.H{
		"statusCode": 200,
		"body":       restaurants,
	})
}

// GetOpenRestaurantsByGenreID returns open restaurants for the specified genre.
func (c *RestaurantController) GetOpenRestaurantsByGenreID(ctx *gin.Context) {
	restaurantRepository := repository.RestaurantRepository{}
	restaurants := restaurantRepository.GetOpenRestaurantsByGenreID(ctx.Param("id"))

	ctx.JSON(http.StatusOK, gin.H{
		"statusCode": 200,
		"body":       restaurants,
	})
}

// GetOpenRestaurantsByDishID returns open restaurants for the specified dish.
func (c *RestaurantController) GetOpenRestaurantsByDishID(ctx *gin.Context) {
	restaurantRepository := repository.RestaurantRepository{}
	restaurants := restaurantRepository.GetOpenRestaurantsByDishID(ctx.Param("id"))

	ctx.JSON(http.StatusOK, gin.H{
		"statusCode": 200,
		"body":       restaurants,
	})
}

// GetOpenRestaurantCount returns number of open restaurants.
func (c *RestaurantController) GetOpenRestaurantCount(ctx *gin.Context) {
	restaurantRepository := repository.RestaurantRepository{}
	restaurantCount := restaurantRepository.GetOpenRestaurantCount()

	ctx.JSON(http.StatusOK, gin.H{
		"statusCode": 200,
		"body":       restaurantCount,
	})
}
