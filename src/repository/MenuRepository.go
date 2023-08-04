package repository

import (
	"github.com/yumaeda/sakaba-api/src/infrastructure"
	"github.com/yumaeda/sakaba-api/src/model"
	"gorm.io/gorm"
)

// MenuRepository is responsible for reading from DB Table `menus`.
type MenuRepository struct {
	DB *gorm.DB
}

// GetMenus returns the menus for the specified restaurants.
func (c MenuRepository) GetMenus(restaurantID string) []model.Menu {
	menus := []model.Menu{}
	c.DB.Table("menus").
		Select("sort_order", "name", "name_jpn", "category", "sub_category", "region", "price", "is_min_price").
		Where("restaurant_id = ?", infrastructure.UUIDToBin(restaurantID)).
		Order("category ASC").
		Order("sub_category ASC").
		Order("region ASC").
		Order("sort_order ASC").
		Scan(&menus)

	return menus
}
