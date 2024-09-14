package database

import (
	"database/sql"
	"zadanie-6105/internal/domain"
)

const (
	sqlBidArchiveGetByIdAndVersion = `
	select
	    bid_id, name, description, status, version, tender_id, author_type, author_id, created
	from
	    bid_archive
	where
	    version = $1
		and
	    bid_id = $2
	`
)

type BidArchive struct {
	db *sql.DB
}

func NewBidArchive(db *sql.DB) BidArchive {
	return BidArchive{
		db: db,
	}
}

func (b BidArchive) GetByIdAndVersion(i string, v int32) (domain.Bid, error) {
	var r domain.Bid
	err := b.db.QueryRow(
		sqlBidArchiveGetByIdAndVersion, v, i,
	).Scan(
		&r.ID, &r.Name, &r.Description, &r.Status, &r.Version, &r.TenderID, &r.AuthorType, &r.AuthorID, &r.CreatedAt,
	)
	return r, err
}
