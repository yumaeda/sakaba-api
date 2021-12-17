package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"sakaba.link/api/src/infrastructure"
	"sakaba.link/api/src/model"
)

type CategoryController struct{}

func (c *CategoryController) GetCategoriesByRestaurantId(ctx *gin.Context) {
	id := ctx.Param("id")
	restaurantId := infrastructure.UuidToBin(id)
	categories := []model.Category{}
	db := infrastructure.ConnectToDB()
	db.Table("categories").
		Select("id", "parent_id", "name").
		Where("restaurant_id = ?", restaurantId).
		Order("parent_id ASC").
		Order("id ASC").
		Scan(&categories)
	infrastructure.CloseDB(db)

	ctx.JSON(http.StatusOK, gin.H{
		"statusCode": 200,
		"body":       categories,
	})
}
