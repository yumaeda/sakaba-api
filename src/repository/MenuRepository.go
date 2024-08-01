package repository

import (
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

// SetMenu sets meta data for the specified menu.
func (c MenuRepository) SetMenu(menu model.MenuView) error {
	dbError := c.TiDB.First(&menu, "id = ?", infrastructure.UUIDToBinForTiDB(menu.ID)).Error
	if dbError == nil {
		dbError = c.TiDB.Model(&menu).Updates(menu).Error
	}

	return dbError
}

// DeleteMenu deletes meta data from the menus table.
func (c MenuRepository) DeleteMenu(id string) error {
	dbError := c.TiDB.Delete(&model.MenuNew{}, "id = ?", infrastructure.UUIDToBinForTiDB(id)).Error

	return dbError
}
