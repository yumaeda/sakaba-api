package repository

import (
	"sakaba.link/api/src/infrastructure"
	"sakaba.link/api/src/model"
)

// DrinkRepository is responsible for reading from and writing to DB Table `drinks`.
type DrinkRepository struct{}

// GetAllDrinks returns all the drinks.
func (c *DrinkRepository) GetAllDrinks() []model.Drink {
	db, closer, err := infrastructure.ConnectToDB()
	if err != nil {
		panic(err.Error())
	}
	defer closer()

	allDrinks := []model.Drink{}
	db.Table("drinks").
		Select("id", "name").
		Order("name ASC").
		Scan(&allDrinks)

	return allDrinks
}

// GetDrinkByID returns the specified drink.
func (c *DrinkRepository) GetDrinkByID(id string) model.Drink {
	db, closer, err := infrastructure.ConnectToDB()
	if err != nil {
		panic(err.Error())
	}
	defer closer()

	drink := model.Drink{}
	db.Table("drinks").
		Select("id", "name").
		Where("id = ?", id).
		Scan(&drink)

	return drink
}
