package repository

import (
	"sakaba.link/api/src/infrastructure"
	"sakaba.link/api/src/model"
)

type CategoryRepository struct{}

func (c *CategoryRepository) GetCategories(restaurantId string) []model.Category {
	categories := []model.Category{}
	db := infrastructure.ConnectToDB()
	db.Table("categories").
		Select("id", "parent_id", "name").
		Where("restaurant_id = ?", infrastructure.UuidToBin(restaurantId)).
		Order("parent_id ASC").
		Order("id ASC").
		Scan(&categories)
	infrastructure.CloseDB(db)

	return categories
}
