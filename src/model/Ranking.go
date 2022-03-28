package model

// Ranking is an entity for Ranking API.
type Ranking struct {
	Dish          string `json:"dish"`
	Rank          int    `json:"rank"`
	Restaurant    string `json:"restaurant"`
	RestaurantID  string `json:"restaurant_id"`
	Photo         string `json:"photo"`
	RestaurantURL string `json:"restaurant_url"`
}
