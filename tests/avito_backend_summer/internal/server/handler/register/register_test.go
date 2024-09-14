package register

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/MakaroffAV/backend-bootcamp-assignment-2024/internal/repo/db"
	"github.com/MakaroffAV/backend-bootcamp-assignment-2024/internal/repo/db/dbmodel"
	"github.com/MakaroffAV/backend-bootcamp-assignment-2024/internal/service"
)

var register Register

func init() {
	c, err := db.Conn()
	if err != nil {
		panic(err)
	}

	repoUser := dbmodel.NewUser(c)

	register = NewRegister(
		service.NewRegister(repoUser),
	)
}

func TestRegister_Do(t *testing.T) {
	// неверные входные данные
	w := httptest.NewRecorder()
	r := httptest.NewRequest(http.MethodPost, "/register", nil)

	register.Do(w, r)
	if w.Code != 400 {
		t.Fatalf(
			"case: 0; got: %d; want: %d; \n", w.Code, 400,
		)
	}
	// неверные входные данные

	// верные входные данные (желаемое поведение)
	d := []byte(
		`{"email": "test@test.com", "password": "test_test_password", "user_type": "test_test_user_type"}`,
	)
	w = httptest.NewRecorder()
	r = httptest.NewRequest(http.MethodPost, "/register", bytes.NewBuffer(d))

	register.Do(w, r)
	if w.Code != 200 {
		t.Fatalf(
			"case: 0; got: %d; want: %d; \n", w.Code, 200,
		)
	}
	// верные входные данные (желаемое поведение)
}
