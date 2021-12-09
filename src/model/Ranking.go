package model

type Ranking struct {
	Dish          string `json:"dish"`
	Rank          int    `json:"rank"`
	Restaurant    string `json:"restaurant"`
	RestaurantId  string `json:"restaurant_id"`
	Photo         string `json:"photo"`
	RestaurantUrl string `json:"restaurant_url"`
}
