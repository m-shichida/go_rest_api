package model

import (
	"log"
	"time"
)

type Fish struct {
	Id int
	Name string
	CreatedAt time.Time
	UpdatedAt time.Time
}

func (fish *Fish) Validate() (messages []string) {
	var name string
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

	err := Db.QueryRow("SELECT name FROM fishes WHERE name = ?", fish.Name).Scan(&name)
	if err == nil {
		message := "すでにその名前は使用されています"
		messages = append(messages, message)
	}

	return
}

func (fish *Fish) Index() (fishes []Fish, err error) {
  rows, err := Db.Query("SELECT id, name, created_at, updated_at FROM fishes")
	if err != nil {
		log.Fatalf("Failed to execute query.")
	}
	for rows.Next() {
		fish := Fish{}
		if err = rows.Scan(&fish.Id, &fish.Name, &fish.CreatedAt, &fish.UpdatedAt); err != nil {
			log.Fatalf("Failed to execute query.")
		}
		fishes = append(fishes, fish)
	}
	rows.Close()
	return
}

func (fish *Fish) GetById(id int) (err error) {
  err =
		Db.QueryRow("SELECT id, name, created_at, updated_at FROM fishes WHERE id = ?", id).Scan(&fish.Id, &fish.Name, &fish.CreatedAt, &fish.UpdatedAt)
	return
}

func (fish *Fish) Create() (err error) {
	statement, err := Db.Prepare("INSERT INTO fishes (name) VALUES (?)")
	if err != nil {
		return
	}
	statement.Exec(fish.Name)

	err = Db.QueryRow("SELECT id, created_at, updated_at FROM fishes WHERE name = ?", fish.Name).Scan(&fish.Id, &fish.CreatedAt, &fish.UpdatedAt)
	return
}

func (fish *Fish) Update() (err error) {
	statement, err := Db.Prepare("UPDATE fishes SET name = ? WHERE id = ?")
	if err != nil {
		return
	}
	statement.Exec(fish.Name, fish.Id)
	err = Db.QueryRow("SELECT updated_at FROM fishes WHERE id = ?", fish.Id).Scan(&fish.UpdatedAt)
	return
}

func (fish *Fish) Delete() (err error) {
	statement, err := Db.Prepare("DELETE FROM fishes WHERE id = ?")
	if err != nil {
		return
	}
	statement.Exec(fish.Id)
	return
}
