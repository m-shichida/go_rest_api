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
	if fish.Name == "" {
		message := "名前を入力してください"
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
