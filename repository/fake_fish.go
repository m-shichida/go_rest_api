package repository

import (
	"go_rest_api/model"
	"time"
)

type FakeFishRepository struct {}

func (ffr *FakeFishRepository) Index() (fishes []model.Fish, err error) {
	fishes = []model.Fish{
		{
			Id: 1,
			Name: "アジ",
			CreatedAt: time.Date(2022, 4, 1, 0, 0, 0, 0, time.Local),
			UpdatedAt: time.Date(2022, 4, 1, 0, 0, 0, 0, time.Local),
		},
		{
			Id: 2,
			Name: "マダイ",
			CreatedAt: time.Date(2022, 4, 1, 0, 0, 0, 0, time.Local),
			UpdatedAt: time.Date(2022, 4, 1, 0, 0, 0, 0, time.Local),
		},
	}
	return
}

func (ffr *FakeFishRepository) FetchFishById(id int) (fish model.Fish, err error) {
	fish = model.Fish{
		Id: 1,
		Name: "アジ",
		CreatedAt: time.Date(2022, 4, 1, 0, 0, 0, 0, time.Local),
		UpdatedAt: time.Date(2022, 4, 1, 0, 0, 0, 0, time.Local),
	}
	return
}

func (ffr *FakeFishRepository) FetchFishByName(name string) (fish model.Fish, err error) {
	return
}

func (ffr *FakeFishRepository) Create(fish *model.Fish) (err error) {
	return
}

func (ffr *FakeFishRepository) Update(fish *model.Fish) (err error) {
	return
}

func (ffr *FakeFishRepository) Delete(id int) (err error) {
	return
}
