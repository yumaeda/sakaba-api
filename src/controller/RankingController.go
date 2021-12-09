package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"sakaba.link/api/src/infrastructure"
	"sakaba.link/api/src/model"
)

type RankingController struct{}

func (c *RankingController) GetAllRankings(ctx *gin.Context) {
	allRankings := []model.Ranking{}
	db := infrastructure.ConnectToDB()

	db.Raw(`SELECT dishes.name AS 'dish',
	               rankings.rank AS 'rank',
	               restaurants.name AS 'restaurant',
	               UuidFromBin(restaurants.id) AS 'restaurant_id',
	               photos.name AS 'photo',
	               restaurants.url AS 'restaurant_url'
                  FROM rankings
                  JOIN dishes
                    ON rankings.dish_id = dishes.id
                  JOIN photos
                    ON rankings.photo_id = photos.id
                  JOIN restaurants
                    ON photos.restaurant_id = restaurants.id
                 ORDER BY dishes.name ASC, rankings.rank ASC`).Scan(&allRankings)
	infrastructure.CloseDB(db)

	ctx.JSON(http.StatusOK, gin.H{
		"statusCode": 200,
		"body":       allRankings,
	})
}
