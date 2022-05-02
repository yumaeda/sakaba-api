package repository

import (
	"gorm.io/gorm"
	"sakaba.link/api/src/infrastructure"
	"sakaba.link/api/src/model"
)

// RestaurantGenreRepository is responsible for reading from and writing to DB Table `restaurant_genres`.
type RestaurantGenreRepository struct{}

// AddRestaurantGenre adds a new genre for the specified restaurant.
func (c *RestaurantGenreRepository) AddRestaurantGenre(restaurantID string, genreID string) *gorm.DB {
	restaurantGenre := model.RestaurantGenre{
		RestaurantID: infrastructure.UUIDToBin(restaurantID),
		GenreID:      genreID,
	}
	db := infrastructure.ConnectToDB()
	result := db.Create(&restaurantGenre)
	infrastructure.CloseDB(db)
	return result
}
