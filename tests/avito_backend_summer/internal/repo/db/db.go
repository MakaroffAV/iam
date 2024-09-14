package db

import (
	"database/sql"

	"github.com/MakaroffAV/backend-bootcamp-assignment-2024/internal/config"

	_ "github.com/lib/pq"
)

func Conn() (*sql.DB, error) {
	c, err := sql.Open(
		"postgres",
		config.DbConnString,
	)

	if err != nil {
		return nil, err
	} else {
		return c, c.Ping()
	}
}
