package register

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/MakaroffAV/backend-bootcamp-assignment-2024/internal/service"
	"github.com/google/uuid"
)

type Register struct {
	servRegister service.Register
}

func NewRegister(servRegister service.Register) Register {
	return Register{
		servRegister: servRegister,
	}
}

func (s Register) Do(w http.ResponseWriter, r *http.Request) {
	i := uuid.New().String()
	log.SetPrefix("server.handler.register.Do ")

	d := json.NewDecoder(r.Body)
	t := struct {
		E string `json:"email"`
		P string `json:"password"`
		U string `json:"user_type"`
	}{}

	if err := d.Decode(&t); err != nil {
		log.Printf(
			"error: %v; data: %v; request_id: %s", err, r.Body, i,
		)
		s.invalidResponse(
			w, i, 400, "ошибка чтения данных",
		)
		return
	}

	if m, err := s.servRegister.Do(t.E, t.P, t.U); err == nil {
		s.validResponse(
			w, m.Uuid,
		)
	} else {
		log.Printf(
			"error: %v; data: %v; request_id: %s", err, r.Body, i,
		)
		s.invalidResponse(
			w, i, 500, "ошибка создания нового пользователя",
		)
	}
}

func (s Register) invalidResponse(
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

func (s Register) validResponse(w http.ResponseWriter, id string) {
	var t = struct {
		UserID string `json:"user_id"`
	}{
		UserID: id,
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
