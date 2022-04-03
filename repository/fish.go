package repository

import (
	"database/sql"
	"go_rest_api/model"
	"log"
)

type FishRepository struct {
	Db *sql.DB
}

func (fr *FishRepository) Index() (fishes []model.Fish, err error) {
  rows, err := fr.Db.Query("SELECT id, name, created_at, updated_at FROM fishes")
	if err != nil {
		log.Fatalf("Failed to execute query.")
		return
	}
	for rows.Next() {
		fish := model.Fish{}
		if err = rows.Scan(&fish.Id, &fish.Name, &fish.CreatedAt, &fish.UpdatedAt); err != nil {
			rows.Close()
			log.Fatalf("Failed to execute query.")
			return
		}
		fishes = append(fishes, fish)
	}
	rows.Close()
	return
}

func (fr *FishRepository) FetchFishById(id int) (fish model.Fish, err error) {
  err =
		fr.Db.QueryRow("SELECT id, name, created_at, updated_at FROM fishes WHERE id = ?", id).Scan(&fish.Id, &fish.Name, &fish.CreatedAt, &fish.UpdatedAt)
	return
}

func (fr *FishRepository) FetchFishByName(name string) (fish model.Fish, err error) {
  err =
		fr.Db.QueryRow("SELECT id, name, created_at, updated_at FROM fishes WHERE name = ?", name).Scan(&fish.Id, &fish.Name, &fish.CreatedAt, &fish.UpdatedAt)
	return
}

func (fr *FishRepository) Create(fish *model.Fish) (err error) {
	statement, err := fr.Db.Prepare("INSERT INTO fishes (name) VALUES (?)")
	if err != nil {
		return
	}
	statement.Exec(fish.Name)

	err = fr.Db.QueryRow("SELECT id, created_at, updated_at FROM fishes WHERE name = ?", fish.Name).Scan(&fish.Id, &fish.CreatedAt, &fish.UpdatedAt)
	return
}

func (fr *FishRepository) Update(fish *model.Fish) (err error) {
	statement, err := fr.Db.Prepare("UPDATE fishes SET name = ? WHERE id = ?")
	if err != nil {
		return
	}
	statement.Exec(fish.Name, fish.Id)
	err = fr.Db.QueryRow("SELECT updated_at FROM fishes WHERE id = ?", fish.Id).Scan(&fish.UpdatedAt)
	return
}

func (fr *FishRepository) Delete(id int) (err error) {
	statement, err := fr.Db.Prepare("DELETE FROM fishes WHERE id = ?")
	if err != nil {
		return
	}
	statement.Exec(id)
	return
}
