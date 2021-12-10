package main

import (
	"github.com/gin-gonic/gin"
	"sakaba.link/api/src/controller"
)

func CORS(ctx *gin.Context) {
	ctx.Header("Access-Control-Allow-Origin", "https://sakaba.link")
	ctx.Header("Access-Control-Allow-Methods", "GET")
	ctx.Header("Access-Control-Allow-Headers", "*")
	ctx.Header("Content-Type", "application/json")
}

func main() {
	homeController := controller.HomeController{}
	photoController := controller.PhotoController{}
	videoController := controller.VideoController{}
	rankingController := controller.RankingController{}
	restaurantController := controller.RestaurantController{}

	router := gin.Default()
	router.Use(CORS)
	router.GET("/", homeController.Index)
	router.GET("/photos/", photoController.GetAllPhotos)
	router.GET("/restaurants/", restaurantController.GetOpenRestaurants)
	router.GET("/rankings/", rankingController.GetAllRankings)
	router.GET("/videos/", videoController.GetAllVideos)
	router.GET("/videos/:id", videoController.GetVideosByRestaurantId)
	router.Run(":8080")
}
