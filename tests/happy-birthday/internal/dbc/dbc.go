// ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++ //

package dbc

// ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++ //

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

// ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++ //

var (
	dbUser = "makarov"
	dbPass = "makarov"
	dbName = "list"
	dbHost = "localhost:3306"
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
