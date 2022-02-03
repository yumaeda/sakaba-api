package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"sakaba.link/api/src/infrastructure"
	"sakaba.link/api/src/model"
)

type GenreController struct{}

func (c *GenreController) GetAllGenres(ctx *gin.Context) {
	allGenres := []model.Genre{}
	db := infrastructure.ConnectToDB()
	db.Table("genres").
		Select("id", "name").
		Order("name ASC").
		Scan(&allGenres)
	infrastructure.CloseDB(db)

	ctx.JSON(http.StatusOK, gin.H{
		"statusCode": 200,
		"body":       allGenres,
	})
}

func (c *GenreController) GetGenreById(ctx *gin.Context) {
	id := ctx.Param("id")
	genre := model.Genre{}
	db := infrastructure.ConnectToDB()
	db.Table("genres").
		Select("id", "name").
		Where("id = ?", id).
		Scan(&genre)
	infrastructure.CloseDB(db)

	ctx.JSON(http.StatusOK, gin.H{
		"statusCode": 200,
		"body":       genre,
	})
}
