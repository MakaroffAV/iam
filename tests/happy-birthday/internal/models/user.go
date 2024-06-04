// ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++ //

package models

// ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++ //

import "database/sql"

// ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++ //

// User is data structure
// to describe  the  user
// database         table
type User struct {
	Id   int
	Name string
	Mail string
}

// ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++ //

// Exists  is  method  of  the  User
// data structure  to  check if user
// has been already  exists into the
// users table by them email address
func (u User) Exists(dbConn *sql.DB) (bool, error) {

	var (

		// e defines  the existence of
		// the  user  into  the  table
		e = false

		// s defines the SQL statement
		// to check if user  has  been
		// already  exists  into   the
		// database              table
		s = `
		SELECT EXISTS (
			SELECT
				id
			FROM
				user
			WHERE
				mail = ?
		)
		`
	)

	// execute SQL  statement to check
	// user                  existence
	qErr := dbConn.QueryRow(s, u.Mail).Scan(&e)
	return e, qErr

}

// ------------------------------------------------------------------------ //

// UserIdBySignInToken is method
// of the User data structure to
// retrieve user id  by one-time
// code and sign in token  value
func (u User) UserIdBySignInToken(dbConn *sql.DB, code string, token string) (int, error) {

	var (

		// e defines  the  flag if user
		// already exists in  the table
		e = false

		// i defines the  new user's id
		i = -1

		// s1 defines the SQL statement
		// to check  if  pair of   user
		// code  and  token  is   valid
		s1 = `
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

		// s2 defines the SQL statement
		// to  retrieve  new  user's id
		s2 = `
		SELECT
			id
		FROM
			user
		WHERE
			user.mail = (
				SELECT
					mail
				FROM
					user_code
				WHERE
					code  = ?
					AND 
					token = ?
			)
		`
	)

	// check if pair of passed code and
	// token          is          valid
	qErr := dbConn.QueryRow(s1, code, token).Scan(&e)
	if qErr != nil {
		return i, qErr
	}

	// if pair of passed code and token
	// is not valid, return negative id
	if !e {
		return i, nil
	}

	// retrieve id of the  created user
	sErr := dbConn.QueryRow(s2, code, token).Scan(&i)
	if sErr != nil {
		return i, sErr
	}

	return i, nil

}

// ------------------------------------------------------------------------ //

// CreateBySignUpToken  is  method
// of the User data  structure  to
// create   new   user  by  passed
// sign up token and one-time code
func (u User) CreateBySignUpToken(dbConn *sql.DB, code string, token string) (int, error) {

	var (

		// e defines  the  flag if user
		// already exists in  the table
		e = false

		// i defines the  new user's id
		i = -1

		// s1 defines the SQL statement
		// to check  if  pair  of  user
		// code  and  token   is  valid
		s1 = `
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

		// s2 defines the SQL statement
		// to create new user by passed
		// user                metadata
		s2 = `
		INSERT INTO user (
			mail,
			name
		) VALUES (
			(
				SELECT
					mail
				FROM
					user_code
				WHERE
					code  = ?
					AND
					token = ?
			),
			?
		)
		`

		// s3 defines the SQL statement
		// to  retrieve  new  user's id
		s3 = `
		SELECT
			id
		FROM
			user
		WHERE
			user.mail = (
				SELECT
					mail
				FROM
					user_code
				WHERE
					code  = ?
					AND 
					token = ?
			)
		`
	)

	// check if pair of passed code and
	// token          is          valid
	qErr := dbConn.QueryRow(s1, code, token).Scan(&e)
	if qErr != nil {
		return i, qErr
	}

	// if pair of passed code and token
	// is not valid, return negative id
	if !e {
		return i, nil
	}

	// create new user by passed values
	_, eErr := dbConn.Exec(s2, code, token, u.Name)
	if eErr != nil {
		return i, eErr
	}

	// retrieve id of the  created user
	sErr := dbConn.QueryRow(s3, code, token).Scan(&i)
	if sErr != nil {
		return i, sErr
	}

	return i, nil

}

// ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++ //
