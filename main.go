package main

import (
	"fmt"
	"net/http"
	"os"
)

func main() {
	address := os.Getenv("ADDRESS")
	server := http.Server {
		Addr: address,
	}
	http.HandleFunc("/hoge", hoge)
	server.ListenAndServe()
}

func hoge(w http.ResponseWriter, r *http.Request) {
	fmt.Println("hoge")
}
