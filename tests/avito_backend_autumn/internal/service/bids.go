package service

import (
	"database/sql"
	"errors"
	"github.com/google/uuid"
	"zadanie-6105/internal/domain"
	repo "zadanie-6105/internal/repository"
)

var (
	ErrBidUsernameNotFound = errors.New(
		"Пользователь не существует или некорректен",
	)
	ErrBidTenderNotFound = errors.New(
		"Тендер не найден",
	)
	ErrBidNotFound = errors.New(
		"Предложение не найдено",
	)
)

type Bids struct {
	repoBid        repo.Bid
	repoEmployee   repo.Employee
	repoTender     repo.Tender
	repoBidArchive repo.BidArchive
}

func NewBids(repoEmployee repo.Employee, repoTender repo.Tender, repoBid repo.Bid, repoBidArchive repo.BidArchive) Bids {
	return Bids{
		repoBid:        repoBid,
		repoTender:     repoTender,
		repoEmployee:   repoEmployee,
		repoBidArchive: repoBidArchive,
	}
}

func (b Bids) New(n, d, ti, at, ai string) (domain.Bid, error) {
	// для начала
	// проверяем, что пользователь существует
	u, err := b.repoEmployee.GetById(ai)
	if err != nil {
		return domain.Bid{}, err
	}

	// Если пользователя не существует,
	// скидываем ошибку errBidUsernameNotFound
	if u == nil {
		return domain.Bid{}, ErrBidUsernameNotFound
	}

	// !!!
	// Вот тут должна быть обработка на
	// права пользователя, но ее механизм не описан
	// !!!

	// Вытаскиваем из базы тендер,
	// предложение по которому хотим разместить
	t, err := b.repoTender.GetFullByID(ti)
	if err != nil {
		if !errors.Is(err, sql.ErrNoRows) {
			return domain.Bid{}, err
		} else {
			return domain.Bid{}, ErrBidTenderNotFound
		}
	}

	// Записываем в базу новой предложение по тендеру
	ni := uuid.NewString()
	if err := b.repoBid.Create(ni, n, d, t.ID, at, u.ID); err != nil {
		return domain.Bid{}, err
	}

	// Вытаскиваем из базы созданное предложение по тендеру
	return b.repoBid.GetByID(ni)
}

func (b Bids) GetByUsername(l, o int32, n string) ([]domain.Bid, error) {
	// для начала
	// проверяем пользователя
	u, err := b.repoEmployee.GetByUsername(n)
	if err != nil {
		return nil, err
	}

	// если вернулся nil
	// скидываем про ошибку пользователя
	if u == nil {
		return nil, ErrBidUsernameNotFound
	}

	// Если пользователь был найден, то
	// вытаскиваем из базы все предложения от пользователя
	return b.repoBid.GetByUserID(l, o, u.ID)
}

func (b Bids) List(l, o int32, tID, un string) ([]domain.Bid, error) {
	// для начала
	// проверяем пользователя
	u, err := b.repoEmployee.GetByUsername(un)
	if err != nil {
		return nil, err
	}

	// Если вернулся nil,
	// скидываем про ошибку пользователя
	if u == nil {
		return nil, ErrBidUsernameNotFound
	}

	// Вытаскиваем тендер по его идентификатору
	t, err := b.repoTender.GetFullByID(tID)
	if err != nil {
		if !errors.Is(err, sql.ErrNoRows) {
			return nil, err
		} else {
			return nil, ErrTenderNotFound
		}
	}

	// !!!
	// вот тут должна быть проверка
	// на права пользователя, НО ее механизм не описан
	// НО
	// По идее мы вытаскиваем предложения полтзователя по тендеру
	// соответственно, надо проверить, чтобы id размещаюшего тендер был равен id запращиваюшего
	// !!!
	if t.EmployeeID != u.ID {
		return nil, ErrTenderUsernameIllegal
	}

	return b.repoBid.GetByIdAndEmployeeID(l, o, t.ID, u.ID)
}

