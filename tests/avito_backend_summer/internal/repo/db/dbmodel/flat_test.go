package dbmodel

import (
	"errors"
	"reflect"
	"testing"
	"time"

	"github.com/MakaroffAV/backend-bootcamp-assignment-2024/internal/domain/model"
)

func TestFlat_ID(t *testing.T) {
	var (
		f = NewFlat(conn)

		exp = struct {
			e error
			m model.Flat
		}{
			e: nil,
			m: model.Flat{
				ID:      int64(1),
				Price:   192837465,
				Rooms:   2,
				HouseID: 1,
				Status:  "created",
			},
		}
	)

	m, err := f.id(exp.m.ID, exp.m.HouseID)
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

func TestFlat_Create(t *testing.T) {
	time.Sleep(time.Second)
	var (
		f = NewFlat(conn)
		i = time.Now().UnixMicro()

		exp = struct {
			e error
			m model.Flat
		}{
			e: nil,
			m: model.Flat{
				ID:      i,
				Price:   192837465,
				Rooms:   2,
				HouseID: 1,
				Status:  "created",
			},
		}
	)

	m, err := f.Create(
		i, 1, 192837465, 2, time.Date(2024, 8, 8, 15, 30, 0, 0, time.FixedZone("", 0)),
	)
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

func TestFlat_Update(t *testing.T) {
	// странный баг, если принудительно не
	// заснуть  -  ошибка  с дублированием
	//
	// по идее логично, одна и  та же сек.
	time.Sleep(time.Second)

	var (
		f = NewFlat(conn)
		i = time.Now().UnixMicro()

		exp = struct {
			e error
			m model.Flat
		}{
			e: nil,
			m: model.Flat{
				ID:      i,
				Price:   192837465,
				Rooms:   2,
				HouseID: 1,
				Status:  "test status",
			},
		}
	)

	m, err := f.Create(
		i, 1, 192837465, 2, time.Date(2024, 8, 8, 15, 30, 0, 0, time.FixedZone("", 0)),
	)
	if err != nil {
		panic(err)
	}

	u, err := f.Update(m.ID, m.HouseID, "test status")
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

func TestFlat_House(t *testing.T) {
	var (
		f = NewFlat(conn)
	)

	m, err := f.House(1, true)
	if err != nil || len(m) == 0 {
		t.Fatalf(
			"got: (%v, %v); want: (>0, nil); \n",
			len(m), err,
		)
	}

	if _, err := f.House(1, false); err != nil {
		t.Fatalf(
			"got: (%v); want: (nil); \n", err,
		)
	}
}
