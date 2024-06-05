// ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++ //

package models

import "database/sql"

// ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++ //

// UserSession is data structure
// to describe metadata  of  the
// user                  session
type UserSession struct {
	UserId int
	Status int
	SToken string
}

// ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++ //

// Create  is   method  of  the
// UserSession  data  structure
// to create new user's session
func (u UserSession) Create(dbConn *sql.DB) error {

	var (

		// s defines  the template
		// of the SQL statement to
		// create    new   session
		s = `
		INSERT INTO user_session (
			token,
			status,
			user_id
		) VALUES (
			?,
			?,
			?
		)
		`
	)

	// execute  SQL  statement  to
	// create new  user's  session
	_, qErr := dbConn.Exec(s, u.SToken, 1, u.UserId)
	return qErr

}

// ------------------------------------------------------------------------ //

// Check is method of the UserSession
// data structure  to  check  if user
// session   is   exists   and  alive
func (u UserSession) Check(dbConn *sql.DB) (bool, error) {

	var (

		// e defines the  flag  if
		// the user session exists
		// and               alive
		e = false

		// s defines  the template
		// of the SQL statement to
		// check   session   token
		s = `
		SELECT EXISTS (
			SELECT
				id
			FROM
				user_session
			WHERE
				token  = ?
				AND
				status = 1
		)
		`
	)

	// execute  SQL  statement  to
	// check     user's    session
	qErr := dbConn.QueryRow(s, u.SToken).Scan(&e)
	return e, qErr

}

// ------------------------------------------------------------------------ //

func (u UserSession) Close(dbConn *sql.DB, sessionToken string) error {

	var s = `
	UPDATE
		user_session
	SET
		user_session.status = 0
	WHERE
		user_session.token  = ?
	`

	_, eErr := dbConn.Exec(s, sessionToken)
	return eErr

}

// ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++ //
