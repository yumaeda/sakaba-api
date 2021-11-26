package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"sakaba.link/api/infrastructure"
	"sakaba.link/api/models"
)

func home(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "hello world"})
}

func UuidFromBin(uuid []byte) string {
	return fmt.Sprintf("%x-%x-%x-%x-%x", uuid[0:4], uuid[4:6], uuid[6:8], uuid[8:10], uuid[10:])
}

func getVideosForRestaurant(c *gin.Context) {
	n := c.Param("id")
	newDB := infrastructure.NewDatabase()

	rows, err := newDB.DB.Raw("SELECT UuidToBin('" + n + "')").Rows()
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"statusCode": 400,
			"body":       err.Error(),
		})
		return
	}
	defer rows.Close()

	var restaurantId string
	for rows.Next() {
		if err := rows.Scan(&restaurantId); err != nil {
			c.JSON(http.StatusOK, gin.H{
				"statusCode": 500,
				"body":       err.Error(),
			})
			return
		}
	}

	allVideos := []models.Video{}
	newDB.DB.Where(models.Video{RestaurantId: restaurantId}).Find(&allVideos)

	c.JSON(http.StatusOK, gin.H{
		"statusCode": 200,
		"body":       allVideos,
	})
}

func main() {
	router := gin.Default()
	router.GET("/", home)
	router.GET("/videos/:id", getVideosForRestaurant)
	router.Run(":8080")
}
