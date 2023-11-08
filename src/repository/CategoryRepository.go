package repository

import (
	"github.com/yumaeda/sakaba-api/src/infrastructure"
	"github.com/yumaeda/sakaba-api/src/model"
	"gorm.io/gorm"
)

// CategoryRepository is responsible for reading from and writing to DB Table `categories`.
type CategoryRepository struct {
	DB *gorm.DB
}

// GetCategories returns the categories for the specified restaurants.
func (c CategoryRepository) GetCategories(restaurantID string) []model.Category {
	categories := []model.Category{}
	c.DB.Table("categories").
		Select("id", "parent_id", "name").
		Where("restaurant_id = ?", infrastructure.UUIDToBin(restaurantID)).
		Order("parent_id ASC").
		Order("sort_order ASC").
		Scan(&categories)

	return categories
}
