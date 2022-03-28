package handler

import (
	"encoding/json"
	"go_rest_api/model"
	"log"
	"net/http"
	"regexp"
	"strconv"
)

func FishesHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		fetchFishes(w)
	case http.MethodPost:
		createFish(w, r)
	default:
		w.WriteHeader(http.StatusNotFound)
		return
	}
}

func FishHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		fetchFish(w, r)
	default:
		w.WriteHeader(http.StatusNotFound)
		return
	}
}

func fetchFishes(w http.ResponseWriter) {
	var fish model.Fish

	fishes, err := fish.Index()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Fatal(err)
		return
	}
	renderJSON(w, fishes, http.StatusOK)
}

func fetchFish(w http.ResponseWriter, r *http.Request) {
	var fish model.Fish

	rgx := regexp.MustCompile(`^/fishes/([0-9]+)$`)
	paths := rgx.FindStringSubmatch(r.URL.Path)
	// /fishes/:id だと [/fishes/:id, :id] として取得できる
	if len(paths) != 2 {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	id, err := strconv.Atoi(paths[1])
	if err != nil {
		log.Fatal(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	err = fish.GetById(id)
	if fish.Id == 0 {
		w.WriteHeader(http.StatusNotFound)
		return
	}
	renderJSON(w, fish, http.StatusOK)
}

func createFish(w http.ResponseWriter, r *http.Request) {
	var fish model.Fish
	var validationMessages ValidationMessages
	body := make([]byte, r.ContentLength)
	r.Body.Read(body)
	defer r.Body.Close()
	json.Unmarshal(body, &fish)

	messages := fish.Validate()
	if messages != nil {
		validationMessages.Messages = messages

		renderJSON(w, validationMessages, http.StatusBadRequest)
		return
	}

	err := fish.Create()
	if err != nil {
		log.Fatal(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	renderJSON(w, fish, http.StatusOK)
	return
}
