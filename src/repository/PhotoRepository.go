package repository

import (
	"github.com/yumaeda/sakaba-api/src/infrastructure"
	"github.com/yumaeda/sakaba-api/src/model"
	"gorm.io/gorm"
)

// PhotoRepository is responsible for reading from and writing to DB Table `photos`.
type PhotoRepository struct {
	DB *gorm.DB
}

// GetAllPhotos returns all the photos.
func (c PhotoRepository) GetAllPhotos() []model.PhotoView {
	allPhotos := []model.PhotoView{}
	c.DB.Raw(`SELECT UuidFromBin(r.id) AS restaurant_id,
                       CONCAT(p.name, '.jpg') AS image,
                       CONCAT(p.name, '.webp') AS image_webp,
                       CONCAT(p.name, '_thumbnail.jpg') AS thumbnail,
                       CONCAT(p.name, '_thumbnail.webp') AS thumbnail_webp
                  FROM photos AS p
                  JOIN restaurants AS r
                    ON p.restaurant_id = r.id
                 ORDER BY p.create_time DESC`).Scan(&allPhotos)

	return allPhotos
}

// AddPhoto adds meta data for the new photo.
func (c PhotoRepository) AddPhoto(restaurantID string, fileName string) error {
	photo := model.Photo{
		RestaurantID: infrastructure.UUIDToBin(restaurantID),
		Type:         "dish",
		Name:         fileName,
	}
	dbError := c.DB.Create(&photo).Error

	return dbError
}
