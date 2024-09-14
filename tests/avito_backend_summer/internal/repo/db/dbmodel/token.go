package dbmodel

import (
	"database/sql"

	"github.com/MakaroffAV/backend-bootcamp-assignment-2024/internal/domain/model"
)

const (
	sqlTokenGet = `
	select
		id,
		value,
		user_id,
		created
	from
		housing.token
	where
		housing.token.value = $1
	`

	sqlTokenUser = `
	select
		t2.id,
		t2.type,
		t2.uuid,
		t2.email,
		t2.password,
		t2.dummy,
		t2.created
	from
		housing.token as t1
	left join
		housing.user as t2
	on
		t1.user_id = t2.id
	where
		t1.value = $1
	`

	sqlTokenCreate = `
	insert into housing.token (
		value,
		user_id
	) values (
		$1,
		$2
	);
	`
)

type Token struct {
	db *sql.DB
}

func NewToken(db *sql.DB) Token {
	return Token{
		db: db,
	}
}

func (t Token) Get(token string) (model.Token, error) {
	var m model.Token

	err := t.db.QueryRow(sqlTokenGet, token).Scan(
		&m.ID, &m.Value, &m.UserID, &m.Created,
	)
	return m, err
}

func (t Token) User(token string) (model.User, error) {
	var m model.User
	err := t.db.QueryRow(sqlTokenUser, token).Scan(
		&m.ID, &m.Type, &m.Uuid, &m.Email,
		&m.Password, &m.Dummy, &m.Created,
	)
	return m, err
}

func (t Token) Create(userID int, token string) (model.Token, error) {
	var m model.Token

	_, err := t.db.Exec(
		sqlTokenCreate, token, userID,
	)
	if err != nil {
		return m, err
	} else {
		return t.Get(token)
	}
}
