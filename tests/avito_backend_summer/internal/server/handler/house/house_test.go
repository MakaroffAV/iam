package house

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/MakaroffAV/backend-bootcamp-assignment-2024/internal/repo/db"
	"github.com/MakaroffAV/backend-bootcamp-assignment-2024/internal/repo/db/dbmodel"
	"github.com/MakaroffAV/backend-bootcamp-assignment-2024/internal/service"
)

var house House

func init() {
	c, err := db.Conn()
	if err != nil {
		panic(err)
	}

	repoFlat := dbmodel.NewFlat(c)
	repoToken := dbmodel.NewToken(c)
	repoHouse := dbmodel.NewHouse(c)

	house = NewHouse(
		service.NewHouse(repoFlat, repoToken, repoHouse),
		service.NewToken(repoToken),
	)
}

func TestHouse_Flats(t *testing.T) {

	// неверный номер дома
	w := httptest.NewRecorder()
	r := httptest.NewRequest(http.MethodGet, "/house/abc", nil)

	house.Flats(w, r)
	if w.Code != 400 {
		t.Fatalf(
			"case: 0; got: %d; want: %d; \n", w.Code, 400,
		)
	}
	// неверный номер дома

	// неверный токен запроса
	w = httptest.NewRecorder()
	r = httptest.NewRequest(http.MethodGet, "/house", nil)
	r.SetPathValue("id", "1")
	r.Header.Set("token", "")

	house.Flats(w, r)
	if w.Code != 401 {
		t.Fatalf(
			"case: 1; got: %d; want: %d; \n", w.Code, 401,
		)
	}
	// неверный токен запроса

	// все валидное
	w = httptest.NewRecorder()
	r = httptest.NewRequest(http.MethodGet, "/house", nil)
	r.SetPathValue("id", "1")
	r.Header.Set("token", "test_token_moderator")

	house.Flats(w, r)
	if w.Code != 200 {
		t.Fatalf(
			"case: 1; got: %d; want: %d; \n", w.Code, 200,
		)
	}
	// все валидное
}

func TestHouse_Create(t *testing.T) {

	// неверный токен
	w := httptest.NewRecorder()
	r := httptest.NewRequest(http.MethodPost, "/house", nil)

	house.Create(w, r)
	if w.Code != 401 {
		t.Fatalf(
			"case: 0; got: %d; want: %d; \n", w.Code, 401,
		)
	}
	// неверный токен

	// неверные данные
	w = httptest.NewRecorder()
	r = httptest.NewRequest(http.MethodPost, "/house", nil)
	r.Header.Add("token", "test_token")

	house.Create(w, r)
	if w.Code != 400 {
		t.Fatalf(
			"case: 1; got: %d; want: %d; \n", w.Code, 400,
		)
	}
	// неверные данные

	// верные данные, верный токен
	d := []byte(
		`{"address": "oz_corp", "year": 2000, "developer": "ozborn"}`,
	)
	w = httptest.NewRecorder()
	r = httptest.NewRequest(http.MethodPost, "/house", bytes.NewReader(d))
	r.Header.Add("token", "test_token_moderator")

	house.Create(w, r)
	if w.Code != 200 {
		t.Fatalf(
			"case: 2; got: %d; want: %d; \n", w.Code, 400,
		)
	}
	// верные данные, верный токен

	// верные данные, неверный токен
	d = []byte(
		`{"address": "oz_corp", "year": 2000, "developer": "ozborn"}`,
	)
	w = httptest.NewRecorder()
	r = httptest.NewRequest(http.MethodPost, "/house", bytes.NewReader(d))
	r.Header.Add("token", "test_token")

	house.Create(w, r)
	if w.Code != 500 {
		t.Fatalf(
			"case: 3; got: %d; want: %d; \n", w.Code, 400,
		)
	}
	// верные данные, неверный токен
}
