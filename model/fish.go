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
