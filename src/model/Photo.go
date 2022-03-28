package model

// Photo is an entity for Photo API.
type Photo struct {
	RestaurantID  string `json:"restaurant_id"`
	Name          string `json:"name"`
	Image         string `json:"image"`
	ImageWebp     string `json:"image_webp"`
	Thumbnail     string `json:"thumbnail"`
	ThumbnailWebp string `json:"thumbnail_webp"`
}
