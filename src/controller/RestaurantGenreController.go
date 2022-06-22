package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"sakaba.link/api/src/model"
	"sakaba.link/api/src/repository"
)

// RestaurantGenreController is a controller for Genre API.
type RestaurantGenreController struct {
	Repository repository.RestaurantGenreRepository
}

// AddRestaurantGenre adds the specified genre to the specified restaurant.
func (c RestaurantGenreController) AddRestaurantGenre(ctx *gin.Context) {
	var json model.RestaurantGenre
	if err := ctx.ShouldBindJSON(&json); err == nil {
		dbError := c.Repository.AddRestaurantGenre(json.RestaurantID, json.GenreID)
		if dbError == nil {
			ctx.JSON(http.StatusOK, gin.H{
				"statusCode": 200,
				"body":       "New restaurant genre is inserted.",
			})
			return
		}
	}

	ctx.JSON(http.StatusBadRequest, gin.H{
		"statusCode": 400,
		"error":      "Restaurant genre insertion failed.",
	})
}
