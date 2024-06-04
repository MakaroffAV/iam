// ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++ //

package models

// ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++ //

import (
	"fmt"
	"happy-birthday/internal/dbc"
	"log"
)

// ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++ //

// init is function to initialize
// the       models       package
func init() {

	// create  instance  of  the
	// mysql database connection
	c, cErr := dbc.Connection()
	if cErr != nil {
		fmt.Println(cErr)
		return

	}
	defer c.Close()

	// create    users   db   table
	if _, t1Err := c.Exec(
		`
		CREATE TABLE IF NOT EXISTS user (
			id   INT(11)      NOT NULL PRIMARY KEY AUTO_INCREMENT,
			name VARCHAR(100) NOT NULL,
			mail VARCHAR(100) NOT NULL UNIQUE
		);
		`,
	); t1Err != nil {
		log.Fatalf(
			`
			creation    the    users
			database table failed %s
			`,
			t1Err.Error(),
		)
	}

	// create  user_code  db  table
	if _, t2Err := c.Exec(
		`
		CREATE TABLE IF NOT EXISTS user_code (
			id    INT(11)      NOT NULL PRIMARY KEY AUTO_INCREMENT,
			mail  VARCHAR(100) NOT NULL,
			code  INT(11)      NOT NULL,
			token VARCHAR(200) NOT NULL
		);
		`,
	); t2Err != nil {
		log.Fatalf(
			`
			creation  the  user_code
			database table failed %s`,
			t2Err.Error(),
		)
	}

	// create user_session db table
	if _, t3Err := c.Exec(
		`
		CREATE TABLE IF NOT EXISTS user_session (
			id      INT(11)      NOT NULL PRIMARY KEY AUTO_INCREMENT,
			token   VARCHAR(200) NOT NULL,
			status  INT          NOT NULL,
			user_id INT(11)      NOT NULL,

			FOREIGN KEY (user_id) REFERENCES user(id)
		);
		`,
	); t3Err != nil {
		log.Fatalf(
			`
			creation the user_session
			database table  failed %s
			`,
			t3Err.Error(),
		)
	}

	// create   birthday  db  table
	if _, t4Err := c.Exec(
		`
		CREATE TABLE IF NOT EXISTS birthday (
			id            INT(11)      NOT NULL PRIMARY KEY AUTO_INCREMENT,
			name          VARCHAR(200) NOT NULL,
			bday          DATE         NOT NULL,
			subscribed    INT(11)      NOT NULL,
			created_by    INT(11)      NOT NULL,
			notify_before INT(11)      NOT NULL,

			FOREIGN KEY (created_by) REFERENCES user(id)
		)
		`,
	); t4Err != nil {
		log.Fatalf(
			`
			creation   the  birthday
			database table failed %s
			`,
			t4Err.Error(),
		)
	}

}

// ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++ //
