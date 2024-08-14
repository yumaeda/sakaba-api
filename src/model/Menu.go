package model

// MenuView is a readonly entity for Menu API.
type MenuView struct {
	ID          string `json:"id"`
	SortOrder   int    `json:"sort_order"`
	Category    int    `json:"category"`
	SubCategory int    `json:"sub_category"`
	Region      int    `json:"region"`
	Name        string `json:"name"`
	NameJpn     string `json:"name_jpn"`
	Price       int    `json:"price"`
	IsMinPrice  int    `json:"is_min_price"`
	IsHidden    int    `json:"is_hidden"`
}

// Menu is an entity for Menu API.
type Menu struct {
	ID           string `json:"id"`
	RestaurantID string `json:"restaurant_id"`
	Name         string `json:"name"`
	NameJpn      string `json:"name_jpn"`
	Price        int    `json:"price"`
}

type MenuUpdate struct {
	ID     string `json:"id"`
	Column string `json:"column"`
	Value  string `json:"value"`
}
