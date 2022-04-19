package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"sakaba.link/api/src/model"
	"sakaba.link/api/src/repository"
)

// RestaurantGenreController is a controller for Genre API.
type RestaurantGenreController struct{}

// AddRestaurantGenre adds the specified genre to the specified restaurant.
func (c *RestaurantGenreController) AddRestaurantGenre(ctx *gin.Context) {
	var json model.RestaurantGenre
	if err := ctx.ShouldBindJSON(&json); err == nil {
		restaurantGenreRepository := repository.RestaurantGenreRepository{}
		result := restaurantGenreRepository.AddRestaurantGenre(json.RestaurantID, json.GenreID)
		if result.Error == nil {
			ctx.JSON(http.StatusOK, gin.H{
				"statusCode": 200,
				"body":       "New restaurant genre is inserted.",
			})
		}
	}

	ctx.JSON(http.StatusBadRequest, gin.H{
		"statusCode": 400,
		"error":      "Restaurant genre insertion failed.",
	})
}
