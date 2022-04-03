package model

import (
	"time"
)

type FakeFish struct {
	Id int
	Name string
	CreatedAt time.Time
	UpdatedAt time.Time
}

func (fakeFish *FakeFish) Validate() (messages []string) {
	return
}

	func (fakeFish *FakeFish) Index() (fishes []Fish, err error) {
	fishes = []Fish{
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

func (fakeFish *FakeFish) GetById(int) (fish Fish, err error) {
	fish = Fish{
		Id: 1,
		Name: "アジ",
		CreatedAt: time.Date(2022, 4, 1, 0, 0, 0, 0, time.Local),
		UpdatedAt: time.Date(2022, 4, 1, 0, 0, 0, 0, time.Local),
	}
	return
}

func (fakeFish *FakeFish) Create() (err error) {
	return
}

func (fakeFish *FakeFish) Update() (err error) {
	return
}

func (fakeFish *FakeFish) Delete() (err error) {
	return
}
