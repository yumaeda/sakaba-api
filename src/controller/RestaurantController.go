package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"sakaba.link/api/src/repository"
)

type RestaurantController struct{}

func (c *RestaurantController) GetOpenRestaurants(ctx *gin.Context) {
	restaurantRepository := repository.RestaurantRepository{}
	restaurants := restaurantRepository.GetOpenRestaurants()

	ctx.JSON(http.StatusOK, gin.H{
		"statusCode": 200,
		"body":       restaurants,
	})
}

func (c *RestaurantController) GetOpenRestaurantsByGenreId(ctx *gin.Context) {
	restaurantRepository := repository.RestaurantRepository{}
	restaurants := restaurantRepository.GetOpenRestaurantsByGenreId(ctx.Param("id"))

	ctx.JSON(http.StatusOK, gin.H{
		"statusCode": 200,
		"body":       restaurants,
	})
}

func (c *RestaurantController) GetOpenRestaurantsByDishId(ctx *gin.Context) {
	restaurantRepository := repository.RestaurantRepository{}
	restaurants := restaurantRepository.GetOpenRestaurantsByDishId(ctx.Param("id"))

	ctx.JSON(http.StatusOK, gin.H{
		"statusCode": 200,
		"body":       restaurants,
	})
}

func (c *RestaurantController) GetOpenRestaurantCount(ctx *gin.Context) {
	restaurantRepository := repository.RestaurantRepository{}
	restaurantCount := restaurantRepository.GetOpenRestaurantCount()

	ctx.JSON(http.StatusOK, gin.H{
		"statusCode": 200,
		"body":       restaurantCount,
	})
}
