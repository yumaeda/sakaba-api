package repository

import (
	"sakaba.link/api/src/infrastructure"
	"sakaba.link/api/src/model"
)

// RestaurantGenreRepository is responsible for reading from and writing to DB Table `restaurant_genres`.
type RestaurantGenreRepository struct{}

// AddRestaurantGenre adds a new genre for the specified restaurant.
func (c *RestaurantGenreRepository) AddRestaurantGenre(restaurantID string, genreID string) error {
	restaurantGenre := model.RestaurantGenre{
		RestaurantID: infrastructure.UUIDToBin(restaurantID),
		GenreID:      genreID,
	}
	db := infrastructure.ConnectToDB()
	dbError := db.Create(&restaurantGenre).Error
	infrastructure.CloseDB(db)

	return dbError
}
