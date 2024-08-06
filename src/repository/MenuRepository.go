package repository

import (
	"errors"

	"github.com/google/uuid"
	"github.com/yumaeda/sakaba-api/src/infrastructure"
	"github.com/yumaeda/sakaba-api/src/model"
	"gorm.io/gorm"
)

// MenuRepository is responsible for reading from DB Table `menus`.
type MenuRepository struct {
	DB   *gorm.DB
	TiDB *gorm.DB
}

// contains checks if the specified string is included in the specified array.
func contains(array []string, str string) bool {
	for _, item := range array {
		if item == str {
			return true
		}
	}

	return false
}

// GetMenus returns the menus for the specified restaurants.
func (c MenuRepository) GetMenus(restaurantID string) []model.MenuView {
	menus := []model.MenuView{}
	c.DB.Raw(`SELECT UuidFromBin(id) AS id,
		             sort_order,
					 name,
					 name_jpn,
					 category,
					 sub_category,
					 region,
					 price,
					 is_min_price,
					 is_hidden
				FROM menus
			   WHERE UuidFromBin(restaurant_id) = '` + restaurantID + `'
			   ORDER BY category ASC, sub_category ASC, region ASC, sort_order ASC`).Scan(&menus)

	return menus
}

// GetMenus returns the menus for the specified restaurants from TiDB.
func (c MenuRepository) GetMenusFromTiDB(restaurantID string) []model.MenuView {
	menus := []model.MenuView{}
	c.TiDB.Raw(`SELECT BIN_TO_UUID(id, 1) AS id,
		               sort_order,
					   name,
					   name_jpn,
					   category,
					   sub_category,
					   region,
					   price,
					   is_min_price,
					   is_hidden
                  FROM menus
                 WHERE BIN_TO_UUID(restaurant_id, 1) = '` + restaurantID + `'
			     ORDER BY category ASC, sub_category ASC, region ASC, sort_order ASC`).Scan(&menus)

	return menus
}

// AddMenu adds meta data for the new menu.
func (c MenuRepository) AddMenu(restaurantID string) (string, error) {
	id := uuid.New().String()
	menu := model.MenuNew{
		ID:           infrastructure.UUIDToBinForTiDB(id),
		RestaurantID: infrastructure.UUIDToBinForTiDB(restaurantID),
		Name:         "",
		NameJpn:      "",
		Price:        0,
	}
	dbError := c.TiDB.Create(&menu).Error

	return id, dbError
}

// SetMenu updates the specified column with specified value for the specified menu.
func (c MenuRepository) SetMenu(ID string, column string, value string) error {
	var COLUMNS = []string{"name", "name_jpn", "price", "category", "sub_category", "region", "sort_order", "is_min_price", "is_hidden"}
	if !contains(COLUMNS, column) {
		return errors.New("invalid column name")
	}

	dbError := c.TiDB.Exec(`UPDATE menus
                               SET ` + column + ` = '` + value + `'
                             WHERE BIN_TO_UUID(id, 1) = '` + ID + `'`).Error

	return dbError
}

// DeleteMenu deletes meta data from the menus table.
func (c MenuRepository) DeleteMenu(id string) error {
	dbError := c.TiDB.Delete(&model.MenuNew{}, "id = ?", infrastructure.UUIDToBinForTiDB(id)).Error

	return dbError
}
