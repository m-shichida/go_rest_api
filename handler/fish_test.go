package handler

import (
	"go_rest_api/repository"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
)

var mux *http.ServeMux
var writer *httptest.ResponseRecorder

func TestMain(m *testing.M) {
	setUp()
	code := m.Run()
	os.Exit(code)
}

func setUp() {
	mux = http.NewServeMux()
	mux.HandleFunc("/fishes", FishesHandler(&repository.FakeFishRepository{}))
	writer = httptest.NewRecorder()
}

func TestFishesHandler(t *testing.T) {
	t.Run("GET /fishes", func(t *testing.T) {
		request, _ := http.NewRequest(http.MethodGet, "/fishes", nil)
		mux.ServeHTTP(writer, request)

		expected := `[{"Id":1,"Name":"アジ","CreatedAt":"2022-04-01T00:00:00+09:00","UpdatedAt":"2022-04-01T00:00:00+09:00"},{"Id":2,"Name":"マダイ","CreatedAt":"2022-04-01T00:00:00+09:00","UpdatedAt":"2022-04-01T00:00:00+09:00"}]`
		body, _ := ioutil.ReadAll(writer.Body)
		actual := string(body)

		if writer.Code != http.StatusOK {
			t.Errorf("It's expected to return 200. But returns %d", writer.Code)
		}
		if expected != actual {
			t.Errorf("Response body unmatched. returns %s", actual)
		}
	})
}
