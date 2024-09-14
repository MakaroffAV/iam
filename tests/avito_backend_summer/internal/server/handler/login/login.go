package login

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/MakaroffAV/backend-bootcamp-assignment-2024/internal/service"
	"github.com/google/uuid"
)

type Login struct {
	servLogin service.Login
}

func NewLogin(servLogin service.Login) Login {
	return Login{
		servLogin: servLogin,
	}
}

func (s Login) invalidResponse(
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

func (s Login) validResponse(w http.ResponseWriter, token string) {
	var t = struct {
		Token string `json:"token"`
	}{
		Token: token,
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

func (s Login) Login(w http.ResponseWriter, r *http.Request) {
	i := uuid.New().String()

	log.SetPrefix("server.handler.Login.Login ")

	d := json.NewDecoder(r.Body)
	t := struct {
		I string `json:"id"`
		P string `json:"password"`
	}{}

	if err := d.Decode(&t); err != nil {
		log.Printf(
			"error: %v; request_id: %s; data: %v;", err, i, r.Body,
		)
		s.invalidResponse(
			w, i, 400, "ошибка чтения данных",
		)
		return
	}

	u, err := s.servLogin.UserExists(t.I, t.P)
	if err != nil {
		log.Printf(
			"error: %v; data: %v; request_id: %s", err, r.Body, i,
		)
		s.invalidResponse(
			w, i, 404, "пользователь не найден",
		)
		return
	}

	if m, err := s.servLogin.Login(u.ID); err == nil {
		s.validResponse(
			w, m.Value,
		)
	} else {
		log.Printf(
			"error: %v; data: %v; request_id: %s", err, r.Body, i,
		)
		s.invalidResponse(
			w, i, 500, "ошибка авторизации пользователя",
		)
	}
}

func (s Login) DummyLogin(w http.ResponseWriter, r *http.Request) {
	i := uuid.New().String()

	log.SetPrefix("server.handler.Login.DummyLogin ")

	t, err := s.servLogin.DummyLogin(
		r.URL.Query().Get("user_type"),
	)
	if err == nil {
		s.validResponse(w, t.Value)
	} else {
		log.Printf(
			"error: %v; user_type: %v; request_id: %s", err, r.URL.Query().Get("user_type"), i,
		)
		s.invalidResponse(
			w, i,
			500, "что-то и правда пошло не так",
		)
	}
}
