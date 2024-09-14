package routes

import (
	"testing"

	"github.com/MakaroffAV/backend-bootcamp-assignment-2024/internal/repo/db"
	"github.com/MakaroffAV/backend-bootcamp-assignment-2024/internal/repo/db/dbmodel"
	"github.com/MakaroffAV/backend-bootcamp-assignment-2024/internal/service"
)

func TestGet(t *testing.T) {

	c, err := db.Conn()
	if err != nil {
		panic(err)
	}
	defer c.Close()

	var (
		n = 7
	)

	var (
		repoUser  = dbmodel.NewUser(c)
		repoFlat  = dbmodel.NewFlat(c)
		repoToken = dbmodel.NewToken(c)
		repoHouse = dbmodel.NewHouse(c)
	)

	var (
		servToken    = service.NewToken(repoToken)
		servRegister = service.NewRegister(repoUser)
		servFlat     = service.NewFlat(repoFlat, repoToken)
		servLogin    = service.NewLogin(repoUser, repoToken)
		servHouse    = service.NewHouse(repoFlat, repoToken, repoHouse)
	)

	r := Get(servFlat, servHouse, servLogin, servToken, servRegister)
	if len(r) != n {
		t.Fatalf(
			"case: %d; got: %d; want: %d; \n", 1, len(r), n,
		)
	}
}
