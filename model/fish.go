package model

import (
	"time"
)

type Fish struct {
	Id int
	Name string
	CreatedAt time.Time
	UpdatedAt time.Time
}

func (fish *Fish) Validate() (messages []string) {
	if fish.Name == "" {
		message := "名前を入力してください"
		messages = append(messages, message)

		return
	}

	if fish.Name != "" && len(fish.Name) > 50 {
		message := "名前は50文字以下で入力してください"
		messages = append(messages, message)

		return
	}

	// err := Db.QueryRow("SELECT name FROM fishes WHERE name = ?", fish.Name).Scan(&name)
	// if err == nil {
	// 	message := "すでにその名前は使用されています"
	// 	messages = append(messages, message)
	// }

	return
}
