package model

// Category is an entity for Category API.
type Category struct {
	ID       int    `json:"id"`
	ParentID int    `json:"parent_id"`
	Name     string `json:"name"`
}
