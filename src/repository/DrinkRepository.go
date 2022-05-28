package repository

import (
	"sakaba.link/api/src/infrastructure"
	"sakaba.link/api/src/model"
)

// DrinkRepository is responsible for reading from and writing to DB Table `drinks`.
type DrinkRepository struct{}

// GetAllDrinks returns all the drinks.
func (c *DrinkRepository) GetAllDrinks() []model.Drink {
	allDrinks := []model.Drink{}
	db := infrastructure.ConnectToDB()
	db.Table("drinks").
		Select("id", "name").
		Order("name ASC").
		Scan(&allDrinks)
	infrastructure.CloseDB(db)

	return allDrinks
}

// GetDrinkByID returns the specified drink.
func (c *DrinkRepository) GetDrinkByID(id string) model.Drink {
	drink := model.Drink{}
	db := infrastructure.ConnectToDB()
	db.Table("drinks").
		Select("id", "name").
		Where("id = ?", id).
		Scan(&drink)
	infrastructure.CloseDB(db)

	return drink
}
