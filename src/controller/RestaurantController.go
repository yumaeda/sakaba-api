package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"sakaba.link/api/src/infrastructure"
	"sakaba.link/api/src/model"
)

type RestaurantController struct{}

func (c *RestaurantController) GetOpenRestaurants(ctx *gin.Context) {
	restaurants := []model.Restaurant{}
	db := infrastructure.ConnectToDB()

	db.Raw(`SELECT UuidFromBin(r.id) AS id,
	               r.url,
	               r.image_name,
	               r.name,
	               r.genre,
	               r.tel,
	               r.business_day_info,
	               r.address,
	               r.latitude,
	               r.longitude,
	               r.area,
	               r.comment,
	               r.takeout_available,
	               COUNT(p.restaurant_id) AS photo_count
                  FROM restaurants AS r
                  LEFT JOIN photos AS p
                    ON r.id = p.restaurant_id
                 WHERE is_closed = 0
                   AND REPLACE(JSON_EXTRACT(r.business_day_info, CONCAT('$.', DAYOFWEEK(CURDATE()), ".Start")), '"', '') <= DATE_FORMAT(CONVERT_TZ(NOW(), '+00:00', '+09:00'), '%H%i')
                   AND REPLACE(JSON_EXTRACT(r.business_day_info, CONCAT('$.', DAYOFWEEK(CURDATE()), ".End")), '"', '') >= DATE_FORMAT(CONVERT_TZ(NOW(), '+00:00', '+09:00'), '%H%i')
                 GROUP BY r.id
                 ORDER BY photo_count DESC`).Scan(&restaurants)
	infrastructure.CloseDB(db)

	ctx.JSON(http.StatusOK, gin.H{
		"statusCode": 200,
		"body":       restaurants,
	})
}

func (c *RestaurantController) GetOpenRestaurantsByGenreId(ctx *gin.Context) {
	genreId := ctx.Param("id")
	restaurants := []model.Restaurant{}
	db := infrastructure.ConnectToDB()

	db.Raw(`SELECT UuidFromBin(r.id) AS id,
	               r.url,
	               r.name,
	               r.genre,
	               r.tel,
	               r.business_day_info,
	               r.address,
	               r.latitude,
	               r.longitude,
	               r.area,
	               r.comment,
	               r.takeout_available,
	               COUNT(p.restaurant_id) AS photo_count
                  FROM restaurants AS r
                  JOIN restaurant_genres AS rg
                    ON r.id = rg.restaurant_id
                  LEFT JOIN photos AS p
                    ON r.id = p.restaurant_id
                 WHERE is_closed = 0
		   AND rg.genre_id = ` + genreId + `
                   AND REPLACE(JSON_EXTRACT(r.business_day_info, CONCAT('$.', DAYOFWEEK(CURDATE()), ".Start")), '"', '') <= DATE_FORMAT(CONVERT_TZ(NOW(), '+00:00', '+09:00'), '%H%i')
                   AND REPLACE(JSON_EXTRACT(r.business_day_info, CONCAT('$.', DAYOFWEEK(CURDATE()), ".End")), '"', '') >= DATE_FORMAT(CONVERT_TZ(NOW(), '+00:00', '+09:00'), '%H%i')
                 GROUP BY r.id
                 ORDER BY photo_count DESC`).Scan(&restaurants)
	infrastructure.CloseDB(db)

	ctx.JSON(http.StatusOK, gin.H{
		"statusCode": 200,
		"body":       restaurants,
	})
}

func (c *RestaurantController) GetOpenRestaurantsByDishId(ctx *gin.Context) {
	dishId := ctx.Param("id")
	restaurants := []model.Restaurant{}
	db := infrastructure.ConnectToDB()

	db.Raw(`SELECT UuidFromBin(r.id) AS id,
	               r.url,
	               r.name,
	               r.genre,
	               r.tel,
	               r.business_day_info,
	               r.address,
	               r.latitude,
	               r.longitude,
	               r.area,
	               r.comment,
	               r.takeout_available,
	               0 AS photo_count
	          FROM dishes AS d
	          JOIN rankings AS rk
                    ON rk.dish_id = d.id
                  JOIN photos AS p
                    ON p.id = rk.photo_id
                  JOIN restaurants AS r
                    ON r.id = p.restaurant_id
                 WHERE r.is_closed = 0
                   AND d.id = ` + dishId + `
                   AND REPLACE(JSON_EXTRACT(r.business_day_info, CONCAT('$.', DAYOFWEEK(CURDATE()), ".Start")), '"', '') <= DATE_FORMAT(CONVERT_TZ(NOW(), '+00:00', '+09:00'), '%H%i')
                   AND REPLACE(JSON_EXTRACT(r.business_day_info, CONCAT('$.', DAYOFWEEK(CURDATE()), ".End")), '"', '') >= DATE_FORMAT(CONVERT_TZ(NOW(), '+00:00', '+09:00'), '%H%i')
                 ORDER BY rk.rank ASC`).Scan(&restaurants)
	infrastructure.CloseDB(db)

	ctx.JSON(http.StatusOK, gin.H{
		"statusCode": 200,
		"body":       restaurants,
	})
}

func (c *RestaurantController) GetOpenRestaurantCount(ctx *gin.Context) {
	restaurantCount := []model.RestaurantCount{}
	db := infrastructure.ConnectToDB()

	db.Raw(`SELECT area,
                       COUNT(area) AS count
                  FROM restaurants
                 WHERE is_closed = 0
                   AND REPLACE(JSON_EXTRACT(business_day_info, CONCAT('$.', DAYOFWEEK(CURDATE()), ".Start")), '"', '') <= DATE_FORMAT(CONVERT_TZ(NOW(), '+00:00', '+09:00'), '%H%i')
                   AND REPLACE(JSON_EXTRACT(business_day_info, CONCAT('$.', DAYOFWEEK(CURDATE()), ".End")), '"', '') >= DATE_FORMAT(CONVERT_TZ(NOW(), '+00:00', '+09:00'), '%H%i')
                 GROUP BY area
                 ORDER BY COUNT(area) DESC`).Scan(&restaurantCount)
	infrastructure.CloseDB(db)

	ctx.JSON(http.StatusOK, gin.H{
		"statusCode": 200,
		"body":       restaurantCount,
	})
}
