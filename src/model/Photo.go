package model

// PhotoView is a readonly entity for Photo API.
type PhotoView struct {
	RestaurantID  string `json:"restaurant_id"`
	Name          string `json:"name"`
	Image         string `json:"image"`
	ImageWebp     string `json:"image_webp"`
	Thumbnail     string `json:"thumbnail"`
	ThumbnailWebp string `json:"thumbnail_webp"`
}

// Photo is an entity for Photo API.
type Photo struct {
	RestaurantID string `json:"restaurant_id"`
	Name         string `json:"name"`
	Type         string `json:"type"`
}

// PhotoRequest is an entity for Photo upload API.
type PhotoRequest struct {
	RestaurantID string `json:"restaurant_id"`
	FileContent  string `json:"file_content"`
}
