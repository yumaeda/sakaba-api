package main

import (
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"sakaba.link/api/src/controller"
)

func main() {
	cagegoyController := controller.CategoryController{}
	dishController := controller.DishController{}
	genreController := controller.GenreController{}
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
	router.GET("/categories/:id", cagegoyController.GetCategoriesByRestaurantID)
	router.GET("/dishes/", dishController.GetAllDishes)
	router.GET("/dishes/:id", dishController.GetDishByID)
	router.GET("/genres/", genreController.GetAllGenres)
	router.GET("/genres/:id", genreController.GetGenreByID)
	router.GET("/photos/", photoController.GetAllPhotos)
	router.GET("/restaurants/", restaurantController.GetOpenRestaurants)
	router.GET("/restaurants/dishes/:id", restaurantController.GetOpenRestaurantsByDishID)
	router.GET("/restaurants/genres/:id", restaurantController.GetOpenRestaurantsByGenreID)
	router.GET("/restaurant-counts/", restaurantController.GetOpenRestaurantCount)
	router.GET("/rankings/", rankingController.GetAllRankings)
	router.GET("/videos/", videoController.GetAllVideos)
	router.GET("/videos/:id", videoController.GetVideosByRestaurantID)
	router.Run(":8080")
}
