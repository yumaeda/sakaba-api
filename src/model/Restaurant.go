package model

// Restaurant is an entity for Restaurant API.
type Restaurant struct {
	ID               string `json:"id"`
	URL              string `json:"url"`
	ImageName        string `json:"image_name"`
	Name             string `json:"name"`
	Genre            string `json:"genre"`
	Tel              string `json:"tel"`
	BusinessDayInfo  string `json:"business_day_info"`
	Address          string `json:"address"`
	Latitude         string `json:"latitude"`
	Longitude        string `json:"longitude"`
	Area             string `json:"area"`
	Comment          string `json:"comment"`
	TakeoutAvailable int    `json:"takeout_available"`
	PhotoCount       int    `json:"photo_count"`
}
