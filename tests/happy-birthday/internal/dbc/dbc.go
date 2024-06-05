// ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++ //

package dbc

// ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++ //

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/go-sql-driver/mysql"
)

// ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++ //

var (
	dbUser = os.Getenv("DB_USER")
	dbPass = os.Getenv("DB_PASS")
	dbName = os.Getenv("DB_NAME")
	dbHost = os.Getenv("DB_ADDR")
)

// ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++ //

// Connection is function to
// create  instance  of  the
// MySQL database connection
func Connection() (*sql.DB, error) {

	c, cErr := sql.Open(
		"mysql",
		fmt.Sprintf(
			"%s:%s@tcp(%s)/%s",
			dbUser,
			dbPass,
			dbHost,
			dbName,
		),
	)
	if cErr != nil {
		return nil, cErr
	}

	if pErr := c.Ping(); pErr != nil {
		return nil, pErr
	}

	return c, nil

}

// ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++ //
