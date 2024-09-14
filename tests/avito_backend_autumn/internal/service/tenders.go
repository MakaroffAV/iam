package service

import (
	"database/sql"
	"errors"
	"github.com/google/uuid"
	"zadanie-6105/internal/domain"
	repo "zadanie-6105/internal/repository"
)

var (
	ErrTenderUsername = errors.New(
		"Пользователь не существует или некорректен",
	)
	ErrTenderNotFound = errors.New(
		"Тендер не найден",
	)
	ErrTenderUsernameIllegal = errors.New(
		"Недостаточно прав для выполнения действия",
	)
)

type Tenders struct {
	repoTender         repo.Tender
	repoEmployee       repo.Employee
	repoTenderArchive  repo.TenderArchive
	repoOrgResponsible repo.OrgResponsible
}

func NewTenders(repoTender repo.Tender, repoEmployee repo.Employee, repoTenderArchive repo.TenderArchive, repoOrgResponsible repo.OrgResponsible) Tenders {
	return Tenders{
		repoTender:         repoTender,
		repoEmployee:       repoEmployee,
		repoTenderArchive:  repoTenderArchive,
		repoOrgResponsible: repoOrgResponsible,
	}
}

func (t Tenders) Status(i string, u string) (string, error) {
	// для начала вытащим тендер по id
	v, err := t.repoTender.GetFullByID(i)
	if err != nil {
		if !errors.Is(err, sql.ErrNoRows) {
			return "", err
		} else {
			// если тендера не существует
			return "", ErrTenderNotFound
		}
	}

	// проверим существует ли пользователь
	e, err := t.repoEmployee.GetByUsername(u)
	if err != nil {
		return "", err
	}

	// если пользователя не существует - кидаем ошибку
	if e == nil {
		return "", ErrTenderUsername
	}

	// теперь у нас есть тендер и пользователь,
	// надо проверить, что тендер принадлежить пользователю
	if e.ID != v.EmployeeID {
		return "", ErrTenderUsernameIllegal
	}

	return v.Status, nil
}

func (t Tenders) StatusUpdate(s, i, u string) (domain.Tender, error) {
	return domain.Tender{}, nil
}

func (t Tenders) GetByUser(l int32, o int32, n string) ([]domain.Tender, error) {
	// для начала
	// проверяем пользователя
	u, err := t.repoEmployee.GetByUsername(n)
	if err != nil {
		return nil, err
	}

	// если вернулся nil
	// скидываем про ошибку пользователя
	if u == nil {
		return nil, ErrTenderUsername
	}

	// если пользователь был найден, то
	// берем из базы данных все его тендеры
	return t.repoTender.GetByEmployeeID(l, o, u.ID)
}

func (t Tenders) Get(l int32, o int32, s map[string]bool) ([]domain.Tender, error) {
	var services []string
	for k, v := range s {
		if v {
			services = append(services, k)
		}
	}
	return t.repoTender.Get(l, o, services)
}

func (t Tenders) New(n string, d string, k string, o string, c string) (domain.Tender, error) {
	// для начала
	// проверяем пользователя
	u, err := t.repoEmployee.GetByUsername(c)
	if err != nil {
		return domain.Tender{}, err
	}

	// Если пользователя не существует,
	// скидываем ошибку ErrTenderUsername
	if u == nil {
		return domain.Tender{}, ErrTenderUsername
	}

	// проверяем может ли пользователь
	// действовать от лица организации
	r, err := t.repoOrgResponsible.Exists(o, u.ID)
	if err != nil {
		return domain.Tender{}, err
	}

	// Если пользователь не может действовать от
	// лица организации - скидываем ошибку ErrTenderUsernameIllegal
	if !r {
		return domain.Tender{}, ErrTenderUsernameIllegal
	}

	// сгенерирую uuid тут, чтобы
	// потом по нему достать клиенту созданный тендер
	i := uuid.NewString()

	// пробуем создать
	// новый тендер от лицп пользователя
	if err := t.repoTender.Create(i, n, d, k, u.ID); err != nil {
		return domain.Tender{}, err
	}

	// после того, как тендер
	// успешно создан - возвращаем информацию о нем пользователю
	return t.repoTender.GetByID(i)
}

