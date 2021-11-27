package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"sakaba.link/api/src/controller"
)

func CORS(ctx *gin.Context) {
	ctx.Header("Access-Control-Allow-Origin", "*")
	ctx.Header("Access-Control-Allow-Methods", "GET, OPTIONS")
	ctx.Header("Access-Control-Allow-Headers", "*")
	ctx.Header("Content-Type", "application/json")

	// Second, we handle the OPTIONS problem
	if ctx.Request.Method != "OPTIONS" {
		ctx.Next()
	} else {

		// Everytime we receive an OPTIONS request,
		// we just return an HTTP 200 Status Code
		// Like this, Angular can now do the real
		// request using any other method than OPTIONS
		ctx.AbortWithStatus(http.StatusOK)
	}
}

/*
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
*/

func main() {
	homeController := controller.HomeController{}
	videoController := controller.VideoController{}

	router := gin.Default()
	router.Use(CORS)
	router.GET("/", homeController.Index)
	router.GET("/videos/", videoController.GetAllVideos)
	router.GET("/videos/:id", videoController.GetVideosByRestaurantId)
	router.Run(":8080")
}
