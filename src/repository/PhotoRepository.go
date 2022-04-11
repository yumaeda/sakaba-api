package repository

import (
	"sakaba.link/api/src/infrastructure"
	"sakaba.link/api/src/model"
)

// PhotoRepository is responsible for reading from and writing to DB Table `photos`.
type PhotoRepository struct{}

// GetAllPhotos returns all the photos.
func (c *PhotoRepository) GetAllPhotos() []model.Photo {
	allPhotos := []model.Photo{}
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
