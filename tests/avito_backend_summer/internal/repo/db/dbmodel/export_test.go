package dbmodel

import (
	"database/sql"

	"github.com/MakaroffAV/backend-bootcamp-assignment-2024/internal/repo/db"
)

var conn *sql.DB

func init() {
	c, err := db.Conn()
	if err != nil {
		panic(err)
	}
	conn = c
}
