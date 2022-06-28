package repository

import (
	"github.com/yumaeda/sakaba-api/src/infrastructure"
	"github.com/yumaeda/sakaba-api/src/model"
	"gorm.io/gorm"
)

// VideoRepository is responsible for reading from and writing to DB Table `videos`.
type VideoRepository struct {
	DB *gorm.DB
}

// GetAllVideos returns all the videos.
func (c VideoRepository) GetAllVideos() []model.Video {
	allVideos := []model.Video{}
	c.DB.Raw("SELECT id, UuidFromBin(restaurant_id) AS restaurant_id, name, url FROM videos ORDER BY name").Scan(&allVideos)

	return allVideos
}

// GetVideosByRestaurantID returns videos for the specified restaurant.
func (c VideoRepository) GetVideosByRestaurantID(id string) []model.SimpleVideo {
	restaurantID := infrastructure.UUIDToBin(id)
	videos := []model.SimpleVideo{}
	c.DB.Table("videos").Select("name", "url").Where("restaurant_id = ?", restaurantID).Scan(&videos)

	return videos
}
