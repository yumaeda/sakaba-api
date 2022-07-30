package repository

import (
	"github.com/yumaeda/sakaba-api/src/model"
	"gorm.io/gorm"
)

// AreaRepository is responsible for reading from and writing to DB Table `areas`.
type AreaRepository struct {
	DB *gorm.DB
}

// GetAllAreas returns all the areas.
func (c AreaRepository) GetAllAreas() []model.Area {
	allAreas := []model.Area{}
	c.DB.Table("areas").
		Select("id", "name", "value").
		Order("name ASC").
		Scan(&allAreas)

	return allAreas
}
