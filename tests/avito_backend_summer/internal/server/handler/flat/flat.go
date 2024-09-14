package flat

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/MakaroffAV/backend-bootcamp-assignment-2024/internal/domain/model"
	"github.com/MakaroffAV/backend-bootcamp-assignment-2024/internal/service"
	"github.com/google/uuid"
)

type Flat struct {
	servFlat  *service.Flat
	servToken service.Token
}

func NewFlat(servToken service.Token, servFlat *service.Flat) Flat {
	return Flat{
		servFlat:  servFlat,
		servToken: servToken,
	}
}

func (f Flat) Update(w http.ResponseWriter, r *http.Request) {
	i := uuid.New().String()
	log.SetPrefix("server.handler.Flat.Update ")

	// если токен недействительный - вернет ошибку,
	// если с токеном все нормально - возвращаем пользователя
	u, err := f.servToken.UserByToken(r.Header.Get("token"))
	if err != nil {
		log.Printf(
			"err: %v; request_id: %s; token: %s", err, i, r.Header.Get("token"),
		)
		f.invalidResponse(
			w, i, 401, "пользователь не существует",
		)
		return
	}

	// готовим данные
	//
	// опять же, номер квартиры по
	// доке не уникальный. уникальная
	// комбинация номер  + номер дома
	d := json.NewDecoder(r.Body)
	t := struct {
		I int64  `json:"id"`
		H int    `json:"house_id"`
		S string `json:"status"`
	}{}

	// читаем данные
	if err := d.Decode(&t); err != nil {
		log.Printf(
			"error: %v; request_id: %s; body: %v", err, i, r.Body,
		)
		f.invalidResponse(
			w, i, 400, "ошибка чтения данных",
		)
		return
	}

	// обновляем квартиру
	if n, err := f.servFlat.Update(u, t.I, t.H, t.S); err == nil {
		f.validResponse(
			w, n,
		)
	} else {
		log.Printf(
			"error: %v; request_id: %s; body: %v", err, i, t,
		)
		f.invalidResponse(
			w, i, 500, "ошибка при обновлении квартиры",
		)
	}
}
func (f Flat) Create(w http.ResponseWriter, r *http.Request) {
	i := uuid.New().String()
	log.SetPrefix("server.handler.Flat.Create ")

	// если токен недействительный - вернет ошибку,
	// если с токеном все нормально - возвращаем пользователя
	u, err := f.servToken.UserByToken(r.Header.Get("token"))
	if err != nil {
		f.invalidResponse(
			w, i, 401, "пользователь не существует",
		)
		return
	}

	// готовим данные
	//
	// номер квартиры по доке не
	// уникальный - надо его читать
	d := json.NewDecoder(r.Body)
	t := struct {
		I int64 `json:"id"`
		H int   `json:"house_id"`
		P int   `json:"price"`
		R int   `json:"rooms"`
	}{}

	// читаем данные
	if err := d.Decode(&t); err != nil {
		log.Printf(
			"error: %v; request_id: %s; body: %v", err, i, r.Body,
		)
		f.invalidResponse(
			w, i, 400, "ошибка чтения данных",
		)
		return
	}

	// создаем квартиру
	if n, err := f.servFlat.Create(u, t.I, t.H, t.P, t.R); err == nil {
		f.validResponse(
			w, n,
		)
	} else {
		log.Printf(
			"error: %v; request_id: %s; body: %v", err, i, t,
		)
		f.invalidResponse(
			w, i, 500, "ошибка при создании квартиры",
		)
	}

}

func (s Flat) invalidResponse(
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

func (s Flat) validResponse(w http.ResponseWriter, flat model.Flat) {
	m, err := json.Marshal(flat)
	if err != nil {
		panic(err)
	}

	w.WriteHeader(200)
	if _, err := w.Write(m); err != nil {
		panic(err)
	}
}
