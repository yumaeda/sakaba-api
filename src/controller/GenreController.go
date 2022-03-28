package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"sakaba.link/api/src/repository"
)

// GenreController is a controller for Genre API.
type GenreController struct{}

// GetAllGenres returns all the genres.
func (c *GenreController) GetAllGenres(ctx *gin.Context) {
	genreRepository := repository.GenreRepository{}
	allGenres := genreRepository.GetAllGenres()

	ctx.JSON(http.StatusOK, gin.H{
		"statusCode": 200,
		"body":       allGenres,
	})
}

// GetGenreByID returns the specified genre.
func (c *GenreController) GetGenreByID(ctx *gin.Context) {
	id := ctx.Param("id")
	genreRepository := repository.GenreRepository{}
	genre := genreRepository.GetGenreByID(id)

	ctx.JSON(http.StatusOK, gin.H{
		"statusCode": 200,
		"body":       genre,
	})
}
