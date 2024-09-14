package main

import (
	"database/sql"
	"log"
	"sync"

	"github.com/MakaroffAV/backend-bootcamp-assignment-2024/internal/repo/db"
	"github.com/MakaroffAV/backend-bootcamp-assignment-2024/internal/repo/db/dbmodel"
	"github.com/MakaroffAV/backend-bootcamp-assignment-2024/internal/server"
	"github.com/MakaroffAV/backend-bootcamp-assignment-2024/internal/service"
)

func runServiceServer(c *sql.DB) error {
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

	return server.NewServer(
		servFlat, servHouse,
		servLogin, servToken, servRegister).Do()

}

func runServiceSender(c *sql.DB) error {
	return nil
}

func main() {
	c, err := db.Conn()
	if err != nil {
		panic(err)
	}
	defer c.Close()

	var wg sync.WaitGroup

	wg.Add(2)
	go func() {
		defer wg.Done()

		// если произойдет ошибка - экстренно завершаем приложение
		if err := runServiceSender(c); err != nil {
			log.Fatalf("sender service stopped; error: %v", err)
		}
	}()
	go func() {
		defer wg.Done()

		// если произойдет ошибка - экстренно завершаем приложение
		if err := runServiceServer(c); err != nil {
			log.Fatalf("server service stopped; error: %v", err)
		}
	}()

	wg.Wait()
}
