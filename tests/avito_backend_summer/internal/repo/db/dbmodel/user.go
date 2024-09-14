package dbmodel

import (
	"database/sql"

	"github.com/MakaroffAV/backend-bootcamp-assignment-2024/internal/domain/model"
)

var (
	sqlUserGet = `
	select
		id,
		type,
		uuid,
		email,
		dummy,
		created,
		password
	from
		housing.user
	where
		housing.user.uuid = $1
	`

	sqlUserCreate = `
	insert into housing.user (
		type,
		uuid,
		email,
		password,
		dummy
	) values (
		$1,
		$2,
		$3,
		$4,
		$5
	)
	`

	sqlUserCredentials = `
	select
		id,
		type,
		uuid,
		email,
		dummy,
		created,
		password
	from
		housing.user
	where
		housing.user.dummy = false
		and
		housing.user.uuid = $1
		and
		housing.user.password = $2
	`
)

type User struct {
	db *sql.DB
}

func NewUser(db *sql.DB) User {
	return User{
		db: db,
	}
}

func (u User) Get(uuid string) (model.User, error) {
	var m model.User

	err := u.db.QueryRow(sqlUserGet, uuid).Scan(
		&m.ID, &m.Type, &m.Uuid, &m.Email,
		&m.Dummy, &m.Created, &m.Password,
	)
	return m, err
}

func (u User) Credentials(uuid, password string) (model.User, error) {
	var m model.User

	err := u.db.QueryRow(
		sqlUserCredentials, uuid, password,
	).Scan(
		&m.ID, &m.Type, &m.Uuid, &m.Email,
		&m.Dummy, &m.Created, &m.Password,
	)
	return m, err
}

func (u User) Create(
	userType, uuid, email, password string, dummy bool) (model.User, error) {

	var m model.User

	_, err := u.db.Exec(
		sqlUserCreate,
		userType, uuid, email, password, dummy,
	)
	if err != nil {
		return m, err
	} else {
		return u.Get(uuid)
	}
}
