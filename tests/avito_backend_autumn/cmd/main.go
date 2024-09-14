package main

import (
	"github.com/go-chi/chi/v5"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"log"
	"net/http"
	"zadanie-6105/internal/api"
	"zadanie-6105/internal/config"
	"zadanie-6105/internal/repository/database"
	"zadanie-6105/internal/service"
)

// !!!
// Before commit to the main
// branch check if sql migrations applied
// !!!

// !!!
// ПРОБЛЕМЫ ТЗ (если проверка автоматизированая - покажите тесты, это же не сложно, но сможет устранить все неточности ТЗ)
//
// 1 Не прозрачно с версией тендера, обнолвение статуса - обновление тендера, увеличиваем при этом версию или нет??? Наверное, нет
// 2 Не написано, кто может обновлять статус тендера (буду обновлять только если статус редактируется тем же, кто создал тендер)
// 3 Что иммено должно происходить при откате версии тендера, должны ли восстанавливаться прошлые данные или просто изменяется версия?
//    Для решения этого вопроса заведем новую таблицу, tender_archive и туда булем кидать данные перед внесением изменений в тендер (чтобы при откате восстанавливать тендер полностью)
// 4 403 код при /api/bids/new. Не описано почему создание предложения должно быть заблокировано
//   Точнее описано, при том, что пользователь не имеет прав на это, нет ответа на вопрос почему он не может его редактировать
// !!!

// !!!
// НЕ РЕАЛИЗОВАНО ИЗ-ЗА доки, я ее не понял, процесс согласования
// (PUT /bids/{bidId}/feedback)
// (PUT /bids/{bidId}/submit_decision)
// (GET /bids/{tenderId}/reviews)
// !!!

func main() {
	d := config.NewDB().MustConn()
	defer d.Close()

	r := chi.NewRouter()
	s := api.NewServer(
		service.NewTenders(
			database.NewTender(d),
			database.NewEmployee(d),
			database.NewTenderArchive(d),
			database.NewOrgResponsible(d),
		),
		service.NewBids(
			database.NewEmployee(d),
			database.NewTender(d),
			database.NewBid(d),
			database.NewBidArchive(d),
		),
	)

	h := api.HandlerFromMuxWithBaseURL(s, r, "/api")

	if err := http.ListenAndServe(":8080", h); err != nil {
		log.Fatalln(
			err, "сервер неожиданно завершил работу",
		)
	}
}
