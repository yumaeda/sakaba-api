package repository

import (
	"sakaba.link/api/src/infrastructure"
	"sakaba.link/api/src/model"
)

// VideoRepository is responsible for reading from and writing to DB Table `videos`.
type VideoRepository struct{}

// GetAllVideos returns all the videos.
func (c *VideoRepository) GetAllVideos() []model.Video {
	allVideos := []model.Video{}
	db := infrastructure.ConnectToDB()
	db.Raw("SELECT id, UuidFromBin(restaurant_id) AS restaurant_id, name, url FROM videos ORDER BY name").Scan(&allVideos)
	infrastructure.CloseDB(db)

	return allVideos
}

// GetVideosByRestaurantID returns videos for the specified restaurant.
func (c *VideoRepository) GetVideosByRestaurantID(id string) []model.SimpleVideo {
	restaurantID := infrastructure.UUIDToBin(id)
	videos := []model.SimpleVideo{}
	db := infrastructure.ConnectToDB()
	db.Table("videos").Select("name", "url").Where("restaurant_id = ?", restaurantID).Scan(&videos)
	infrastructure.CloseDB(db)

	return videos
}
