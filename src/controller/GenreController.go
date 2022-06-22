package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"sakaba.link/api/src/repository"
)

// GenreController is a controller for Genre API.
type GenreController struct {
	Repository repository.GenreRepository
}

// GetAllGenres returns all the genres.
func (c GenreController) GetAllGenres(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"statusCode": 200,
		"body":       c.Repository.GetAllGenres(),
	})
}

// GetGenreByID returns the specified genre.
func (c GenreController) GetGenreByID(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"statusCode": 200,
		"body":       c.Repository.GetGenreByID(ctx.Param("id")),
	})
}
