package repository

import (
	"gorm.io/gorm"
	"sakaba.link/api/src/infrastructure"
	"sakaba.link/api/src/model"
)

// PhotoRepository is responsible for reading from and writing to DB Table `photos`.
type PhotoRepository struct{}

// GetAllPhotos returns all the photos.
func (c *PhotoRepository) GetAllPhotos() []model.PhotoView {
	allPhotos := []model.PhotoView{}
	db := infrastructure.ConnectToDB()
	db.Raw(`SELECT UuidFromBin(r.id) AS restaurant_id,
                       CONCAT(p.name, '.jpg') AS image,
                       CONCAT(p.name, '.webp') AS image_webp,
                       CONCAT(p.name, '_thumbnail.jpg') AS thumbnail,
                       CONCAT(p.name, '_thumbnail.webp') AS thumbnail_webp
                  FROM photos AS p
                  JOIN restaurants AS r
                    ON p.restaurant_id = r.id
                 ORDER BY p.create_time DESC`).Scan(&allPhotos)
	infrastructure.CloseDB(db)

	return allPhotos
}

// AddPhoto adds meta data for the new photo.
func (c *PhotoRepository) AddPhoto(restaurantID string, fileName string) *gorm.DB {
	photo := model.Photo{
		RestaurantID: infrastructure.UUIDToBin(restaurantID),
		Name:         fileName,
	}
	db := infrastructure.ConnectToDB()
	result := db.Create(&photo)
	infrastructure.CloseDB(db)

	return result
}
