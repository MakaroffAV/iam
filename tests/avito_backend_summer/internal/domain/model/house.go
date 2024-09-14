package model

import "time"

type House struct {
	// номер дома,  согласно требованиям
	// это будет просто уникальное число
	ID int

	// Адрес дома
	Address string

	// Имя застройщика
	Developer string

	// Год постройки дома
	Year int

	// Дата и время создания дома в базе
	Created time.Time

	// По умолчанию текущая дата и время
	// По факту - дата и время добавления квартиры
	Updated time.Time
}
