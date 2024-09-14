package login

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/MakaroffAV/backend-bootcamp-assignment-2024/internal/repo/db"
	"github.com/MakaroffAV/backend-bootcamp-assignment-2024/internal/repo/db/dbmodel"
	"github.com/MakaroffAV/backend-bootcamp-assignment-2024/internal/service"
)

var login Login

func init() {
	c, err := db.Conn()
	if err != nil {
		panic(err)
	}

	repoUser := dbmodel.NewUser(c)
	repoToken := dbmodel.NewToken(c)

	login = NewLogin(
		service.NewLogin(repoUser, repoToken),
	)
}

func TestLogin_Login(t *testing.T) {

	// неверные данные
	w := httptest.NewRecorder()
	r := httptest.NewRequest(http.MethodPost, "/login", nil)

	login.Login(w, r)
	if w.Code != 400 {
		t.Fatalf(
			"case: 0; got: %d; want: %d; \n", w.Code, 400,
		)
	}
	// неверные данные

	// неверный пользователь и пароль
	d := []byte(
		`{"id": "", "password": ""}`,
	)
	w = httptest.NewRecorder()
	r = httptest.NewRequest(http.MethodPost, "/login", bytes.NewBuffer(d))

	login.Login(w, r)
	if w.Code != 404 {
		t.Fatalf(
			"case: 1; got: %d; want: %d; \n", w.Code, 404,
		)
	}
	// неверный пользователь и пароль

	// все верное (желаемое поведение)
	d = []byte(
		`{"id": "test_uuid", "password": "test_password"}`,
	)

	w = httptest.NewRecorder()
	r = httptest.NewRequest(http.MethodPost, "/login", bytes.NewBuffer(d))

	login.Login(w, r)
	if w.Code != 200 {
		t.Fatalf(
			"case: 2; got: %d; want: %d; \n", w.Code, 200,
		)
	}
	// все верное (желаемое поведение)

}

func TestLogin_DummyLogin(t *testing.T) {

	// верный тип пользователя
	w := httptest.NewRecorder()
	r := httptest.NewRequest(http.MethodGet, "/dummyLogin?user_type=client", nil)

	login.DummyLogin(w, r)
	if w.Code != 200 {
		t.Fatalf(
			"case: 0; got: %d; want: %d; \n", w.Code, 200,
		)
	}
	// верный тип пользователя

	time.Sleep(time.Second)

	// неверный тип пользователя
	w = httptest.NewRecorder()
	r = httptest.NewRequest(http.MethodGet, "/dummyLogin?user_type=x", nil)

	login.DummyLogin(w, r)
	if w.Code != 500 {
		t.Fatalf(
			"case: 1; got: %d; want: %d; \n", w.Code, 500,
		)
	}
	// неверный тип пользователя
}
