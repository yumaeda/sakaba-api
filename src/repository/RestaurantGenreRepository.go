package repository

import (
	"gorm.io/gorm"
	"sakaba.link/api/src/infrastructure"
	"sakaba.link/api/src/model"
)

// RestaurantGenreRepository is responsible for reading from and writing to DB Table `restaurant_genres`.
type RestaurantGenreRepository struct {
	DB *gorm.DB
}

// AddRestaurantGenre adds a new genre for the specified restaurant.
func (c RestaurantGenreRepository) AddRestaurantGenre(restaurantID string, genreID string) error {
	restaurantGenre := model.RestaurantGenre{
		RestaurantID: infrastructure.UUIDToBin(restaurantID),
		GenreID:      genreID,
	}
	dbError := c.DB.Create(&restaurantGenre).Error

	return dbError
}
