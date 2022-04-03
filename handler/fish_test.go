package handler

import (
	"go_rest_api/model"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestFishesHandler(t *testing.T) {
	res := httptest.NewRecorder()

	t.Run("fetchFishes", func(t *testing.T) {
		httptest.NewRequest(http.MethodGet, "http://example.com/fishes", nil)
		FishesHandler(&model.FakeFish{})
		expected := `[{"id":1,"name":"アジ","created_at":"2022-04-01T00:00:17Z","updated_at":"2022-04-01T00:00:17Z"},{"id":2,"name":"マダイ","created_at":"2022-04-01T00:00:17Z","updated_at":"2022-04-01T00:00:17Z"}]`
		body, _ := ioutil.ReadAll(res.Body)
		actual := string(body)

		if res.Code != http.StatusOK {
			t.Errorf("It's expected to return 200. But returns %d", res.Code)
		}
		if expected != actual {
			t.Errorf("Response body unmatched. returns %s", actual)
		}
	})

	t.Run("createFish", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodPost, "http://example.com/fishes", nil)
		FishesHandler(res, req)
	})

	t.Run("HTTP method did't matched", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodPatch, "http://example.com/fishes", nil)
		FishesHandler(res, req)

		if res.Code != http.StatusNotFound {
			t.Errorf("It's expected to return 404. But returns %d", res.Code)
		}
	})
}
