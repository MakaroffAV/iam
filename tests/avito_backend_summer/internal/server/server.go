package server

import (
	"net/http"

	"github.com/MakaroffAV/backend-bootcamp-assignment-2024/internal/server/routes"
	"github.com/MakaroffAV/backend-bootcamp-assignment-2024/internal/service"
)

type Server struct {
	servFlat     *service.Flat
	servHouse    service.House
	servLogin    service.Login
	servToken    service.Token
	servRegister service.Register
}

func NewServer(
	servFlat *service.Flat,
	servHouse service.House,
	servLogin service.Login,
	servToken service.Token,
	servRegister service.Register,
) Server {

	return Server{
		servFlat:     servFlat,
		servHouse:    servHouse,
		servLogin:    servLogin,
		servToken:    servToken,
		servRegister: servRegister,
	}
}

func (s Server) Do() error {
	m := http.NewServeMux()
	r := routes.Get(
		s.servFlat,
		s.servHouse,
		s.servLogin,
		s.servToken,
		s.servRegister,
	)

	for _, i := range r {
		m.HandleFunc(i.Path, i.Handler)
	}
	return http.ListenAndServe(":20501", m)
}
