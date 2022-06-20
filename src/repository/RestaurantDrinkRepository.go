package repository

import (
	"sakaba.link/api/src/infrastructure"
	"sakaba.link/api/src/model"
)

// RestaurantDrinkRepository is responsible for reading from and writing to DB Table `restaurant_drinks`.
type RestaurantDrinkRepository struct{}

// AddRestaurantDrink adds a new drink for the specified restaurant.
func (c *RestaurantDrinkRepository) AddRestaurantDrink(restaurantID string, drinkID string) error {
	db, closer, err := infrastructure.ConnectToDB()
	if err != nil {
		panic(err.Error())
	}
	defer closer()

	restaurantDrink := model.RestaurantDrink{
		RestaurantID: infrastructure.UUIDToBin(restaurantID),
		DrinkID:      drinkID,
	}
	dbError := db.Create(&restaurantDrink).Error

	return dbError
}
