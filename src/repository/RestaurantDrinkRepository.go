package repository

import (
	"github.com/yumaeda/sakaba-api/src/infrastructure"
	"github.com/yumaeda/sakaba-api/src/model"
	"gorm.io/gorm"
)

// RestaurantDrinkRepository is responsible for reading from and writing to DB Table `restaurant_drinks`.
type RestaurantDrinkRepository struct {
	DB *gorm.DB
}

// AddRestaurantDrink adds a new drink for the specified restaurant.
func (c RestaurantDrinkRepository) AddRestaurantDrink(restaurantID string, drinkID string) error {
	restaurantDrink := model.RestaurantDrink{
		RestaurantID: infrastructure.UUIDToBinForTiDB(restaurantID),
		DrinkID:      drinkID,
	}
	dbError := c.DB.Create(&restaurantDrink).Error

	return dbError
}
