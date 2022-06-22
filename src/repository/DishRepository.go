package repository

import (
	"sakaba.link/api/src/infrastructure"
	"sakaba.link/api/src/model"
)

// DishRepository is responsible for reading from and writing to DB Table `dishes`.
type DishRepository struct{}

// GetAllDishes returns all the dishes.
func (c DishRepository) GetAllDishes() []model.Dish {
	db, closer, err := infrastructure.ConnectToDB()
	if err != nil {
		panic(err.Error())
	}
	defer closer()

	allDishes := []model.Dish{}
	db.Table("dishes").
		Select("id", "name").
		Order("name ASC").
		Scan(&allDishes)

	return allDishes
}

// GetDishByID returns the specified dish.
func (c DishRepository) GetDishByID(id string) model.Dish {
	db, closer, err := infrastructure.ConnectToDB()
	if err != nil {
		panic(err.Error())
	}
	defer closer()

	dish := model.Dish{}
	db.Table("dishes").
		Select("id", "name").
		Where("id = ?", id).
		Scan(&dish)

	return dish
}
