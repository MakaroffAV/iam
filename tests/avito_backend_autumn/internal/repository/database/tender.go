package database

import (
	"database/sql"
	"github.com/lib/pq"
	"zadanie-6105/internal/domain"
)

const (
	sqlTenderGet = `
	select
		id,
		name,
		description,
		service_type,
		status,
		version,
		created_at
	from
	    tender
	where
	    service_type = any($1::tender_service[])
	limit
		$2
	offset
		$3
	`

	sqlTenderGetByID = `
	select
		id,
		name,
		description,
		service_type,
		status,
		version,
		created_at
	from
	    tender
	where
	    id = $1
	`

	sqlTenderGetFullByID = `
	select 
	    id,
		name,
		description,
		service_type,
		status,
		version,
		created_at,
		employee_id
	from
	    tender
	where
	    id = $1
	`

	sqlTenderCreate = `
	insert into tender (
	    id,
	    name,
	    description,
		service_type,
	    employee_id
	) values (
	    $1,
	    $2,
	    $3,
	    $4,
	    $5
	)
	`

	sqlTenderGetByEmployeeID = `
	select
		id,
		name,
		description,
		service_type,
		status,
		version,
		created_at
	from
	    tender
	where
	    employee_id = $1
	limit 
		$2
	offset 
		$3
	`

	sqlTenderUpdateStatus = `
	update tender set status = $1 where id = $2
	`

	sqlTenderUpdateMoveToArchive = `
	insert into tender_archive (
	    tender_id, name, description, service_type, status, employee_id, version, created_at
	) select id, name, description, service_type, status, employee_id, version, created_at from tender where id = $1
	`

	sqlTenderUpdate = `
	update tender 
		set version = version + 1,
		    name = COALESCE(NULLIF($1, ''), name),
			description = COALESCE(NULLIF($2, ''), description),
			service_type = COALESCE(NULLIF($3, '')::tender_service, service_type)
	where
	    id = $4
	`

	sqlTenderRollback = `
	update tender
	set version = version+1,
	    name = subquery.name,
	    description = subquery.description,
	    service_type = subquery.service_type
	from (select name, description, service_type from tender_archive where tender_id = $1 and version = $2) as subquery
	where id = $3`
)

type Tender struct {
	db *sql.DB
}

func NewTender(db *sql.DB) Tender {
	return Tender{
		db: db,
	}
}

func (t Tender) GetByID(i string) (domain.Tender, error) {
	var r domain.Tender
	err := t.db.QueryRow(
		sqlTenderGetByID, i).Scan(&r.ID, &r.Name, &r.Description, &r.ServiceType, &r.Status, &r.Version, &r.CreatedAt)

	return r, err
}

func (t Tender) GetFullByID(i string) (domain.Tender, error) {
	var r domain.Tender
	err := t.db.QueryRow(
		sqlTenderGetFullByID, i).Scan(&r.ID, &r.Name, &r.Description, &r.ServiceType, &r.Status, &r.Version, &r.CreatedAt, &r.EmployeeID)

	return r, err
}

func (t Tender) Create(i, n, d, k, e string) error {
	_, err := t.db.Exec(sqlTenderCreate, i, n, d, k, e)
	return err
}

func (t Tender) Get(l int32, o int32, s []string) ([]domain.Tender, error) {
	var r = []domain.Tender{}

	q, err := t.db.Query(sqlTenderGet, pq.Array(s), l, o)
	if err != nil {
		return nil, err
	}
	defer q.Close()

	for q.Next() {
		var c domain.Tender
		if err := q.Scan(&c.ID, &c.Name, &c.Description, &c.ServiceType, &c.Status, &c.Version, &c.CreatedAt); err != nil {
			return nil, err
		}
		r = append(r, c)
	}
	return r, nil
}

func (t Tender) GetByEmployeeID(l int32, o int32, i string) ([]domain.Tender, error) {
	var r = []domain.Tender{}
	q, err := t.db.Query(sqlTenderGetByEmployeeID, i, l, o)
	if err != nil {
		return nil, err
	}
	defer q.Close()

	for q.Next() {
		var c domain.Tender
		if err := q.Scan(&c.ID, &c.Name, &c.Description, &c.ServiceType, &c.Status, &c.Version, &c.CreatedAt); err != nil {
			return nil, err
		}
		r = append(r, c)
	}
	return r, nil
}

func (t Tender) UpdateStatus(i, s string) error {
	_, err := t.db.Exec(sqlTenderUpdateStatus, s, i)
	return err
}

func (t Tender) Update(i, n, d, k string) error {
	// начинаем транзакцию
	tx, err := t.db.Begin()
	if err != nil {
		return err
	}

	// Скидываем текущий
	// тендер в врхивную таблицу
	if _, err := tx.Exec(sqlTenderUpdateMoveToArchive, i); err != nil {
		tx.Rollback()
		return err
	}

	// Обновляем тендер в актуальной таблице
	if _, err := tx.Exec(sqlTenderUpdate, n, d, k, i); err != nil {
		tx.Rollback()
		return err
	}

	return tx.Commit()
}

func (t Tender) Rollback(i string, v int32) error {
	_, err := t.db.Exec(sqlTenderRollback, i, v, i)
	return err
}
