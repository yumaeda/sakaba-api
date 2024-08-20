package repository

import (
	"errors"

	"github.com/yumaeda/sakaba-api/src/infrastructure"
	"github.com/yumaeda/sakaba-api/src/model"
	"gorm.io/gorm"
)

// MenuRepository is responsible for reading from DB Table `menus`.
type MenuRepository struct {
	DB *gorm.DB
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
	c.DB.Raw(`SELECT BIN_TO_UUID(id, 1) AS id,
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
func (c MenuRepository) AddMenu(ID string, restaurantID string, name string, nameJpn string, price int) error {
	menu := model.Menu{
		ID:           infrastructure.UUIDToBin(ID),
		RestaurantID: infrastructure.UUIDToBin(restaurantID),
		Name:         name,
		NameJpn:      nameJpn,
		Price:        price,
	}
	dbError := c.DB.Create(&menu).Error

	return dbError
}

// SetMenu updates the specified column with specified value for the specified menu.
func (c MenuRepository) SetMenu(ID string, column string, value string) error {
	var COLUMNS = []string{"name", "name_jpn", "price", "category", "sub_category", "region", "sort_order", "is_min_price", "is_hidden"}
	if !contains(COLUMNS, column) {
		return errors.New("invalid column name")
	}

	dbError := c.DB.Exec(`UPDATE menus
                               SET ` + column + ` = '` + value + `'
                             WHERE BIN_TO_UUID(id, 1) = '` + ID + `'`).Error

	return dbError
}

// DeleteMenu deletes meta data from the menus table.
func (c MenuRepository) DeleteMenu(ID string) error {
	dbError := c.DB.Exec(`DELETE FROM menus
                           WHERE BIN_TO_UUID(id, 1) = '` + ID + `'`).Error

	return dbError
}
