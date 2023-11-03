package main

import (
	"log"
	"time"

	jwt "github.com/appleboy/gin-jwt/v2"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/yumaeda/sakaba-api/src/controller"
	"github.com/yumaeda/sakaba-api/src/infrastructure"
	"github.com/yumaeda/sakaba-api/src/middleware"
	"github.com/yumaeda/sakaba-api/src/repository"
)

var realm = "Sakaba Link Zone"
var identityKey = "id"
var secretKey = "testKey"

func main() {
	router := gin.Default()
	router.Use(cors.New(cors.Config{
		AllowOrigins:  []string{"https://sakaba.link", "https://admin.sakaba.link", "https://admin.tokyo-dinner.com", "http://www.baroceans.com"},
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

	db, closer, err := infrastructure.ConnectToDB()
	if err != nil {
		panic(err.Error())
	}
	defer closer()

	adminController := controller.AdminController{}
	areaController := controller.AreaController{Repository: repository.AreaRepository{DB: db}}
	cagegoyController := controller.CategoryController{Repository: repository.CategoryRepository{DB: db}}
	dishController := controller.DishController{Repository: repository.DishRepository{DB: db}}
	drinkController := controller.DrinkController{Repository: repository.DrinkRepository{DB: db}}
	genreController := controller.GenreController{Repository: repository.GenreRepository{DB: db}}
	healtchCheckController := controller.HealthCheckController{}
	menuController := controller.MenuController{Repository: repository.MenuRepository{DB: db}}
	photoController := controller.PhotoController{Repository: repository.PhotoRepository{DB: db}}
	videoController := controller.VideoController{Repository: repository.VideoRepository{DB: db}}
	rankingController := controller.RankingController{Repository: repository.RankingRepository{DB: db}}
	restaurantController := controller.RestaurantController{Repository: repository.RestaurantRepository{DB: db}}
	restaurantDrinkController := controller.RestaurantDrinkController{Repository: repository.RestaurantDrinkRepository{DB: db}}
	restaurantGenreController := controller.RestaurantGenreController{Repository: repository.RestaurantGenreRepository{DB: db}}

	router.GET("/", healtchCheckController.GetStatus)
	router.GET("/areas/", areaController.GetAllAreas)
	router.GET("/categories/:id", cagegoyController.GetCategoriesByRestaurantID)
	router.GET("/dishes/", dishController.GetAllDishes)
	router.GET("/dishes/:id", dishController.GetDishByID)
	router.GET("/drinks/", drinkController.GetAllDrinks)
	router.GET("/drinks/:id", drinkController.GetDrinkByID)
	router.GET("/genres/", genreController.GetAllGenres)
	router.GET("/genres/:id", genreController.GetGenreByID)
	router.GET("/health/", healtchCheckController.GetStatus)
	router.GET("/photos/", photoController.GetAllPhotos)
	router.GET("/latest-photos/", photoController.GetLatestPhotos)
	router.GET("/menus/:id", menuController.GetMenusByRestaurantID)
	router.GET("/restaurants/", restaurantController.GetRestaurants)
	router.GET("/restaurants/areas/:id/:latitude/:longitude", restaurantController.GetRestaurantsByArea)
	router.GET("/restaurants/dishes/:id/:latitude/:longitude", restaurantController.GetRestaurantsByDishID)
	router.GET("/restaurants/dishes/:id", restaurantController.GetOpenRestaurantsByDishID)
	router.GET("/restaurants/drinks/:id/:latitude/:longitude", restaurantController.GetRestaurantsByDrinkID)
	router.GET("/restaurants/drinks/:id", restaurantController.GetOpenRestaurantsByDrinkID)
	router.GET("/restaurants/genres/:id/:latitude/:longitude", restaurantController.GetRestaurantsByGenreID)
	router.GET("/restaurants/genres/:id", restaurantController.GetOpenRestaurantsByGenreID)
	router.GET("/restaurant-counts/:latitude/:longitude", restaurantController.GetOpenRestaurantCount)
	router.GET("/rankings/", rankingController.GetAllRankings)
	router.GET("/videos/", videoController.GetAllVideos)
	router.GET("/videos/:id", videoController.GetVideosByRestaurantID)

	auth := router.Group("/auth")
	auth.GET("/refresh_token", middleware.RefreshHandler)
	auth.Use(middleware.MiddlewareFunc())
	{
		auth.GET("/home", adminController.Index)
		auth.POST("/photo/", photoController.AddPhoto)
		auth.POST("/restaurant-drink/", restaurantDrinkController.AddRestaurantDrink)
		auth.POST("/restaurant-genre/", restaurantGenreController.AddRestaurantGenre)
		auth.POST("/restaurant/", restaurantController.AddRestaurant)
	}

	router.Run(":8080")
}
