package repository

import (
	"sakaba.link/api/src/infrastructure"
	"sakaba.link/api/src/model"
)

// CategoryRepository is responsible for reading from and writing to DB Table `categories`.
type CategoryRepository struct{}

// GetCategories returns the categories for the specified restaurants.
func (c *CategoryRepository) GetCategories(restaurantID string) []model.Category {
	categories := []model.Category{}
	db := infrastructure.ConnectToDB()
	db.Table("categories").
		Select("id", "parent_id", "name").
		Where("restaurant_id = ?", infrastructure.UUIDToBin(restaurantID)).
		Order("parent_id ASC").
		Order("id ASC").
		Scan(&categories)
	infrastructure.CloseDB(db)

	return categories
}
