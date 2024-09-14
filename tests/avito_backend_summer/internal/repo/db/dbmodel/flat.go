package dbmodel

import (
	"database/sql"
	"fmt"
	"time"

	"github.com/MakaroffAV/backend-bootcamp-assignment-2024/internal/domain/model"
)

const (
	sqlFlatID = `
	select
		id,
		price,
		rooms,
		status,
		house_id
	from
		housing.flat
	where
		housing.flat.id = $1
		and
		housing.flat.house_id = $2
	`

	sqlFlatCreate = `
	insert into housing.flat (
		id,
		price,
		rooms,
		house_id
	) values (
		$1,
		$2,
		$3,
		$4
	);
	`

	sqlFlatHouseAll = `
	select
		id,
		price,
		rooms,
		status,
		house_id
	from
		housing.flat
	where
		housing.flat.house_id = $1
	`

	sqlFlatHouseApproved = `
	select
		id,
		price,
		rooms,
		status,
		house_id
	from
		housing.flat
	where
		housing.flat.house_id = $1
		and
		housing.flat.status = 'approved'
	`

	sqlFlatUpdate = `
	update
		housing.flat
	set
		status = $1
	where
		housing.flat.id = $2
		and
		housing.flat.house_id = $3
	`

	sqlFlatHouseUpdate = `
	update 
		housing.house
	set
		updated = $1
	where 
		housing.house.id = $2
	`
)

type Flat struct {
	db *sql.DB
}

func NewFlat(db *sql.DB) Flat {
	return Flat{
		db: db,
	}
}

func (s Flat) id(id int64, houseID int) (model.Flat, error) {
	var m model.Flat

	err := s.db.QueryRow(
		sqlFlatID, id, houseID,
	).Scan(
		&m.ID, &m.Price,
		&m.Rooms, &m.Status, &m.HouseID,
	)

	return m, err
}

func (s Flat) Get(
	id int64, houseID int) (model.Flat, error) {

	return s.id(id, houseID)
}

func (s Flat) House(
	houseID int, all bool) ([]model.Flat, error) {

	var (
		stmt string
		m    []model.Flat
	)

	stmt = sqlFlatHouseAll
	if !all {
		stmt = sqlFlatHouseApproved
	}

	q, err := s.db.Query(
		stmt, houseID,
	)
	if err != nil {
		return m, err
	}
	defer q.Close()

	for q.Next() {
		var r model.Flat

		err := q.Scan(
			&r.ID, &r.Price,
			&r.Rooms, &r.Status, &r.HouseID,
		)
		if err != nil {
			return m, err
		} else {
			m = append(m, r)
		}
	}

	return m, nil
}

func (s Flat) Update(
	id int64, houseID int, status string) (model.Flat, error) {

	// В opeapi написано, что обновление
	// статуса  квартиры  происходит  по
	// роуту /flat/update, передается id
	// квартиры
	//
	// В то же время в доке написано, что
	// id квартиры не является уникальным
	//
	// я прошу при обновлениии квартиры от
	// пользователя,  кроме  статуса, и id
	// квартиры, и id дома

	var m model.Flat

	_, err := s.db.Exec(
		sqlFlatUpdate,
		status, id, houseID,
	)
	if err != nil {
		return m, err
	} else {
		return s.id(id, houseID)
	}
}

func (s Flat) Create(
	id int64, houseID int, price int, rooms int, n time.Time) (model.Flat, error) {

	// по требованиям - при создании квартиры надо обновить поле updated в
	// таблице house. Для этого возпользуемся механизмом транзации, вообще
	// можно было написать триггер, но сделаю так, чтобы вся логика
	// приложения лежала в приложении

	var m model.Flat

	// начинаем транзакцию
	t, err := s.db.Begin()
	if err != nil {
		return m, err
	}

	// создаем квартиру
	if _, err := t.Exec(
		sqlFlatCreate,
		id, price, rooms, houseID,
	); err != nil {
		t.Rollback()
		return m, err
	}

	fmt.Println("here")

	// обновляем поле
	// updated у дома
	if _, err := t.Exec(
		sqlFlatHouseUpdate, n, houseID,
	); err != nil {
		t.Rollback()
		return m, err
	}

	// пишем, что надо отправить
	// уведомления пользователям
	// todo

	// завершаем транзакцию
	if err := t.Commit(); err != nil {
		return m, err
	}

	return s.id(id, houseID)
}
