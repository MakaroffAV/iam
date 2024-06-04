// ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++ //

package models

// ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++ //

import "database/sql"

// ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++ //

// UserCode is data structure
// to describe meta  data  of
// the     user_code    table
type UserCode struct {
	Id    int
	Mail  string
	Code  int64
	Token string
}

// ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++ //

// Insert is method of  the UserCode
// data structure to create new  row
// into the user_code database table
func (c UserCode) Insert(dbConn *sql.DB) error {

	var (

		// s defines the template
		// of  the  SQL statement
		s = `
		INSERT INTO user_code (
			mail,
			code,
			token
		) VALUES (
			?,
			?,
			?
		);
		`
	)

	// execute  SQL  statement to
	// create  new  row into  the
	// user_code  database  table
	_, eErr := dbConn.Exec(s, c.Mail, c.Code, c.Token)
	return eErr

}

// ------------------------------------------------------------------------ //

// Exists is method of the UserCode
// data structure to check  if  row
// exists  by   the   token   field
func (c UserCode) Exists(dbConn *sql.DB) (bool, error) {

	var (

		// e   defines    the   row
		// existence     in     the
		// user_code database table
		e = false

		// s  defines the  template
		// of the  SQL -  statement
		s = `
		SELECT EXISTS (
			SELECT
				id
			FROM
				user_code
			WHERE
				code  = ?
				AND
				token = ?
		)
		`
	)

	// execute the SQL statement to
	// check if row  already exists
	if qErr := dbConn.QueryRow(s, c.Code, c.Token).Scan(&e); qErr == nil {
		return e, nil
	} else {
		return e, qErr
	}

}

// ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++ //
