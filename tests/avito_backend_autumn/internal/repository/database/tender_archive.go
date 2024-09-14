package database

import (
	"database/sql"
	"zadanie-6105/internal/domain"
)

const (
	sqlTenderArchiveGetByIdAndVersion = `
	select
		tender_id,
		name,
		description,
		service_type,
		status,
		version,
		created_at,
		employee_id
	from
	    tender_archive
	where
	    version = $1
	    and
	    tender_id = $2
	`
)

type TenderArchive struct {
	db *sql.DB
}

func NewTenderArchive(db *sql.DB) TenderArchive {
	return TenderArchive{
		db: db,
	}
}

func (t TenderArchive) GetByIdAndVersion(i string, v int32) (domain.Tender, error) {
	var r domain.Tender
	err := t.db.QueryRow(
		sqlTenderArchiveGetByIdAndVersion, v, i).Scan(&r.ID, &r.Name, &r.Description, &r.ServiceType, &r.Status, &r.Version, &r.CreatedAt, &r.EmployeeID)

	return r, err
}
