package repository

import (
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
	c.DB.Raw(`SELECT id,
	                 parent_id,
					 name
                FROM categories
               WHERE BIN_TO_UUID(restaurant_id, 1) = '` + restaurantID + `'
               ORDER BY parent_id ASC, sort_order ASC`).Scan(&categories)

	return categories
}
