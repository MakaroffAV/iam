package service

import (
	"testing"

	"github.com/MakaroffAV/backend-bootcamp-assignment-2024/internal/repo/db"
	"github.com/MakaroffAV/backend-bootcamp-assignment-2024/internal/repo/db/dbmodel"
)

func TestToken_UserByToken(t *testing.T) {
	c, err := db.Conn()
	if err != nil {
		panic(err)
	}
	defer c.Close()

	var (
		tk = NewToken(dbmodel.NewToken(c))
	)

	if _, err := tk.UserByToken("test_token"); err != nil {
		t.Fatalf(
			"case: 1; got: %v; want: nil; \n", err,
		)
	}
}
