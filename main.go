package main

import (
	"go_rest_api/handler"
	"go_rest_api/model"
	"net/http"
	"os"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	address := os.Getenv("ADDRESS")
	model.DatabaseInit()
	server := http.Server {
		Addr: address,
	}

	http.HandleFunc("/fishes", handler.GetFishes)
	http.HandleFunc("/fishes/", handler.GetFishById)

	server.ListenAndServe()
}
