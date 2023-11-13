package model

// Menu is an entity for Menu API.
type Menu struct {
	ID          string `json:"id"`
	SortOrder   int    `json:"sort_order"`
	Category    int    `json:"category"`
	SubCategory int    `json:"sub_category"`
	Region      int    `json:"region"`
	Name        string `json:"name"`
	NameJpn     string `json:"name_jpn"`
	Price       int    `json:"price"`
	isMinPrice  int    `json:"is_min_price"`
	isHidden    int    `json:"is_hidden"`
}
