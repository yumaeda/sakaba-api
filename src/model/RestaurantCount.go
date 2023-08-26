package model

// RestaurantCount is an entity for Restaurant Count API.
type RestaurantCount struct {
	Area  string `json:"area"`
	Name  string `json:"name"`
	Count int    `json:"count"`
}
