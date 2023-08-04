package model

// Mneu is an entity for Menu API.
type Menu struct {
	ID          int    `json:"id"`
	SortOrder   int    `json:"sort_order"`
	Name        string `json:"name"`
	NameJpn     string `json:"name_jpn"`
	Category    string `json:"category"`
	SubCategory string `json:"sub_category"`
	Region      string `json:"region"`
	Price       int    `json:"price"`
	isMinPrice  bool   `json:"is_min_price"`
}
