package database

import (
	"database/sql"
	"zadanie-6105/internal/domain"
)

const (
	sqlBidGetByID = `
	select
		id,
		name,
		description,
		status,
		version,
		tender_id,
		author_type,
		author_id
	from
	    bid
	where
	    id = $1
	`

	sqlBidCreate = `
	insert into bid (
	    id, name, description, tender_id, author_type, author_id
	) values (
	    $1, $2, $3, $4, $5, $6
	) 
	`

	sqlBidGetByUserID = `
	select
		id,
		name,
		description,
		status,
		version,
		tender_id,
		author_type,
		author_id
	from
	    bid
	where
	    author_id = $1
	limit
		$2
	offset
		$3
	`

	sqlBidGetByIdAndEmployeeID = `
	select
		id,
		name,
		description,
		status,
		version,
		tender_id,
		author_type,
		author_id
	from
	    bid
	where
	    tender_id = $1 and author_id = $2
	limit
		$3
	offset
		$4
	`

	sqlBidUpdateStatus = `
	update bid set status = $1 where id = $2
	`

	sqlBidEditMoveToArchive = `
	insert into bid_archive (
	    bid_id, name, description, status, version, tender_id, author_type, author_id, created
	) select id, name, description, status, version, tender_id, author_type, author_id, created from bid where id = $1`

	sqlBidEdit = `
	update bid
		set version = version + 1,
			name = COALESCE(NULLIF($1, ''), name),
			description = COALESCE(NULLIF($2, ''), name)
	where
	    id = $3
	`

	sqlBidRollback = `
	update bid
	    set	version = version + 1,
	    	name = subquery.name,
	    	description = subquery.description
	from (select name, description from bid_archive where bid_id = $1 and version = $2) as subquery
	where id = $3
	`
)

type Bid struct {
	db *sql.DB
}

func NewBid(db *sql.DB) Bid {
	return Bid{
		db: db,
	}
}

func (b Bid) GetByID(i string) (domain.Bid, error) {
	var r domain.Bid
	err := b.db.QueryRow(
		sqlBidGetByID, i,
	).Scan(
		&r.ID, &r.Name, &r.Description, &r.Status, &r.Version, &r.TenderID, &r.AuthorType, &r.AuthorID,
	)
	return r, err

}

func (b Bid) Create(i, n, d, t, at, ai string) error {
	_, err := b.db.Exec(sqlBidCreate, i, n, d, t, at, ai)
	return err
}

func (b Bid) GetByUserID(l, o int32, u string) ([]domain.Bid, error) {
	var r = []domain.Bid{}
	q, err := b.db.Query(sqlBidGetByUserID, u, l, o)
	if err != nil {
		return r, err
	}
	defer q.Close()

	for q.Next() {
		var c domain.Bid
		sErr := q.Scan(
			&c.ID, &c.Name, &c.Description, &c.Status, &c.Version, &c.TenderID, &c.AuthorType, &c.AuthorID,
		)
		if sErr != nil {
			return r, sErr
		}
		r = append(r, c)
	}
	return r, nil
}

func (b Bid) GetByIdAndEmployeeID(l, o int32, tID, uID string) ([]domain.Bid, error) {
	var r = []domain.Bid{}
	q, err := b.db.Query(sqlBidGetByIdAndEmployeeID, tID, uID, l, o)
	if err != nil {
		return r, err
	}
	defer q.Close()

	for q.Next() {
		var c domain.Bid
		sErr := q.Scan(
			&c.ID, &c.Name, &c.Description, &c.Status, &c.Version, &c.TenderID, &c.AuthorType, &c.AuthorID,
		)
		if sErr != nil {
			return r, sErr
		}
		r = append(r, c)
	}
	return r, nil
}

func (b Bid) UpdateStatus(id string, s string) error {
	_, err := b.db.Exec(sqlBidUpdateStatus, s, id)
	return err
}

func (b Bid) Edit(id, bn, bd string) error {
	// начинаем транзакцию
	tx, err := b.db.Begin()
	if err != nil {
		return err
	}

	// Скидываем текущее
	// предложение в архив
	if _, err := tx.Exec(sqlBidEditMoveToArchive, id); err != nil {
		tx.Rollback()
		return err
	}

	// Обновляем предложение
	// в актуальной  таблице
	if _, err := tx.Exec(sqlBidEdit, bn, bd, id); err != nil {
		tx.Rollback()
		return err
	}
	return tx.Commit()

}

func (b Bid) Rollback(i string, v int32) error {
	_, err := b.db.Exec(
		sqlBidRollback, i, v, i,
	)
	return err
}
