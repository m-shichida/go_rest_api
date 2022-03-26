package seed

import (
	"go_rest_api/model"
	"time"
)

var Fishes = []model.Fish{
	{
		Id: 1,
		Name: "ヒラメ",
		CreatedAt: time.Now().Add(-2 * time.Hour),
		UpdatedAt: time.Now().Add(-2 * time.Hour),
	},
	{
		Id: 2,
		Name: "マダイ",
		CreatedAt: time.Now().Add(-2 * time.Hour),
		UpdatedAt: time.Now().Add(-2 * time.Hour),
	},
	{
		Id: 3,
		Name: "ブリ",
		CreatedAt: time.Now().Add(-2 * time.Hour),
		UpdatedAt: time.Now().Add(-2 * time.Hour),
	},
}
