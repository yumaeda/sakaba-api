package model

type Category struct {
	Id       int    `json:"id"`
	ParentId int    `json:"parent_id"`
	Name     string `json:"name"`
}
