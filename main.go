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

// そして GET 関数の内部でJSON処理を書く

func get(h http.HandlerFunc) http.HandlerFunc {
	return func (w http.ResponseWriter, r *http.Request) {
		if r.Method != "GET" {
			return
		}

		log.Printf("called %s", r.URL.Path)
		h(w, r)
	}
}

func renderJSON() {}
