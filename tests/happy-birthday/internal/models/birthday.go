// ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++ //

package models

// ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++ //

import "database/sql"

// ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++ //

// Birthday is  data structure
// to describe 'Birthday' man
type Birthday struct {
	Id           int
	Name         string
	BDay         string
	Subscribed   int
	CreatedBy    int
	NotifyBefore int
}

// ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++ //

// Create  is  function   to
// create new 'Birthday' man
func (b Birthday) Create(dbConn *sql.DB, sessionToken string) error {

	// s defines the template
	// of the  SQL  statement
	var s = `
	INSERT INTO birthday (
		name,
		bday,
		subscribed,
		created_by,
		notify_before
	) VALUES (
		?,
		?,
		1,
		(
			SELECT
				user_id
			FROM
				user_session
			WHERE
				token = ?
		),
		?
	)
	`

	// execute SQL  statement to
	// create new 'Birthday' man
	_, eErr := dbConn.Exec(s, b.Name, b.BDay, sessionToken, b.NotifyBefore)
	return eErr

}

// ------------------------------------------------------------------------ //

// GetAll is method  of the  Birthday
// data  structure  to  retrieve  all
// 'birthday' men created by the user
func (b Birthday) GetAll(dbConn *sql.DB, sessionToken string) ([]Birthday, error) {

	var bDays = []Birthday{}

	// s defines the template
	// of the  SQL  statement
	var s = `
	SELECT
		id,
		name,
		bday,
		subscribed,
		notify_before
	FROM
		birthday
	WHERE
		birthday.created_by = (
			SELECT
				user_id
			FROM
				user_session
			WHERE
				token   =  ?
		)
	`

	r, rErr := dbConn.Query(s, sessionToken)
	if rErr != nil {
		return nil, rErr
	}

	for r.Next() {
		var bDay Birthday
		if sErr := r.Scan(&bDay.Id, &bDay.Name, &bDay.BDay, &bDay.Subscribed, &bDay.NotifyBefore); sErr != nil {
			return nil, sErr
		}
		bDays = append(bDays, bDay)
	}

	return bDays, nil

}

// ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++ //
