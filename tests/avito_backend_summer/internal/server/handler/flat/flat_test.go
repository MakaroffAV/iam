package flat

import (
	"bytes"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/MakaroffAV/backend-bootcamp-assignment-2024/internal/repo/db"
	"github.com/MakaroffAV/backend-bootcamp-assignment-2024/internal/repo/db/dbmodel"
	"github.com/MakaroffAV/backend-bootcamp-assignment-2024/internal/service"
)

var flat Flat

func init() {
	c, err := db.Conn()
	if err != nil {
		panic(err)
	}

	repoFlat := dbmodel.NewFlat(c)
	repoToken := dbmodel.NewToken(c)

	flat = NewFlat(
		service.NewToken(repoToken),
		service.NewFlat(repoFlat, repoToken),
	)
}

func TestFlat_Update(t *testing.T) {

	// невалидный токен
	w := httptest.NewRecorder()
	r := httptest.NewRequest(http.MethodPost, "/flat/update", nil)

	r.Header.Set("token", "")

	flat.Update(w, r)
	if w.Code != 401 {
		t.Fatalf(
			"case: 1; got: %d; want: %d; \n", w.Code, 401,
		)
	}
	// невалидный токен

	// невалидные данные
	w = httptest.NewRecorder()
	r = httptest.NewRequest(http.MethodPost, "/flat/update", nil)

	r.Header.Set("token", "test_token_moderator")

	flat.Update(w, r)
	if w.Code != 400 {
		t.Fatalf(
			"case: 2; got: %d; want: %d; \n", w.Code, 400,
		)
	}
	// невалидные данные

	// валидные данные, успешный запрос
	d := []byte(
		`{"id": 1, "house_id": 1, "status": "on moderation"}`,
	)
	w = httptest.NewRecorder()
	r = httptest.NewRequest(http.MethodPost, "/flat/update", bytes.NewBuffer(d))
	r.Header.Set("token", "test_token_moderator")

	flat.Update(w, r)
	if w.Code != 200 {
		t.Fatalf(
			"case: 3; got: %d; want: %d; \n", w.Code, 200,
		)
	}
	// валидные данные, успешный запрос

	// валидные данные, неуспешный запрос
	d = []byte(
		`{"id": 1, "house_id": 1, "status": "on moderation"}`,
	)
	w = httptest.NewRecorder()
	r = httptest.NewRequest(http.MethodPost, "/flat/update", bytes.NewBuffer(d))
	r.Header.Set("token", "test_token_moderator")

	flat.Update(w, r)
	if w.Code != 500 {
		t.Fatalf(
			"case: 4; got: %d; want: %d; \n", w.Code, 200,
		)
	}
	// валидные данные, неуспешный запрос

	// возвращаем исходные данные
	d = []byte(
		`{"id": 1, "house_id": 1, "status": "created"}`,
	)
	w = httptest.NewRecorder()
	r = httptest.NewRequest(http.MethodPost, "/flat/update", bytes.NewBuffer(d))
	r.Header.Set("token", "test_token_moderator")

	flat.Update(w, r)
	if w.Code == 500 {
		panic("not valid request")
	}
	// возвращаем исходные данные
}

func TestFlat_Create(t *testing.T) {
	// невалидный токен
	w := httptest.NewRecorder()
	r := httptest.NewRequest(http.MethodPost, "/flat/create", nil)

	r.Header.Set("token", "")

	flat.Create(w, r)
	if w.Code != 401 {
		t.Fatalf(
			"case: 1; got: %d; want: %d; \n", w.Code, 401,
		)
	}
	// невалидный токен

	// невалидные данные
	w = httptest.NewRecorder()
	r = httptest.NewRequest(http.MethodPost, "/flat/create", nil)

	r.Header.Set("token", "test_token_moderator")

	flat.Create(w, r)
	if w.Code != 400 {
		t.Fatalf(
			"case: 2; got: %d; want: %d; \n", w.Code, 400,
		)
	}
	// невалидные данные

	flatID := time.Now().UnixMicro()

	// валидные данные, успешный запрос
	d := []byte(
		fmt.Sprintf(
			`{"id": %d, "house_id": 1, "price": 800, "rooms": 300}`,
			flatID,
		),
	)
	w = httptest.NewRecorder()
	r = httptest.NewRequest(http.MethodPost, "/flat/create", bytes.NewBuffer(d))
	r.Header.Set("token", "test_token_moderator")

	flat.Create(w, r)
	if w.Code != 200 {
		t.Fatalf(
			"case: 3; got: %d; want: %d; \n", w.Code, 200,
		)
	}
	// валидные данные, успешный запрос

	// валидные данные, неуспешный запрос (дубль)
	d = []byte(
		fmt.Sprintf(
			`{"id": %d, "house_id": 1, "price": 800, "rooms": 300}`,
			flatID,
		),
	)
	w = httptest.NewRecorder()
	r = httptest.NewRequest(http.MethodPost, "/flat/create", bytes.NewBuffer(d))
	r.Header.Set("token", "test_token_moderator")

	flat.Create(w, r)
	if w.Code != 500 {
		t.Fatalf(
			"case: 3; got: %d; want: %d; \n", w.Code, 200,
		)
	}
	// валидные данные, неуспешный запрос (дубль)

}
