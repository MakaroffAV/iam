package dbmodel

import (
	"errors"
	"reflect"
	"testing"
	"time"

	"github.com/MakaroffAV/backend-bootcamp-assignment-2024/internal/domain/model"
)

func TestUser_Get(t *testing.T) {
	var (
		u = NewUser(conn)

		arg = "test_uuid"
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

	m, err := u.Get(arg)
	if !reflect.DeepEqual(m, exp.m) || !errors.Is(err, exp.e) {
		t.Fatalf(
			"got: (%v, %v); want: (%v, %v); \n",
			m,
			err,
			exp.m,
			exp.e,
		)
	}
}

func TestUser_Create(t *testing.T) {
	var (
		u = NewUser(conn)

		arg = struct {
			userType, uuid, email, password string
			dummy                           bool
		}{
			userType: "moderator",
			uuid:     "test_uuid",
			email:    "test_email",
			password: "test_password",
			dummy:    false,
		}
	)

	_, err := u.Create(
		arg.userType, arg.uuid, arg.email, arg.password, arg.dummy,
	)
	if err != nil {
		t.Fatalf(
			"got: %v; want: nil; \n", err,
		)
	}
}

func TestUser_Credentials(t *testing.T) {
	var (
		u = NewUser(conn)

		arg = struct {
			u string
			p string
		}{
			u: "test_uuid",
			p: "test_password",
		}

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

	m, err := u.Credentials(arg.u, arg.p)
	if !reflect.DeepEqual(m, exp.m) || !errors.Is(err, exp.e) {
		t.Fatalf(
			"got: (%v, %v); want: (%v, %v); \n",
			m,
			err,
			exp.m,
			exp.e,
		)
	}
}
