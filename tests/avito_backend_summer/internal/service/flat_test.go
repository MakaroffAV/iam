package service

import (
	"testing"
	"time"

	"github.com/MakaroffAV/backend-bootcamp-assignment-2024/internal/domain/model"
	"github.com/MakaroffAV/backend-bootcamp-assignment-2024/internal/repo/db"
	"github.com/MakaroffAV/backend-bootcamp-assignment-2024/internal/repo/db/dbmodel"
)

func TestFlatStatusCheck(t *testing.T) {
	if r := flatStatusCheck(""); r != false {
		t.Fatalf(
			"case: 1; got: %t; want: %t; \n", r, false,
		)
	}

	if r := flatStatusCheck(model.FlatStatusApproved); r != true {
		t.Fatalf(
			"case: 2; got: %t; want: %t; \n", r, true,
		)
	}
}

func TestFlat_Create(t *testing.T) {
	c, err := db.Conn()
	if err != nil {
		panic(err)
	}
	defer c.Close()

	var (
		f = NewFlat(
			dbmodel.NewFlat(c),
			dbmodel.NewToken(c),
		)
	)

	if _, err := f.Create(
		model.User{}, time.Now().Unix(), 1, 3243, 2,
	); err != nil {
		t.Fatalf(
			"got: %v; want: nil; \n", err,
		)
	}
}

func TestFlatUpdate(t *testing.T) {
	c, err := db.Conn()
	if err != nil {
		panic(err)
	}
	defer c.Close()

	var (
		f = NewFlat(
			dbmodel.NewFlat(c),
			dbmodel.NewToken(c),
		)
	)

	// не должно пройти
	// проверку на смену статуса
	if _, err := f.Update(
		model.User{Type: model.UserTypeModeratorName}, time.Now().UnixMicro(), 1, ""); err == nil {

		t.Fatalf(
			"case: 1; got: %v; want: !nil; \n", err,
		)
	}

	// не должно пройти
	// проверку на тип пользователя
	if _, err := f.Update(
		model.User{Type: model.UserTypeClientName}, time.Now().UnixMicro(), 1, model.FlatStatusApproved); err == nil {

		t.Fatalf(
			"case: 2; got: %v; want: !nill; \n", err,
		)
	}

	time.Sleep(time.Second)

	// попытка изменить квартиру
	// которая находится на модерации
	flatID := time.Now().Unix()
	flatUs := model.User{Type: model.UserTypeModeratorName}

	// создадим новую квартиру
	if _, err := f.Create(flatUs, flatID, 1, 100, 200); err != nil {
		panic(err)
	}

	// переведем квартиру
	// на статус модерации
	// заодно проверим желаемое поведение
	if _, err := f.Update(flatUs, flatID, 1, model.FlatStatusOnModeration); err != nil {
		t.Fatalf(
			"case: 3; got: %v; want: nil; \n", err,
		)
	}

	// попытка изменить квартиру
	// которая находится на модерации
	if _, err := f.Update(
		model.User{Type: model.UserTypeModeratorName}, flatID, 1, model.FlatStatusOnModeration); err == nil {

		t.Fatalf(
			"case: 4; got: %v; want: !nil; \n", err,
		)
	}

}
