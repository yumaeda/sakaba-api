package models

import (
	"time"
)

type Video struct {
	Id           int
	RestaurantId string
	Name         string
	Url          string
	CreateTime   time.Time
	UpdateTime   time.Time
}
