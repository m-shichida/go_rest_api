package handler

import (
	"go_rest_api/model"
	"log"
	"net/http"
	"regexp"
	"strconv"
)

func GetFishes(w http.ResponseWriter, r *http.Request) {
	var fish model.Fish

	fishes, err := fish.Index()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Fatal(err)
		return
	}
	renderJSON(w, fishes)
}

func GetFishById(w http.ResponseWriter, r *http.Request) {
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
	renderJSON(w, fish)
}
