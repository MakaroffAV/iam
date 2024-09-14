package database

import "database/sql"

const (
	sqlOrgResponsibleExists = `
	select exists (
		select
			id 
		from 
		    organization_responsible
		where
		    organization_id = $1 and user_id = $2
	)
	`
)

type OrgResponsible struct {
	db *sql.DB
}

func NewOrgResponsible(db *sql.DB) OrgResponsible {
	return OrgResponsible{
		db: db,
	}
}

func (o OrgResponsible) Exists(orgID, userID string) (bool, error) {
	var e bool
	err := o.db.QueryRow(
		sqlOrgResponsibleExists, orgID, userID).Scan(&e)
	return e, err
}
