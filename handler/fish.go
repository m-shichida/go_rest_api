package handler

import (
	"encoding/json"
	"go_rest_api/model"
	"log"
	"net/http"
	"regexp"
	"strconv"
)

type FishText interface {
	Index() ([]model.Fish, error)
	FetchFishById(int) (model.Fish, error)
	FetchFishByName(string) (model.Fish, error)
	Create(*model.Fish) (error)
	Update(*model.Fish) (error)
	Delete(int) (error)
}

func FishesHandler(t FishText) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var err error

		switch r.Method {
		case http.MethodGet:
			err = fetchFishes(w, t)
		case http.MethodPost:
			err = createFish(w, r, t)
		default:
			w.WriteHeader(http.StatusNotFound)
			return
		}
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			log.Fatal(err)
			return
		}
	}
}

func FishHandler(t FishText) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var err error

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

		switch r.Method {
		case http.MethodGet:
			err = fetchFish(w, id, t)
		case http.MethodPatch:
			err = updateFish(w, r, id, t)
		case http.MethodDelete:
			err = deleteFish(w, id, t)
		default:
			w.WriteHeader(http.StatusNotFound)
			return
		}
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			log.Fatal(err)
			return
		}
	}
}

func fetchFishes(w http.ResponseWriter, t FishText) (err error) {
	fishes, err := t.Index()
	if err != nil {
		return
	}

	renderJSON(w, fishes, http.StatusOK)
	return
}

func fetchFish(w http.ResponseWriter, id int, t FishText) (err error) {
	fish, err := t.FetchFishById(id)
	if err != nil {
		return
	}

	if fish.Id == 0 {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	renderJSON(w, fish, http.StatusOK)
	return
}

func createFish(w http.ResponseWriter, r *http.Request, t FishText) (err error) {
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

	err = t.Create(&fish)
	if err != nil {
		return
	}
	renderJSON(w, fish, http.StatusOK)
	return
}

func updateFish(w http.ResponseWriter, r *http.Request, id int, t FishText) (err error) {
	var validationMessages ValidationMessages
	fish, err := t.FetchFishById(id)
	if err != nil {
		return
	}

	if fish.Id == 0 {
		w.WriteHeader(http.StatusNotFound)
		return
	}

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

	err = t.Update(&fish)
	if err != nil {
		return
	}

	renderJSON(w, fish, http.StatusOK)
	return
}

func deleteFish(w http.ResponseWriter, id int, t FishText) (err error) {
	fish, err := t.FetchFishById(id)
	if err != nil {
		return
	}

	if fish.Id == 0 {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	err = t.Delete(fish.Id)
	if err != nil {
		return
	}

	w.WriteHeader(http.StatusOK)
	return
}
