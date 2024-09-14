package service

import (
	"errors"

	"github.com/MakaroffAV/backend-bootcamp-assignment-2024/internal/domain/model"
	"github.com/MakaroffAV/backend-bootcamp-assignment-2024/internal/repo"
	"github.com/google/uuid"
)

type Login struct {
	repoUser  repo.User
	repoToken repo.Token
}

func NewLogin(
	repoUser repo.User,
	repoToken repo.Token) Login {

	return Login{
		repoUser:  repoUser,
		repoToken: repoToken,
	}
}

func (s Login) Login(userID int) (model.Token, error) {
	// uuid.New().String()
	//
	// внутри есть mutex,
	// можно безопасно использовать
	return s.repoToken.Create(userID, uuid.New().String())
}

func (s Login) DummyLogin(userType string) (model.Token, error) {
	var t model.Token

	//  проверка на допустимые
	// значения типов пользователей
	if !(userType == model.UserTypeClientName || userType == model.UserTypeModeratorName) {
		return t, errors.New(
			"неизвестный тип пользователя",
		)
	}

	u, err := s.repoUser.Create(
		userType, "", "", "", true,
	)
	if err != nil {
		return t, err
	}

	// uuid.New().String()
	//
	// внутри есть mutex,
	// можно безопасно использовать
	return s.repoToken.Create(u.ID, uuid.New().String())
}

func (s Login) UserExists(userUuid, password string) (model.User, error) {
	// вернет ошибку если
	// нет   пользователя
	return s.repoUser.Credentials(
		userUuid, password,
	)
}
