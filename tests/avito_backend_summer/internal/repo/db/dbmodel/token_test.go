package dbmodel

import (
	"errors"
	"reflect"
	"testing"
	"time"

	"github.com/MakaroffAV/backend-bootcamp-assignment-2024/internal/domain/model"
	"github.com/google/uuid"
)

func TestToken_Get(t *testing.T) {
	var (
		tk = NewToken(conn)

		exp = struct {
			e error
			m model.Token
		}{
			e: nil,
			m: model.Token{
				ID:      1,
				Value:   "test_token",
				UserID:  1,
				Created: time.Date(2024, 8, 8, 22, 0, 0, 0, time.FixedZone("", 0)),
			},
		}
	)

	v, err := tk.Get("test_token")
	if !reflect.DeepEqual(v, exp.m) || !errors.Is(err, exp.e) {
		t.Fatalf(
			"got: (%v, %v); want: (%v, %v); \n",
			v,
			err,
			exp.m,
			exp.e,
		)
	}
}

func TestToken_User(t *testing.T) {
	var (
		tk = NewToken(conn)

		exp = struct {
			e error
			m model.User
		}{
			e: nil,
			m: model.User{
				ID:       1,
				Type:     "client",
				Uuid:     "test_uuid",
				Email:    "test_email",
				Password: "test_password",
				Dummy:    false,
				Created:  time.Date(2024, 8, 8, 21, 30, 0, 0, time.FixedZone("", 0)),
			},
		}
	)

	u, err := tk.User("test_token")
	if !reflect.DeepEqual(u, exp.m) || !errors.Is(err, exp.e) {
		t.Fatalf(
			"got: (%v, %v); want: (%v, %v); \n",
			u,
			err,
			exp.m,
			exp.e,
		)
	}
}

func TestToken_Create(t *testing.T) {
	var (
		tk = NewToken(conn)

		exp error = nil
		arg       = struct {
			u int
			t string
		}{
			u: 1,
			t: uuid.New().String(),
		}
	)

	_, err := tk.Create(arg.u, arg.t)
	if err != exp {
		t.Fatalf(
			"got: %v; want: %v; \n", err, exp,
		)
	}
}
