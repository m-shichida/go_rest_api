package handler

import (
	"bytes"
	"encoding/json"
	"go_rest_api/seed"
	"net/http"
)

func GetPost(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	if r.Method != "GET" {
		w.WriteHeader(http.StatusNotFound)
	}

	// TODO: DB から取ってくる
}

func GetPosts(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	if r.Method != "GET" {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	var buf bytes.Buffer
	enc := json.NewEncoder(&buf)
	if err := enc.Encode(seed.Posts); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Write(buf.Bytes())
	return
}
