package repository

import (
	"gorm.io/gorm"
	"sakaba.link/api/src/model"
)

// DrinkRepository is responsible for reading from and writing to DB Table `drinks`.
type DrinkRepository struct {
	DB *gorm.DB
}

// GetAllDrinks returns all the drinks.
func (c DrinkRepository) GetAllDrinks() []model.Drink {
	allDrinks := []model.Drink{}
	c.DB.Table("drinks").
		Select("id", "name").
		Order("name ASC").
		Scan(&allDrinks)

	return allDrinks
}

// GetDrinkByID returns the specified drink.
func (c DrinkRepository) GetDrinkByID(id string) model.Drink {
	drink := model.Drink{}
	c.DB.Table("drinks").
		Select("id", "name").
		Where("id = ?", id).
		Scan(&drink)

	return drink
}
