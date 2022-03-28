package handler

import (
	"bytes"
	"encoding/json"
	"log"
	"net/http"
)

type ValidationMessages struct {
	Messages []string
}

func renderJSON(w http.ResponseWriter, context interface{}, statusCode int) {
	var buf bytes.Buffer
	enc := json.NewEncoder(&buf)
	if err := enc.Encode(context); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Fatal(err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	w.Write(buf.Bytes())
	return
}