func (b Bids) GetStatus(id, un string) (string, error) {
	// для начала
	// проверяем пользователя
	u, err := b.repoEmployee.GetByUsername(un)
	if err != nil {
		return "", err
	}

	// Если вернулся nil,
	// скидываем про ошибку пользователя
	if u == nil {
		return "", ErrBidUsernameNotFound
	}

	// Вытаскиваем предложение по его ID
	v, err := b.repoBid.GetByID(id)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return "", ErrBidNotFound
		}
		return "", err
	}

	// Проверяем, что пользователь, который
	// разместил предложение хочет посмотреть его статус
	if v.AuthorID != u.ID {
		return "", ErrTenderUsernameIllegal
	}

	// Если все проверки выполнены успешно, отдаем статус предложения
	return v.Status, nil
}

func (b Bids) UpdateStatus(id, un, st string) (domain.Bid, error) {
	// для начала
	// проверяем пользователя
	u, err := b.repoEmployee.GetByUsername(un)
	if err != nil {
		return domain.Bid{}, err
	}

	// Если вернулся nil,
	// скидываем про ошибку пользователя
	if u == nil {
		return domain.Bid{}, ErrBidUsernameNotFound
	}

	// Вытаскиваем предложение по его ID
	v, err := b.repoBid.GetByID(id)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return domain.Bid{}, ErrBidNotFound
		}
		return domain.Bid{}, err
	}

	// !!!
	// Опять же в ТЗ нет четких слов про
	// то, кто может обновлять статус предложения
	//
	// Будем обновлять статус предложения
	// если об этом просит только его создатель
	// !!!

	// Проверяем, что пользователь, который
	// разместил предложение хочет обновить его статус
	if v.AuthorID != u.ID {
		return domain.Bid{}, ErrTenderUsernameIllegal
	}

	// Если все проверки выполнены, то
	// идем в базу и обновляем статус предложения
	if err := b.repoBid.UpdateStatus(v.ID, st); err != nil {
		return domain.Bid{}, err
	}

	// Обновили статус предложения, идем в базу и отдаем предложение с новым статусом
	return b.repoBid.GetByID(v.ID)
}

func (b Bids) Edit(id, un, bn, bd string) (domain.Bid, error) {
	// для начала
	// проверяем пользователя
	u, err := b.repoEmployee.GetByUsername(un)
	if err != nil {
		return domain.Bid{}, err
	}

	// Если вернулся nil,
	// скидываем про ошибку пользователя
	if u == nil {
		return domain.Bid{}, ErrBidUsernameNotFound
	}

	// Вытаскиваем предложение по его ID
	v, err := b.repoBid.GetByID(id)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return domain.Bid{}, ErrBidNotFound
		}
		return domain.Bid{}, err
	}

	// Проверяем, что пользователь, который
	// разместил предложение хочет отредактировать его
	if v.AuthorID != u.ID {
		return domain.Bid{}, ErrTenderUsernameIllegal
	}

	// Идем в базу и в одной транзакции скидываем
	// предложение старое и редактируем, как просит клиент
	if err := b.repoBid.Edit(v.ID, bn, bd); err != nil {
		return domain.Bid{}, err
	}
	// Вытаскиваем из базы обновленное предложение
	return b.repoBid.GetByID(v.ID)
}

func (b Bids) RollBack(id string, ver int32, un string) (domain.Bid, error) {
	// для начала
	// проверяем пользователя
	u, err := b.repoEmployee.GetByUsername(un)
	if err != nil {
		return domain.Bid{}, err
	}

	// Если вернулся nil,
	// скидываем про ошибку пользователя
	if u == nil {
		return domain.Bid{}, ErrBidUsernameNotFound
	}

	// Проверяем в архиве есть ли
	// версия предложения к которой мы хотим откатиться
	v, err := b.repoBidArchive.GetByIdAndVersion(id, ver)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return domain.Bid{}, ErrBidNotFound
		}
		return domain.Bid{}, err
	}

	// Проверяем может ли
	// редактировать предложение пользователь
	if v.AuthorID != u.ID {
		return domain.Bid{}, ErrTenderUsernameIllegal
	}

	// Откатываем версию предложения
	if err := b.repoBid.Rollback(id, ver); err != nil {
		return domain.Bid{}, err
	}

	// Возвращаем пользователю предложение с обновленными данными
	return b.repoBid.GetByID(v.ID)

}
