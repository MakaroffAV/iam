package dbmodel

import (
	"database/sql"

	"github.com/MakaroffAV/backend-bootcamp-assignment-2024/internal/domain/model"
)

const (
	sqlHouseID = `
	select
		id,
		address,
		developer,
		year,
		created,
		updated
	from
		housing.house
	where
		housing.house.id = $1
	`

	sqlHouseAddress = `
	select
		id,
		address,
		developer,
		year,
		created,
		updated
	from
		housing.house
	where
		housing.house.address = $1
	`

	sqlHouseCreate = `
	insert into housing.house (
		year,
		address,
		developer
	) values (
		$1,
		$2,
		$3
	);
	`
)

type House struct {
	db *sql.DB
}

func NewHouse(db *sql.DB) House {
	return House{
		db: db,
	}
}

func (s House) ID(id int) (model.House, error) {
	var m model.House

	err := s.db.QueryRow(
		sqlHouseID, id,
	).Scan(
		&m.ID, &m.Address, &m.Developer,
		&m.Year, &m.Created, &m.Updated,
	)

	return m, err
}

func (s House) address(address string) (model.House, error) {
	var m model.House

	err := s.db.QueryRow(
		sqlHouseAddress, address,
	).Scan(
		&m.ID, &m.Address, &m.Developer,
		&m.Year, &m.Created, &m.Updated,
	)

	return m, err
}

func (s House) Create(
	year int, address string, developer string) (model.House, error) {

	var m model.House

	_, err := s.db.Exec(
		sqlHouseCreate,
		year, address, developer,
	)

	if err != nil {
		return m, err
	} else {
		return s.address(address)
	}
}
