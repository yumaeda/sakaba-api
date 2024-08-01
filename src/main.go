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
		AllowOrigins:  []string{"https://sakabas.com", "https://admin.tokyo-dinner.com", "http://www.baroceans.com"},
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

	tidb, tidbCloser, tidbErr := infrastructure.ConnectToTiDB()
	if tidbErr != nil {
		panic(tidbErr.Error())
	}
	defer tidbCloser()

	adminController := controller.AdminController{}
	areaController := controller.AreaController{Repository: repository.AreaRepository{DB: tidb}}
	categoryController := controller.CategoryController{Repository: repository.CategoryRepository{DB: tidb}}
	dishController := controller.DishController{Repository: repository.DishRepository{DB: tidb}}
	drinkController := controller.DrinkController{Repository: repository.DrinkRepository{DB: tidb}}
	genreController := controller.GenreController{Repository: repository.GenreRepository{DB: tidb}}
	healtchCheckController := controller.HealthCheckController{}
	menuController := controller.MenuController{Repository: repository.MenuRepository{DB: db, TiDB: tidb}}
	photoController := controller.PhotoController{Repository: repository.PhotoRepository{DB: tidb}}
	videoController := controller.VideoController{Repository: repository.VideoRepository{DB: tidb}}
	rankingController := controller.RankingController{Repository: repository.RankingRepository{DB: tidb}}
	restaurantController := controller.RestaurantController{Repository: repository.RestaurantRepository{TiDB: tidb}}
	restaurantDrinkController := controller.RestaurantDrinkController{Repository: repository.RestaurantDrinkRepository{DB: tidb}}
	restaurantGenreController := controller.RestaurantGenreController{Repository: repository.RestaurantGenreRepository{DB: tidb}}

	router.GET("/", healtchCheckController.GetStatus)
	router.GET("/areas/", areaController.GetAllAreas)
	router.GET("/categories/:id", categoryController.GetCategoriesByRestaurantID)
	router.GET("/dishes/", dishController.GetAllDishes)
	router.GET("/dishes/:id", dishController.GetDishByID)
	router.GET("/drinks/", drinkController.GetAllDrinks)
	router.GET("/drinks/:id", drinkController.GetDrinkByID)
	router.GET("/genres/", genreController.GetAllGenres)
	router.GET("/genres/:id", genreController.GetGenreByID)
	router.GET("/health/", healtchCheckController.GetStatus)
	router.GET("/photos/:id", photoController.GetPhotosByRestaurantID)
	router.GET("/latest-photos/", photoController.GetLatestPhotos)
	router.GET("/menus/:id", menuController.GetMenusByRestaurantID)
	router.GET("/menus2/:id", menuController.GetMenusByRestaurantIDFromTiDB)
	router.GET("/restaurants/", restaurantController.GetRestaurants)
	router.GET("/restaurants/areas/:id", restaurantController.GetRestaurantsByArea)
	router.GET("/restaurants/dishes/:id", restaurantController.GetRestaurantsByDishID)
	router.GET("/restaurants/drinks/:id", restaurantController.GetRestaurantsByDrinkID)
	router.GET("/restaurants/genres/:id", restaurantController.GetRestaurantsByGenreID)
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
		auth.POST("/restaurant-drink/", restaurantDrinkController.AddRestaurantDrink)
		auth.POST("/restaurant-genre/", restaurantGenreController.AddRestaurantGenre)
		auth.POST("/restaurant/", restaurantController.AddRestaurant)
		auth.PUT("/menu/", menuController.SetMenu)
		auth.POST("/menu/", menuController.AddMenu)
		auth.DELETE("/menu/:id", menuController.DeleteMenu)
	}

	router.Run(":8080")
}
