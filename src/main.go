package main

import (
	"log"
	"time"

	jwt "github.com/appleboy/gin-jwt/v2"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"sakaba.link/api/src/controller"
	"sakaba.link/api/src/middleware"
)

var realm = "Sakaba Link Zone"
var identityKey = "id"
var secretKey = "testKey"

func main() {
	router := gin.Default()
	router.Use(cors.New(cors.Config{
		AllowOrigins:  []string{"https://sakaba.link", "https://admin.tokyo-takeout.com"},
		AllowMethods:  []string{"GET", "POST"},
		AllowHeaders:  []string{"Origin", "Content-Type", "Authorization"},
		ExposeHeaders: []string{"Content-Length"},
		MaxAge:        12 * time.Hour,
	}))

	authMiddleware := middleware.AuthMiddleware{}
	middleware, err := authMiddleware.Init(realm, identityKey, secretKey)
	if err != nil {
		log.Fatal("JWT Error:" + err.Error())
	}

	errInit := middleware.MiddlewareInit()
	if errInit != nil {
		log.Fatal(errInit.Error())
	}

	router.POST("/login", middleware.LoginHandler)
	router.NoRoute(middleware.MiddlewareFunc(), func(c *gin.Context) {
		claims := jwt.ExtractClaims(c)
		log.Printf("NoRoute claims: %#v\n", claims)
		c.JSON(404, gin.H{"code": "PAGE_NOT_FOUND", "message": "Page not found"})
	})

	adminController := controller.AdminController{}
	cagegoyController := controller.CategoryController{}
	dishController := controller.DishController{}
	genreController := controller.GenreController{}
	healtchCheckController := controller.HealthCheckController{}
	photoController := controller.PhotoController{}
	videoController := controller.VideoController{}
	rankingController := controller.RankingController{}
	restaurantController := controller.RestaurantController{}
	restaurantGenreController := controller.RestaurantGenreController{}

	router.GET("/", healtchCheckController.GetStatus)
	router.GET("/categories/:id", cagegoyController.GetCategoriesByRestaurantID)
	router.GET("/dishes/", dishController.GetAllDishes)
	router.GET("/dishes/:id", dishController.GetDishByID)
	router.GET("/genres/", genreController.GetAllGenres)
	router.GET("/genres/:id", genreController.GetGenreByID)
	router.GET("/health/", healtchCheckController.GetStatus)
	router.GET("/photos/", photoController.GetAllPhotos)
	router.GET("/restaurants/", restaurantController.GetOpenRestaurants)
	router.GET("/restaurants/dishes/:id", restaurantController.GetOpenRestaurantsByDishID)
	router.GET("/restaurants/genres/:id", restaurantController.GetOpenRestaurantsByGenreID)
	router.GET("/restaurant-counts/", restaurantController.GetOpenRestaurantCount)
	router.GET("/rankings/", rankingController.GetAllRankings)
	router.GET("/videos/", videoController.GetAllVideos)
	router.GET("/videos/:id", videoController.GetVideosByRestaurantID)

	auth := router.Group("/auth")
	auth.GET("/refresh_token", middleware.RefreshHandler)
	auth.Use(middleware.MiddlewareFunc())
	{
		auth.GET("/home", adminController.Index)
		auth.POST("/photo/", photoController.AddPhoto)
		auth.POST("/restaurant-genre/", restaurantGenreController.AddRestaurantGenre)
		auth.POST("/restaurant/", restaurantController.AddRestaurant)
	}

	router.Run(":8080")
}
