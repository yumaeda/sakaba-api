package repository

import (
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
	c.DB.Raw(`SELECT UuidFromBin(id) AS id,
	                 sort_order,
					 name,
					 name_jpn,
					 category,
					 sub_category,
					 region,
					 price,
					 is_min_price
			    FROM menus
			   WHERE restaurant_id = UuidToBin('` + restaurantID + `')
			   ORDER BY category ASC, sub_category ASC, region ASC, sort_order ASC`).Scan(&menus)

	return menus
}
