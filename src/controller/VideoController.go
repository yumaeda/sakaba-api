package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"sakaba.link/api/src/infrastructure"
	"sakaba.link/api/src/model"
)

type VideoController struct{}

func setCORS(ctx *gin.Context) {
	ctx.Header("Access-Control-Allow-Origin", "https://sakaba.link")
	ctx.Header("Access-Control-Allow-Methods", "GET, OPTIONS")
}

func (c *VideoController) GetAllVideos(ctx *gin.Context) {
	allVideos := []model.Video{}
	db := infrastructure.ConnectToDB()
	db.Raw("SELECT id, UuidFromBin(restaurant_id) AS restaurant_id, name, url FROM videos ORDER BY name").Scan(&allVideos)

	setCORS(ctx)
	ctx.JSON(http.StatusOK, gin.H{
		"statusCode": 200,
		"body":       allVideos,
	})
}

func (c *VideoController) GetVideosByRestaurantId(ctx *gin.Context) {
	id := ctx.Param("id")
	restaurantId := infrastructure.UuidToBin(id)
	videos := []model.SimpleVideo{}
	db := infrastructure.ConnectToDB()
	db.Table("videos").Select("name", "url").Where("restaurant_id = ?", restaurantId).Scan(&videos)

	setCORS(ctx)
	ctx.JSON(http.StatusOK, gin.H{
		"statusCode": 200,
		"body":       videos,
	})
}
