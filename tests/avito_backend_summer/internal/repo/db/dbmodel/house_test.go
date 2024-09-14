package dbmodel

import (
	"errors"
	"reflect"
	"testing"
	"time"

	"github.com/MakaroffAV/backend-bootcamp-assignment-2024/internal/domain/model"
)

func TestHouse_Address(t *testing.T) {
	var (
		h = NewHouse(conn)

		exp = struct {
			e error
			m model.House
		}{
			e: nil,
			m: model.House{
				ID:        1,
				Address:   "test address",
				Developer: "test developer",
				Year:      2001,
				Created:   time.Date(2024, 8, 8, 15, 30, 0, 0, time.FixedZone("", 0)),
				Updated:   time.Date(2024, 8, 8, 15, 30, 0, 0, time.FixedZone("", 0)),
			},
		}
	)

	m, err := h.address("test address")
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

func TestHouse_ID(t *testing.T) {
	var (
		h = NewHouse(conn)

		exp = struct {
			e error
			m model.House
		}{
			e: nil,
			m: model.House{
				ID:        1,
				Address:   "test address",
				Developer: "test developer",
				Year:      2001,
				Created:   time.Date(2024, 8, 8, 15, 30, 0, 0, time.FixedZone("", 0)),
				Updated:   time.Date(2024, 8, 8, 15, 30, 0, 0, time.FixedZone("", 0)),
			},
		}
	)

	m, err := h.ID(1)
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

func TestHouse_Create(t *testing.T) {
	var (
		h = NewHouse(conn)

		exp error = nil
		arg       = struct {
			year      int
			address   string
			developer string
		}{
			year:      2000,
			address:   "test address create",
			developer: "test developer create",
		}
	)

	_, err := h.Create(
		arg.year, arg.address, arg.developer,
	)
	if !errors.Is(err, exp) {
		t.Fatalf(
			"got: %v; want: %v; \n",
			err,
			exp,
		)
	}
}
