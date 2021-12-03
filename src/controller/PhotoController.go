package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"sakaba.link/api/src/infrastructure"
	"sakaba.link/api/src/model"
)

type PhotoController struct{}

func (c *PhotoController) GetAllPhotos(ctx *gin.Context) {
	allPhotos := []model.Photo{}
	db := infrastructure.ConnectToDB()

	db.Raw(`SELECT UuidFromBin(r.id) AS restaurant_id,
                       CONCAT(p.name, '.jpg') AS image,
                       CONCAT(p.name, '.webp') AS image_webp,
                       CONCAT(p.name, '_thumbnail.jpg') AS thumbnail,
                       CONCAT(p.name, '_thumbnail.webp') AS thumbnail_webp
                  FROM photos AS p
                  JOIN restaurants AS r
                    ON p.restaurant_id = r.id
                 ORDER BY p.create_time DESC`).Scan(&allPhotos)
	infrastructure.CloseDB(db)

	ctx.JSON(http.StatusOK, gin.H{
		"statusCode": 200,
		"body":       allPhotos,
	})
}
