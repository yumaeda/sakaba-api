package model

// RestaurantCount is an entity for Restaurant Count API.
type RestaurantCount struct {
	Area  string `json:"area"`
	Count int    `json:"count"`
}
