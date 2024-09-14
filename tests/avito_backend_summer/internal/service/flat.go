package service

import (
	"errors"
	"sync"
	"time"

	"github.com/MakaroffAV/backend-bootcamp-assignment-2024/internal/domain/model"
	"github.com/MakaroffAV/backend-bootcamp-assignment-2024/internal/repo"
)

type Flat struct {

	// защитим мютексом, чтобы два пользователя
	// одновременно не могли работать с квартирами
	mu sync.Mutex

	repoFlat  repo.Flat
	repoToken repo.Token
}

func NewFlat(
	repoFlat repo.Flat,
	repoToken repo.Token,
) *Flat {
	return &Flat{
		repoFlat:  repoFlat,
		repoToken: repoToken,
	}
}

func flatStatusCheck(status string) bool {
	var s = map[string]bool{
		model.FlatStatusCreated:      true,
		model.FlatStatusApproved:     true,
		model.FlatStatusDeclined:     true,
		model.FlatStatusOnModeration: true,
	}

	// если статус известный - вернет true,
	// если нет - вернет false, false по умолчанию у bool
	return s[status]
}

func (f *Flat) Create(
	u model.User, id int64, houseID, price, rooms int) (model.Flat, error) {

	f.mu.Lock()
	defer f.mu.Unlock()

	// в доке openapi передается что-то вроде этого:
	//	{
	// 		"house_id": 12345,
	//		"price": 10000,
	//		"rooms": 4
	//	}
	//
	// но тогда, если где взять  id  квартиры ????????
	// сгенерировать самому - не  вариант, буду думать
	// что в openapi ошибка и id просто забыли указать
	return f.repoFlat.Create(id, houseID, price, rooms, time.Now())
}

func (f *Flat) Update(
	u model.User, id int64, houseID int, status string) (model.Flat, error) {

	var m model.Flat

	f.mu.Lock()
	defer f.mu.Unlock()

	if !flatStatusCheck(status) {
		return m, errors.New(
			"попытка смены на неизвестный статус",
		)
	}

	// проверим является ли
	// пользователь модератором
	if u.Type != model.UserTypeModeratorName {
		return m, errUserModerator
	}

	// самопальная, но проверка, что одну и ту
	// же  квартиру  нельзя взять на модерацию
	//
	// но по идее repo.Flat защищен мьютексом
	// => просмотреть квартиры, а потом взять
	// на модерацию и изменить нельзя одновременно
	c, err := f.repoFlat.Get(id, houseID)
	if err != nil {
		return m, err
	}
	if c.Status == model.FlatStatusOnModeration && status == model.FlatStatusOnModeration {
		return m, errors.New(
			"квартира уже находится на модерации",
		)
	}

	// в доке openapi передается что-то вроде
	//	{
	//		"id": 123456,
	//		"status": "approved"
	//	}
	//
	// опять же, в доке github  написано, что
	// номер квартиры не уникальный, но id, тогда
	// по логике  вещей  нам  нужен еше  id  дома
	return f.repoFlat.Update(id, houseID, status)
}
