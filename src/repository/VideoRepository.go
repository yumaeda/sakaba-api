package repository

import (
	"sakaba.link/api/src/infrastructure"
	"sakaba.link/api/src/model"
)

type VideoRepository struct{}

func (c *VideoRepository) GetAllVideos() []model.Video {
	allVideos := []model.Video{}
	db := infrastructure.ConnectToDB()
	db.Raw("SELECT id, UuidFromBin(restaurant_id) AS restaurant_id, name, url FROM videos ORDER BY name").Scan(&allVideos)
	infrastructure.CloseDB(db)

	return allVideos
}

func (c *VideoRepository) GetVideosByRestaurantId(id string) []model.SimpleVideo {
	restaurantId := infrastructure.UuidToBin(id)
	videos := []model.SimpleVideo{}
	db := infrastructure.ConnectToDB()
	db.Table("videos").Select("name", "url").Where("restaurant_id = ?", restaurantId).Scan(&videos)
	infrastructure.CloseDB(db)

	return videos
}
