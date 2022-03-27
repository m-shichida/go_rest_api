package handler

import (
	"bytes"
	"encoding/json"
	"go_rest_api/model"
	"log"
	"net/http"
	"regexp"
	"strconv"
)

// 全てのメソッドはもうJSONで返すのでヘッダのセットとかはもう関数でラップして各関数で設定しないようにしたみ
// GET 以外では return する、みたいな関数もあっても良い

func GetFishes(w http.ResponseWriter, r *http.Request) {
	var fish model.Fish
	if r.Method != "GET" {
		return
	}

	w.Header().Set("Content-Type", "application/json")

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
		log.Fatal(err)
		return
	}

	w.Write(buf.Bytes())
	return
}

func GetFishById(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		return
	}

	w.Header().Set("Content-Type", "application/json")
	var fish model.Fish
	var buf bytes.Buffer

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
	if err != nil {
		log.Fatal(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	enc := json.NewEncoder(&buf)
	if err := enc.Encode(fish); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Fatal(err)
		return
	}

	w.Write(buf.Bytes())
	return
}
