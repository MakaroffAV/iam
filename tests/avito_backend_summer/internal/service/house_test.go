package service

import (
	"errors"
	"testing"

	"github.com/MakaroffAV/backend-bootcamp-assignment-2024/internal/domain/model"
	"github.com/MakaroffAV/backend-bootcamp-assignment-2024/internal/repo/db"
	"github.com/MakaroffAV/backend-bootcamp-assignment-2024/internal/repo/db/dbmodel"
)

func TestHouse_Get(t *testing.T) {
	c, err := db.Conn()
	if err != nil {
		panic(err)
	}
	defer c.Close()

	var (
		h = NewHouse(
			dbmodel.NewFlat(c),
			dbmodel.NewToken(c),
			dbmodel.NewHouse(c),
		)
	)

	if _, err := h.Get(model.User{}, 1); err == nil {
		t.Fatalf(
			"case: 1; got: %v; want: !nil; \n", err,
		)
	}

	if _, err := h.Get(model.User{Type: model.UserTypeClientName}, 1); err != nil {
		t.Fatalf(
			"case: 2; got: %v; want: nil; \n", err,
		)
	}

	if _, err := h.Get(model.User{Type: model.UserTypeModeratorName}, 1); err != nil {
		t.Fatalf(
			"case: 3; got: %v; want: nil; \n", err,
		)
	}
}

func TestHouse_Create(t *testing.T) {
	c, err := db.Conn()
	if err != nil {
		panic(err)
	}
	defer c.Close()

	var (
		h = NewHouse(
			dbmodel.NewFlat(c),
			dbmodel.NewToken(c),
			dbmodel.NewHouse(c),
		)

		arg = struct {
			address   string
			year      int
			developer string
		}{
			address:   "forest str",
			year:      1999,
			developer: "oz_corp",
		}
	)

	// клиент не может создавать дом
	if _, err := h.Create(
		model.User{Type: model.UserTypeClientName}, arg.address, arg.year, arg.developer); err == nil {

		t.Fatalf(
			"case: 1; got: %v; want: %v; \n", err, errors.New("пользователь не модератор"),
		)
	}

	// модератор может создать дом
	if _, err := h.Create(
		model.User{Type: model.UserTypeModeratorName}, arg.address, arg.year, arg.developer); err != nil {

		t.Fatalf(
			"case: 1; got: %v; want: nil; \n", err,
		)

	}

}
