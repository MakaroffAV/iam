package model

import "time"

const (
	UserTypeClientID   = 0
	UserTypeClientName = "client"

	UserTypeModeratorID   = 1
	UserTypeModeratorName = "moderator"
)

type User struct {
	// Serial DB
	ID int

	// client, moderator
	Type string

	// Уникальный id пользователя
	Uuid string

	Email    string
	Password string

	// у  нас  есть  ручка  dummyLogin,
	// при ее использовании  мы заведем
	// пустого пользователя в БД, чтобы
	// выдать  ему  токен,  посредством
	// которого будем  отслеживать  его
	Dummy bool

	// дата и время создания
	Created time.Time
}
