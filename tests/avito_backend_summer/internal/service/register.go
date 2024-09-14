package service

import (
	"github.com/MakaroffAV/backend-bootcamp-assignment-2024/internal/domain/model"
	"github.com/MakaroffAV/backend-bootcamp-assignment-2024/internal/repo"
	"github.com/google/uuid"
)

type Register struct {
	repoUser repo.User
}

func NewRegister(repoUser repo.User) Register {
	return Register{
		repoUser: repoUser,
	}
}

func (r Register) Do(email, password, userType string) (model.User, error) {
	return r.repoUser.Create(
		userType, uuid.New().String(), email, password, false,
	)
}
