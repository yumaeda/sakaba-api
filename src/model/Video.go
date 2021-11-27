package model

type Video struct {
	Id           int    `json:"id"`
	RestaurantId string `json:"restaurant_id"`
	Name         string `json:"name"`
	Url          string `json:"url"`
}

type SimpleVideo struct {
	Name string `json:"name"`
	Url  string `json:"url"`
}
