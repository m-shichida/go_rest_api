package seed

import (
	"go_rest_api/model"
	"time"
)

var Posts = []model.Post{
	{
		Id: 1,
		ImageUrl: "https://user-images.githubusercontent.com/42334697/160239680-e5562dd9-3a02-43e0-b559-ec9d000e466b.jpeg",
		Fish: model.Fish{},
		FishingAt: time.Now().Add(-1 * time.Hour),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	},
	{
		Id: 2,
		ImageUrl: "https://user-images.githubusercontent.com/42334697/160240005-ebabb7a2-3fe2-419f-80b0-5d1c6d05be76.jpeg",
		Fish: model.Fish{},
		FishingAt: time.Now().Add(-1 * time.Hour),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	},
	{
		Id: 3,
		ImageUrl: "https://user-images.githubusercontent.com/42334697/160240040-9e50abb0-3860-46f2-8990-ddc230360121.jpeg",
		Fish: model.Fish{},
		FishingAt: time.Now().Add(-1 * time.Hour),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	},
}
