package handler

import (
	"bytes"
	"encoding/json"
	"log"
	"net/http"
)

func renderJSON(w http.ResponseWriter, context interface{}) {
	var buf bytes.Buffer
	enc := json.NewEncoder(&buf)
	if err := enc.Encode(context); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Fatal(err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(buf.Bytes())
	return
}
