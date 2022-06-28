package repository

import (
	"github.com/yumaeda/sakaba-api/src/model"
	"gorm.io/gorm"
)

// DishRepository is responsible for reading from and writing to DB Table `dishes`.
type DishRepository struct {
	DB *gorm.DB
}

// GetAllDishes returns all the dishes.
func (c DishRepository) GetAllDishes() []model.Dish {
	allDishes := []model.Dish{}
	c.DB.Table("dishes").
		Select("id", "name").
		Order("name ASC").
		Scan(&allDishes)

	return allDishes
}

// GetDishByID returns the specified dish.
func (c DishRepository) GetDishByID(id string) model.Dish {
	dish := model.Dish{}
	c.DB.Table("dishes").
		Select("id", "name").
		Where("id = ?", id).
		Scan(&dish)

	return dish
}
