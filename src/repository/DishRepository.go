package repository

import (
	"sakaba.link/api/src/infrastructure"
	"sakaba.link/api/src/model"
)

type DishRepository struct{}

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

func (c *DishRepository) GetDishById(id string) model.Dish {
	dish := model.Dish{}
	db := infrastructure.ConnectToDB()
	db.Table("dishes").
		Select("id", "name").
		Where("id = ?", id).
		Scan(&dish)
	infrastructure.CloseDB(db)

	return dish
}
