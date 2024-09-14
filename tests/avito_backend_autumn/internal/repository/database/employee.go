package database

import (
	"database/sql"
	"errors"
	"zadanie-6105/internal/domain"
)

const (
	sqlEmployeeGetByUsername = `
	select
		id,
		username,
		first_name,
		last_name,
		created_at,
		updated_at
	from
	    employee
	where
	    username = $1
	`

	sqlEmployeeGetById = `
	select
		id,
		username,
		first_name,
		last_name,
		created_at,
		updated_at
	from
	    employee
	where
	    id = $1
	`
)

type Employee struct {
	db *sql.DB
}

func NewEmployee(db *sql.DB) Employee {
	return Employee{
		db: db,
	}
}

func (e Employee) GetById(i string) (*domain.Employee, error) {
	var r domain.Employee
	err := e.db.QueryRow(
		sqlEmployeeGetById, i,
	).Scan(
		&r.ID, &r.Username, &r.FirstName, &r.LastName, &r.CreatedAt, &r.UpdatedAt,
	)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil
		} else {
			return nil, err
		}
	}
	return &r, nil
}

func (e Employee) GetByUsername(u string) (*domain.Employee, error) {
	var r domain.Employee
	err := e.db.QueryRow(
		sqlEmployeeGetByUsername, u,
	).Scan(
		&r.ID, &r.Username, &r.FirstName, &r.LastName, &r.CreatedAt, &r.UpdatedAt,
	)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil
		} else {
			return nil, err
		}
	}

	return &r, nil
}
