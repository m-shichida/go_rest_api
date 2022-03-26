package model

import "time"

type Post struct {
	Id int
	ImageUrl string
	Address string
	Description string
	Fish Fish
	FishingAt time.Time
	CreatedAt time.Time
	UpdatedAt time.Time
}
