package service

import (
	"testing"

	"github.com/MakaroffAV/backend-bootcamp-assignment-2024/internal/repo/db"
	"github.com/MakaroffAV/backend-bootcamp-assignment-2024/internal/repo/db/dbmodel"
)

func TestRegister_Do(t *testing.T) {
	c, err := db.Conn()
	if err != nil {
		panic(err)
	}
	defer c.Close()

	var (
		r = NewRegister(dbmodel.NewUser(c))
	)

	if _, err := r.Do(
		"test_email",
		"test_password", "moderator"); err != nil {

		t.Fatalf(
			"got: %v; want: nil; \n", err,
		)
	}
}
