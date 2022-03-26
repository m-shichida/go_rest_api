package handler

import (
	"bytes"
	"encoding/json"
	"go_rest_api/model"
	"log"
	"net/http"
)

func GetFishes(w http.ResponseWriter, r *http.Request) {
	var fish model.Fish
	w.Header().Set("Content-Type", "application/json")
	if r.Method != "GET" {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	fishes, err := fish.Index()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Fatal(err)
		return
	}

	var buf bytes.Buffer
	enc := json.NewEncoder(&buf)
	if err := enc.Encode(fishes); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Write(buf.Bytes())
	return
}
