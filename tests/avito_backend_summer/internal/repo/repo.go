package repo

import (
	"time"

	"github.com/MakaroffAV/backend-bootcamp-assignment-2024/internal/domain/model"
)

type Flat interface {
	Get(id int64, houseID int) (model.Flat, error)
	House(houseID int, all bool) ([]model.Flat, error)
	Update(id int64, houseID int, status string) (model.Flat, error)
	Create(id int64, houseID int, price int, rooms int, n time.Time) (model.Flat, error)
}

type User interface {
	Get(uuid string) (model.User, error)
	Credentials(uuid, password string) (model.User, error)
	Create(userType, uuid, email, password string, dummy bool) (model.User, error)
}

type House interface {
	ID(id int) (model.House, error)
	Create(year int, address string, developer string) (model.House, error)
}

type Token interface {
	Get(token string) (model.Token, error)
	User(token string) (model.User, error)
	Create(userID int, token string) (model.Token, error)
}
