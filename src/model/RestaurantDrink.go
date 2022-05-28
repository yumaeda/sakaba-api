package model

// RestaurantDrink is an entity for Restaurant Drink API.
type RestaurantDrink struct {
	RestaurantID string `json:"restaurant_id"`
	DrinkID      string `json:"drink_id"`
}
