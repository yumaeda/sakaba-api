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

// GetPhotosByRestaurantID returns all the photos for the specified restaurant.
func (c PhotoRepository) GetPhotosByRestaurantID(restaurantID string) []model.PhotoView {
	photos := []model.PhotoView{}
	c.DB.Raw(`SELECT CONCAT(name, '.jpg') AS image,
                     CONCAT(name, '.webp') AS image_webp,
                     CONCAT(name, '_thumbnail.jpg') AS thumbnail,
                     CONCAT(name, '_thumbnail.webp') AS thumbnail_webp
                FROM photos
   		       WHERE UuidFromBin(restaurant_id) = '` + restaurantID + `'
               ORDER BY create_time DESC`).Scan(&photos)

	return photos
}

// GetLatestPhotos returns all the photos.
func (c PhotoRepository) GetLatestPhotos() []model.PhotoView {
	photos := []model.PhotoView{}
	c.DB.Raw(`SELECT UuidFromBin(restaurant_id) AS restaurant_id,
	                 CONCAT(name, '.jpg') AS image,
                     CONCAT(name, '.webp') AS image_webp,
                     CONCAT(name, '_thumbnail.jpg') AS thumbnail,
                     CONCAT(name, '_thumbnail.webp') AS thumbnail_webp
                FROM photos
               ORDER BY create_time DESC
               LIMIT 20`).Scan(&photos)

	return photos
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
