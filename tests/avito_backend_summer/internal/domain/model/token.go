package model

import "time"

type Token struct {
	// Serial DB
	ID int

	// token value
	Value string

	// housing.user.id
	UserID int

	// token created timestamp
	Created time.Time
}
