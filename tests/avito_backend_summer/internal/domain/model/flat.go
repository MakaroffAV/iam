package model

const (
	FlatStatusCreated      = "created"
	FlatStatusApproved     = "approved"
	FlatStatusDeclined     = "declined"
	FlatStatusOnModeration = "on moderation"
)

type Flat struct {

	// ID - номер квартиры
	ID    int64
	Price int
	Rooms int

	// id дома, House.ID
	HouseID int

	// created, moderate, approved, declined
	// ! хорошим решением будет внешний ключ
	Status string
}
