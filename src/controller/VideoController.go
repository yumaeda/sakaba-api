package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"sakaba.link/api/src/infrastructure"
	"sakaba.link/api/src/model"
)

type VideoController struct{}

func (c *VideoController) GetAllVideos(ctx *gin.Context) {
	newDB := infrastructure.NewDatabase()
	allVideos := []model.Video{}
	newDB.DB.Find(&allVideos)

	ctx.JSON(http.StatusOK, gin.H{
		"statusCode": 200,
		"body":       allVideos,
	})
}

func (c *VideoController) GetVideosByRestaurantId(ctx *gin.Context) {
	n := ctx.Param("id")
	newDB := infrastructure.NewDatabase()

	rows, err := newDB.DB.Raw("SELECT UuidToBin('" + n + "')").Rows()
	if err != nil {
		ctx.JSON(http.StatusOK, gin.H{
			"statusCode": 400,
			"body":       err.Error(),
		})
		return
	}
	defer rows.Close()

	var restaurantId string
	for rows.Next() {
		if err := rows.Scan(&restaurantId); err != nil {
			ctx.JSON(http.StatusOK, gin.H{
				"statusCode": 500,
				"body":       err.Error(),
			})
			return
		}
	}

	allVideos := []model.Video{}
	newDB.DB.Where(model.Video{RestaurantId: restaurantId}).Find(&allVideos)

	ctx.JSON(http.StatusOK, gin.H{
		"statusCode": 200,
		"body":       allVideos,
	})
}
