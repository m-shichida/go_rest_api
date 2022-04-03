package main

import (
	"go_rest_api/handler"
	"go_rest_api/repository"
	"go_rest_api/utils"
	"net/http"
	"os"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	address := os.Getenv("ADDRESS")
	db := utils.DatabaseInit()
	server := http.Server {
		Addr: address,
	}

	http.HandleFunc("/fishes", handler.FishesHandler(&repository.FishRepository{Db: db}))
	http.HandleFunc("/fishes/", handler.FishHandler(&repository.FishRepository{Db: db})) // /fishes/:id

	server.ListenAndServe()
}
