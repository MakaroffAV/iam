package repository

import "zadanie-6105/internal/domain"

type Bid interface {
	Rollback(i string, v int32) error
	Edit(id, bn, bd string) error
	UpdateStatus(id string, s string) error
	GetByIdAndEmployeeID(l, o int32, tID, uID string) ([]domain.Bid, error)
	GetByUserID(l, o int32, u string) ([]domain.Bid, error)
	GetByID(i string) (domain.Bid, error)
	Create(i, n, d, t, at, ai string) error
}
type OrgResponsible interface {
	Exists(orgID, userID string) (bool, error)
}

type Employee interface {
	GetById(i string) (*domain.Employee, error)
	GetByUsername(n string) (*domain.Employee, error)
}

type Tender interface {
	Rollback(i string, v int32) error
	Update(i, n, d, k string) error
	UpdateStatus(i, s string) error
	GetFullByID(i string) (domain.Tender, error)
	Create(i, n, d, k, e string) error
	GetByID(i string) (domain.Tender, error)
	Get(l int32, o int32, s []string) ([]domain.Tender, error)
	GetByEmployeeID(l int32, o int32, u string) ([]domain.Tender, error)
}

type TenderArchive interface {
	GetByIdAndVersion(i string, v int32) (domain.Tender, error)
}

type BidArchive interface {
	GetByIdAndVersion(i string, v int32) (domain.Bid, error)
}
