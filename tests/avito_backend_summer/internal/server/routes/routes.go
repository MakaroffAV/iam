package routes

import (
	"net/http"

	"github.com/MakaroffAV/backend-bootcamp-assignment-2024/internal/server/handler/flat"
	"github.com/MakaroffAV/backend-bootcamp-assignment-2024/internal/server/handler/house"
	"github.com/MakaroffAV/backend-bootcamp-assignment-2024/internal/server/handler/login"
	"github.com/MakaroffAV/backend-bootcamp-assignment-2024/internal/server/handler/register"
	"github.com/MakaroffAV/backend-bootcamp-assignment-2024/internal/service"
)

type route struct {
	Path    string
	Handler func(w http.ResponseWriter, r *http.Request)
}

func Get(
	servFlat *service.Flat,
	servHouse service.House,
	servLogin service.Login,
	servToken service.Token,
	servRegister service.Register) []route {

	return []route{
		{
			Path:    "/login",
			Handler: login.NewLogin(servLogin).Login,
		},
		{
			Path:    "/dummyLogin",
			Handler: login.NewLogin(servLogin).DummyLogin,
		},
		{
			Path:    "/register",
			Handler: register.NewRegister(servRegister).Do,
		},
		{
			Path:    "/house/{id}",
			Handler: house.NewHouse(servHouse, servToken).Flats,
		},
		{
			Path:    "/flat/create",
			Handler: flat.NewFlat(servToken, servFlat).Create,
		},
		{
			Path:    "/house/create",
			Handler: house.NewHouse(servHouse, servToken).Create,
		},
		{
			Path:    "/flat/update",
			Handler: flat.NewFlat(servToken, servFlat).Update,
		},
	}
}
