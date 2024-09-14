package service

import (
	"testing"

	"github.com/MakaroffAV/backend-bootcamp-assignment-2024/internal/repo/db"
	"github.com/MakaroffAV/backend-bootcamp-assignment-2024/internal/repo/db/dbmodel"
)

func TestLogin_DummyLogin(t *testing.T) {
	c, err := db.Conn()
	if err != nil {
		panic(err)
	}
	defer c.Close()

	var (
		l = NewLogin(
			dbmodel.NewUser(c),
			dbmodel.NewToken(c),
		)
	)
	if _, err := l.DummyLogin("client"); err != nil {
		t.Fatalf(
			"got: %v; want: nil; \n", err,
		)
	}
}

func TestLogin_Login(t *testing.T) {
	c, err := db.Conn()
	if err != nil {
		panic(err)
	}
	defer c.Close()

	var (
		l = NewLogin(
			dbmodel.NewUser(c),
			dbmodel.NewToken(c),
		)
	)

	if _, err := l.Login(1); err != nil {
		t.Fatalf(
			"got: %v; want: nil; \n", err,
		)
	}
}

func TestLogin_UserExists(t *testing.T) {
	c, err := db.Conn()
	if err != nil {
		panic(err)
	}
	defer c.Close()

	var (
		l = NewLogin(
			dbmodel.NewUser(c),
			dbmodel.NewToken(c),
		)
	)

	if _, err := l.UserExists("test_uuid", "test_password"); err != nil {
		t.Fatalf(
			"got: %v; want: nil; \n", err,
		)
	}
}
