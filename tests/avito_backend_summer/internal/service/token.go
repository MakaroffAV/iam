package service

import (
	"github.com/MakaroffAV/backend-bootcamp-assignment-2024/internal/domain/model"
	"github.com/MakaroffAV/backend-bootcamp-assignment-2024/internal/repo"
)

type Token struct {
	repoToken repo.Token
}

func NewToken(repoToken repo.Token) Token {
	return Token{
		repoToken: repoToken,
	}
}

func (t Token) UserByToken(token string) (model.User, error) {
	return t.repoToken.User(token)
}
