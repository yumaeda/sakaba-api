package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"sakaba.link/api/src/repository"
)

type GenreController struct{}

func (c *GenreController) GetAllGenres(ctx *gin.Context) {
	genreRepository := repository.GenreRepository{}
	allGenres := genreRepository.GetAllGenres()

	ctx.JSON(http.StatusOK, gin.H{
		"statusCode": 200,
		"body":       allGenres,
	})
}

func (c *GenreController) GetGenreById(ctx *gin.Context) {
	id := ctx.Param("id")
	genreRepository := repository.GenreRepository{}
	genre := genreRepository.GetGenreById(id)

	ctx.JSON(http.StatusOK, gin.H{
		"statusCode": 200,
		"body":       genre,
	})
}
