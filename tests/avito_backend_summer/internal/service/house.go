package service

import (
	"errors"

	"github.com/MakaroffAV/backend-bootcamp-assignment-2024/internal/domain/model"
	"github.com/MakaroffAV/backend-bootcamp-assignment-2024/internal/repo"
)

var errUserModerator = errors.New(
	"пользователь не модератор",
)

type House struct {
	repoFlat  repo.Flat
	repoToken repo.Token
	repoHouse repo.House
}

func NewHouse(
	repoFlat repo.Flat,
	repoToken repo.Token,
	repoHouse repo.House,
) House {
	return House{
		repoFlat:  repoFlat,
		repoToken: repoToken,
		repoHouse: repoHouse,
	}
}

func (h House) Get(u model.User, id int) ([]model.Flat, error) {
	var m []model.Flat

	switch u.Type {
	// если тип пользователя не найден
	// отдаем                   ошибку
	default:
		{
			return m, errors.New(
				"тип пользователя",
			)
		}
	// если  пользователь НЕ модератор
	// отдаем только approved квартиры
	case model.UserTypeClientName:
		{
			return h.repoFlat.House(id, false)
		}

	// если  пользователь  модератор, то
	// отдаем ему все возможные квартиры
	case model.UserTypeModeratorName:
		{
			return h.repoFlat.House(id, true)
		}
	}
}

func (h House) Create(u model.User, address string, year int, developer string) (model.House, error) {
	var m model.House

	// проверим является ли
	// пользователь модератором
	if u.Type != model.UserTypeModeratorName {
		return m, errUserModerator
	}

	return h.repoHouse.Create(year, address, developer)
}
