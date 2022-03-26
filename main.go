package main

import (
	"go_rest_api/handler"
	"net/http"
	"os"
)

func main() {
	address := os.Getenv("ADDRESS")
	server := http.Server {
		Addr: address,
	}

	http.HandleFunc("/posts", handler.GetPost) // ById から先に判断させる
	http.HandleFunc("/posts", handler.GetPosts)

	server.ListenAndServe()
}
