package repository

import (
	"sakaba.link/api/src/infrastructure"
	"sakaba.link/api/src/model"
)

// DishRepository is responsible for reading from and writing to DB Table `dishes`.
type DishRepository struct{}

// GetAllDishes returns all the dishes.
func (c *DishRepository) GetAllDishes() []model.Dish {
	allDishes := []model.Dish{}
	db := infrastructure.ConnectToDB()
	db.Table("dishes").
		Select("id", "name").
		Order("name ASC").
		Scan(&allDishes)
	infrastructure.CloseDB(db)

	return allDishes
}

// GetDishByID returns the specified dish.
func (c *DishRepository) GetDishByID(id string) model.Dish {
	dish := model.Dish{}
	db := infrastructure.ConnectToDB()
	db.Table("dishes").
		Select("id", "name").
		Where("id = ?", id).
		Scan(&dish)
	infrastructure.CloseDB(db)

	return dish
}
