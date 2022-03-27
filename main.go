package main

import (
	"go_rest_api/handler"
	"go_rest_api/model"
	"log"
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

	http.HandleFunc("/fishes", get(handler.GetFishes))
	http.HandleFunc("/fishes/", get(handler.GetFishById))

	server.ListenAndServe()
}

func get(h http.HandlerFunc) http.HandlerFunc {
	return func (w http.ResponseWriter, r *http.Request) {
		if r.Method != "GET" {
			return
		}

		log.Printf("GET: %s", r.URL)
		h(w, r)
	}
}
