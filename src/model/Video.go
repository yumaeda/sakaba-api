package model

// Video is an entity for Video API.
type Video struct {
	ID           int    `json:"id"`
	RestaurantID string `json:"restaurant_id"`
	Name         string `json:"name"`
	URL          string `json:"url"`
}

// SimpleVideo is an entity for Video API.
type SimpleVideo struct {
	Name string `json:"name"`
	URL  string `json:"url"`
}
