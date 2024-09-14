package domain

import "time"

type Bid struct {
	ID          string
	Name        string
	Description string
	Status      string
	Version     int
	TenderID    string
	AuthorType  string
	AuthorID    string
	CreatedAt   time.Time
}
