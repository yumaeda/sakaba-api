package main

import (
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"sakaba.link/api/src/controller"
)

func getCorsConfig() cors.Config {
	return cors.Config{
		AllowMethods: []string{
			"GET",
			"OPTIONS",
		},
		AllowHeaders: []string{
			"Access-Control-Allow-Headers",
			"Content-Type",
			"Content-Length",
			"Accept-Encoding",
		},
		AllowOrigins: []string{
			"https://sakaba.link",
		},
		MaxAge: 12 * time.Hour,
	}
}

func main() {
	homeController := controller.HomeController{}
	videoController := controller.VideoController{}

	router := gin.Default()
	router.Use(cors.New(getCorsConfig()))
	router.GET("/", homeController.Index)
	router.GET("/videos/", videoController.GetAllVideos)
	router.GET("/videos/:id", videoController.GetVideosByRestaurantId)
	router.Run(":8080")
}