func (t Tenders) UpdateStatus(i, n, s string) (domain.Tender, error) {
	// для начала
	// проверяем пользователя
	u, err := t.repoEmployee.GetByUsername(n)
	if err != nil {
		return domain.Tender{}, err
	}

	// Если пользователя не существует,
	// скидываем ошибку ErrTenderUsername
	if u == nil {
		return domain.Tender{}, ErrTenderUsername
	}

	// Берем из базы тендер
	v, err := t.repoTender.GetFullByID(i)
	if err != nil {
		if !errors.Is(err, sql.ErrNoRows) {
			return domain.Tender{}, err
		} else {
			// проверка, если ошибка связана с тем,
			// что тендера нет, кидаем клиенту об этом информацию
			return domain.Tender{}, ErrTenderNotFound
		}
	}

	// проверяем - может ли
	// редактировать тендер пользователь
	if v.EmployeeID != u.ID {
		return domain.Tender{}, ErrTenderUsernameIllegal
	}

	// отправляем в бд запрос на изменение статуса у тендера
	if err := t.repoTender.UpdateStatus(i, s); err != nil {
		return domain.Tender{}, err
	}

	// после того, как статус тендера успешно
	// изменен - возвращаем информацию о нем пользователю
	return t.repoTender.GetByID(i)
}

func (t Tenders) Update(i, un, n, d, k string) (domain.Tender, error) {
	// для начала
	// проверяем пользователя
	u, err := t.repoEmployee.GetByUsername(un)
	if err != nil {
		return domain.Tender{}, err
	}

	// Если пользователя не существует,
	// скидываем ошибку ErrTenderUsername
	if u == nil {
		return domain.Tender{}, ErrTenderUsername
	}

	// Берем из базы тендер
	v, err := t.repoTender.GetFullByID(i)
	if err != nil {
		if !errors.Is(err, sql.ErrNoRows) {
			return domain.Tender{}, err
		} else {
			// проверка, если ошибка связана с тем,
			// что тендера нет, кидаем клиенту об этом информацию
			return domain.Tender{}, ErrTenderNotFound
		}
	}

	// проверяем - может ли
	// редактировать тендер пользователь
	if v.EmployeeID != u.ID {
		return domain.Tender{}, ErrTenderUsernameIllegal
	}

	// Отправляем запрос на редактирование тендера в БД
	// Редактирование происходит в несколько действий
	// 1. Скидываем текущий тендер в архивную таблицу
	// 2. Обновляем параметры у тендера

	if err := t.repoTender.Update(i, n, d, k); err != nil {
		return domain.Tender{}, err
	}

	// Возвращаем пользователю тендер с обновленными данными
	return t.repoTender.GetByID(i)
}

func (t Tenders) Rollback(i, n string, ver int32) (domain.Tender, error) {
	// для начала
	// проверяем пользователя
	u, err := t.repoEmployee.GetByUsername(n)
	if err != nil {
		return domain.Tender{}, err
	}

	// Если пользователя не существует,
	// скидываем ошибку ErrTenderUsername
	if u == nil {
		return domain.Tender{}, ErrTenderUsername
	}

	// Проверяем есть версия и тендер в архиве
	v, err := t.repoTenderArchive.GetByIdAndVersion(i, ver)
	if err != nil {
		if !errors.Is(err, sql.ErrNoRows) {
			return domain.Tender{}, err
		} else {
			return domain.Tender{}, ErrTenderNotFound
		}
	}

	// проверяем - может ли
	// редактировать тендер пользователь
	if v.EmployeeID != u.ID {
		return domain.Tender{}, ErrTenderUsernameIllegal
	}

	// Откатываем версию тендера
	if err := t.repoTender.Rollback(i, ver); err != nil {
		return domain.Tender{}, err
	}

	// Возвращаем пользователю тендер с обновленными данными
	return t.repoTender.GetByID(i)
}
