package domain

import "time"

type Employee struct {
	ID        string
	Username  string
	FirstName string
	LastName  string
	CreatedAt time.Time
	UpdatedAt time.Time
}
