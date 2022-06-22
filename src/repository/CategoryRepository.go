package repository

import (
	"sakaba.link/api/src/infrastructure"
	"sakaba.link/api/src/model"
)

// CategoryRepository is responsible for reading from and writing to DB Table `categories`.
type CategoryRepository struct{}

// GetCategories returns the categories for the specified restaurants.
func (c CategoryRepository) GetCategories(restaurantID string) []model.Category {
	db, closer, err := infrastructure.ConnectToDB()
	if err != nil {
		panic(err.Error())
	}
	defer closer()

	categories := []model.Category{}
	db.Table("categories").
		Select("id", "parent_id", "name").
		Where("restaurant_id = ?", infrastructure.UUIDToBin(restaurantID)).
		Order("parent_id ASC").
		Order("id ASC").
		Scan(&categories)

	return categories
}
