package main

import (
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"sakaba.link/api/src/controller"
)

func main() {
	homeController := controller.HomeController{}
	photoController := controller.PhotoController{}
	videoController := controller.VideoController{}
	rankingController := controller.RankingController{}
	restaurantController := controller.RestaurantController{}

	router := gin.Default()
	router.Use(cors.New(cors.Config{
		AllowOrigins:  []string{"https://sakaba.link", "https://admin.tokyo-takeout.com"},
		AllowMethods:  []string{"GET"},
		AllowHeaders:  []string{"Origin", "Content-Type"},
		ExposeHeaders: []string{"Content-Length"},
		MaxAge:        12 * time.Hour,
	}))

	router.GET("/", homeController.Index)
	router.GET("/photos/", photoController.GetAllPhotos)
	router.GET("/restaurants/", restaurantController.GetOpenRestaurants)
	router.GET("/restaurant-counts/", restaurantController.GetOpenRestaurantCount)
	router.GET("/rankings/", rankingController.GetAllRankings)
	router.GET("/videos/", videoController.GetAllVideos)
	router.GET("/videos/:id", videoController.GetVideosByRestaurantId)
	router.Run(":8080")
}
