package model

// RestaurantGenre is an entity for Restaurant Genre API.
type RestaurantGenre struct {
	RestaurantID string `json:"restaurant_id"`
	GenreID      string `json:"genre_id"`
}
