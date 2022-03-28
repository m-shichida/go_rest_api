package handler

import (
	"encoding/json"
	"log"
	"net/http"
)

type ValidationMessages struct {
	Messages []string
}

func renderJSON(w http.ResponseWriter, context interface{}, statusCode int) {
	json, err := json.Marshal(context)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Fatal(err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	w.Write(json)
	return
}
