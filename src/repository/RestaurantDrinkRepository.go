package repository

import (
	"sakaba.link/api/src/infrastructure"
	"sakaba.link/api/src/model"
)

// RestaurantDrinkRepository is responsible for reading from and writing to DB Table `restaurant_drinks`.
type RestaurantDrinkRepository struct{}

// AddRestaurantDrink adds a new drink for the specified restaurant.
func (c *RestaurantDrinkRepository) AddRestaurantDrink(restaurantID string, drinkID string) error {
	restaurantDrink := model.RestaurantDrink{
		RestaurantID: infrastructure.UUIDToBin(restaurantID),
		DrinkID:      drinkID,
	}
	db := infrastructure.ConnectToDB()
	dbError := db.Create(&restaurantDrink).Error
	infrastructure.CloseDB(db)

	return dbError
}
