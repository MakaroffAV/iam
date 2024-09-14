package house

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/MakaroffAV/backend-bootcamp-assignment-2024/internal/domain/model"
	"github.com/MakaroffAV/backend-bootcamp-assignment-2024/internal/service"
	"github.com/google/uuid"
)

type House struct {
	servToken service.Token
	servHouse service.House
}

func NewHouse(servHouse service.House, servToken service.Token) House {
	return House{
		servToken: servToken,
		servHouse: servHouse,
	}
}

func (h House) Flats(w http.ResponseWriter, r *http.Request) {
	i := uuid.New().String()
	log.SetPrefix("server.handler.House.Flats ")

	// номер дома
	n, err := strconv.Atoi(r.PathValue("id"))
	if err != nil {
		log.Printf(
			"error: %v; id: %s; house_id: %s", err, i, r.PathValue("id"),
		)
		h.invalidResponse(
			w, i, 400, "неверный номер дома",
		)
		return
	}

	// если токен недействительный - вернет ошибку,
	// если с токеном все нормально - возвращаем пользователя
	u, err := h.servToken.UserByToken(r.Header.Get("token"))
	if err != nil {
		log.Printf(
			"error: %v; request_id: %s; token: %s", err, i, r.Header.Get("token"),
		)
		h.invalidResponse(
			w, i, 401, "пользоваатель не существует",
		)
		return
	}

	// внутри само решит квартиры какого типа отдавать клиенту
	if f, err := h.servHouse.Get(u, n); err == nil {
		h.validResponse(
			w, f,
		)
	} else {
		log.Printf(
			"error: %v; request_id: %s; token: %s;", err, i, r.Header.Get("token"),
		)
		h.invalidResponse(
			w, i, 500, "ошибка при извлечении квартир",
		)
	}
}

func (h House) Create(w http.ResponseWriter, r *http.Request) {
	i := uuid.New().String()
	log.SetPrefix("server.handler.House.Create ")

	// если токен недействительный - вернет ошибку,
	// если с токеном все нормально - возвращаем пользователя
	u, err := h.servToken.UserByToken(r.Header.Get("token"))
	if err != nil {
		log.Printf(
			"error: %v; id: %s; token: %s", err, i, r.Header.Get("token"),
		)
		h.invalidResponse(
			w, i, 401, "пользоваатель не существует",
		)
		return
	}

	// готовим данные
	d := json.NewDecoder(r.Body)
	t := struct {
		A string `json:"address"`
		Y int    `json:"year"`
		D string `json:"developer"`
	}{}

	// читаем данные
	if err := d.Decode(&t); err != nil {
		log.Printf(
			"error: %v; id: %s; data: %v", err, i, r.Body,
		)
		h.invalidResponse(
			w, i, 400, "ошибка чтения данных",
		)
		return
	}

	// создаем квартиру, если пользователь
	// НЕ модератор - вернет ошибку
	if house, err := h.servHouse.Create(u, t.A, t.Y, t.D); err == nil {
		h.validResponseHouse(
			w, house,
		)
	} else {
		log.Printf(
			"error: %v; id: %s; data: %v; user: %v", err, i, r.Body, u,
		)
		h.invalidResponse(
			w, i, 500, "ошибка при создании дома",
		)
	}
}

func (s House) invalidResponse(
	w http.ResponseWriter, reqID string, code int, msg string) {

	var t = struct {
		C int    `json:"code"`
		M string `json:"message"`
		I string `json:"request_id"`
	}{
		C: code,
		M: msg,
		I: reqID,
	}

	m, err := json.Marshal(t)
	if err != nil {
		panic(err)
	}

	if code == 500 {
		w.Header().Add(
			"Retry-After", "10",
		)
	}

	w.WriteHeader(code)
	if _, err := w.Write(m); err != nil {
		panic(err)
	}
}

func (s House) validResponse(w http.ResponseWriter, flats []model.Flat) {
	var t = struct {
		Flats []model.Flat `json:"flats"`
	}{
		Flats: flats,
	}

	m, err := json.Marshal(t)
	if err != nil {
		panic(err)
	}

	w.WriteHeader(200)
	if _, err := w.Write(m); err != nil {
		panic(err)
	}
}

func (s House) validResponseHouse(w http.ResponseWriter, house model.House) {
	m, err := json.Marshal(house)
	if err != nil {
		panic(err)
	}

	w.WriteHeader(200)
	if _, err := w.Write(m); err != nil {
		panic(err)
	}
}
