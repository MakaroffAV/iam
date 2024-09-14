package api

import (
	"encoding/json"
	"errors"
	"log"
	"net/http"
	"zadanie-6105/internal/service"
)

type Server struct {
	serviceBids    service.Bids
	serviceTenders service.Tenders
}

func NewServer(serviceTenders service.Tenders, serviceBids service.Bids) Server {
	return Server{
		serviceBids:    serviceBids,
		serviceTenders: serviceTenders,
	}
}

func (_ Server) write(w http.ResponseWriter, b []byte, c int) error {
	w.WriteHeader(c)
	w.Header().Add(
		"Content-Type", "application/json",
	)
	if _, err := w.Write(b); err != nil {
		return err
	} else {
		return nil
	}
}

func (s Server) response(w http.ResponseWriter, b any, err error) error {
	var c int
	switch {
	default:
		{
			c = http.StatusInternalServerError
		}
	case errors.Is(err, nil):
		{
			c = http.StatusOK
		}
	case errors.Is(err, service.ErrTenderNotFound):
		{
			c = http.StatusNotFound
		}
	case errors.Is(err, service.ErrTenderUsername):
		{
			c = http.StatusUnauthorized
		}
	case errors.Is(err, service.ErrTenderUsernameIllegal):
		{
			c = http.StatusForbidden
		}
	case errors.Is(err, service.ErrBidUsernameNotFound):
		{
			c = http.StatusUnauthorized
		}
	case errors.Is(err, service.ErrBidTenderNotFound):
		{
			c = http.StatusNotFound
		}
	case errors.Is(err, service.ErrBidNotFound):
		{
			c = http.StatusNotFound
		}
	}

	if c == http.StatusOK {
		m, mErr := json.Marshal(b)
		if mErr == nil {
			return s.write(w, m, c)
		} else {
			return s.write(w, nil, http.StatusInternalServerError)
		}
	} else {
		m, mErr := json.Marshal(
			struct {
				Reason string
			}{
				Reason: err.Error(),
			},
		)
		if mErr == nil {
			return s.write(w, m, c)
		} else {
			return s.write(w, nil, http.StatusInternalServerError)
		}
	}
}

// Получение списка ваших предложений
// (GET /bids/my)
func (s Server) GetUserBids(w http.ResponseWriter, r *http.Request, params GetUserBidsParams) {
	// изначально задаем
	// стандартные значения для всех параметров
	var (
		l int32  = 5
		o int32  = 0
		u string = ""
	)

	// проверяем, если есть значения от
	// пользователя, то обновляем значения параметров
	if params.Limit != nil {
		l = *params.Limit
	}
	if params.Offset != nil {
		o = *params.Offset
	}
	if params.Username != nil {
		u = *params.Username
	}

	// бизнес-логика
	v, err := s.serviceBids.GetByUsername(l, o, u)
	if rErr := s.response(w, v, err); rErr != nil {
		log.Println(
			err, "ошибка при отправке ответа клиенту",
		)
	}
}

// Создание нового предложения
// (POST /bids/new)
func (s Server) CreateBid(w http.ResponseWriter, r *http.Request) {
	var (
		b CreateBidJSONRequestBody
		d = json.NewDecoder(r.Body)
	)

	// Читаем тело запроса
	if err := d.Decode(&b); err != nil {
		if rErr := s.response(w, nil, err); rErr != nil {
			log.Println(
				err, "ошибка при отправке ответа клиенту",
			)
		}
	}

	// Бизнес-логика
	v, err := s.serviceBids.New(b.Name, b.Description, b.TenderId, string(b.AuthorType), b.AuthorId)
	if rErr := s.response(w, v, err); rErr != nil {
		log.Println(
			err, "ошибка при отправке ответа клиенту",
		)
	}
}

// Редактирование параметров предложения
// (PATCH /bids/{bidId}/edit)
func (s Server) EditBid(w http.ResponseWriter, r *http.Request, bidId BidId, params EditBidParams) {
	var (
		b EditBidJSONRequestBody
		d = json.NewDecoder(r.Body)

		rn string
		rd string
	)

	// Вот тут пробуем почитать данные от
	// клиента, если не вышло - отправляем ему ошибку
	if err := d.Decode(&b); err != nil {
		if rErr := s.response(w, nil, err); rErr != nil {
			log.Println(
				err, "ошибка при отправке ответа клиенту",
			)
		}
	}

	// Дальше смотрим на переданные параметы,
	// для не пустых - получаем их значения в переменные
	if b.Name != nil {
		rn = *b.Name
	}
	if b.Description != nil {
		rd = *b.Description
	}

	// Отправляем все это в бизнес-логику
	// Процедура такая же, что и при редактировании тендеров
	// 1. проверяем
	// 2. скидываем предложение в архивную таблицу
	// 3. обновляем предложение
	// 4. отдаем клиенту обновленное предложение
	v, err := s.serviceBids.Edit(
		bidId, params.Username, rn, rd,
	)
	if rErr := s.response(w, v, err); rErr != nil {
		log.Println(
			err, "ошибка при отправке ответа клиенту",
		)
	}
}

// Отправка отзыва по предложению
// (PUT /bids/{bidId}/feedback)
func (_ Server) SubmitBidFeedback(w http.ResponseWriter, r *http.Request, bidId BidId, params SubmitBidFeedbackParams) {
	w.WriteHeader(http.StatusNotImplemented)
}

// Откат версии предложения
// (PUT /bids/{bidId}/rollback/{version})
func (s Server) RollbackBid(w http.ResponseWriter, r *http.Request, bidId BidId, version int32, params RollbackBidParams) {
	v, err := s.serviceBids.RollBack(bidId, version, params.Username)
	if rErr := s.response(w, v, err); rErr != nil {
		log.Println(
			err, "ошибка при отправке ответа клиенту",
		)
	}
}

// Получение текущего статуса предложения
// (GET /bids/{bidId}/status)
func (s Server) GetBidStatus(w http.ResponseWriter, r *http.Request, bidId BidId, params GetBidStatusParams) {
	v, err := s.serviceBids.GetStatus(bidId, params.Username)
	if rErr := s.response(w, v, err); rErr != nil {
		log.Println(
			err, "ошибка при отправке ответа клиенту",
		)
	}
}

// Изменение статуса предложения
// (PUT /bids/{bidId}/status)
func (s Server) UpdateBidStatus(w http.ResponseWriter, r *http.Request, bidId BidId, params UpdateBidStatusParams) {
	v, err := s.serviceBids.UpdateStatus(bidId, params.Username, string(params.Status))
	if rErr := s.response(w, v, err); rErr != nil {
		log.Println(
			err, "ошибка при отправке ответа клиенту",
		)
	}
}

// Отправка решения по предложению
// (PUT /bids/{bidId}/submit_decision)
func (_ Server) SubmitBidDecision(w http.ResponseWriter, r *http.Request, bidId BidId, params SubmitBidDecisionParams) {
	w.WriteHeader(http.StatusNotImplemented)
}

// Получение списка предложений для тендера
// (GET /bids/{tenderId}/list)
func (s Server) GetBidsForTender(w http.ResponseWriter, r *http.Request, tenderId TenderId, params GetBidsForTenderParams) {
	// изначально задаем
	// стандартные значения для всех параметров
	var (
		l int32 = 5
		o int32 = 0
	)

	// заменяем стандартные
	// значения значениями от пользователя
	if params.Limit != nil {
		l = *params.Limit
	}
	if params.Offset != nil {
		o = *params.Offset
	}

	// бизнес-логика
	v, err := s.serviceBids.List(l, o, tenderId, params.Username)
	if rErr := s.response(w, v, err); rErr != nil {
		log.Println(
			err, "ошибка при отправке ответа клиенту",
		)
	}
}

// Просмотр отзывов на прошлые предложения
// (GET /bids/{tenderId}/reviews)
func (_ Server) GetBidReviews(w http.ResponseWriter, r *http.Request, tenderId TenderId, params GetBidReviewsParams) {
	w.WriteHeader(http.StatusNotImplemented)
}

// Проверка доступности сервера
// (GET /ping)
func (_ Server) CheckServer(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	if _, err := w.Write([]byte("ok\n")); err != nil {
		log.Fatalln(
			err, "ошибка при отправке ответа клиенту",
		)
	}
}

// Получение списка тендеров
// (GET /tenders)
func (server Server) GetTenders(w http.ResponseWriter, r *http.Request, params GetTendersParams) {
	// изначально задаем
	// стандартные значения для всех параметров
	var (
		l int32           = 5
		o int32           = 0
		s map[string]bool = map[string]bool{
			string(Delivery):     false,
			string(Manufacture):  false,
			string(Construction): false,
		}
	)

	// проверяем, если есть значения от
	// пользователя, то обновляем значения параметров
	if params.Limit != nil {
		l = *params.Limit
	}
	if params.Offset != nil {
		o = *params.Offset
	}
	if params.ServiceType == nil {
		// если не выбраны никакие
		// определенные сервисы - отдаем все
		for k := range s {
			s[k] = true
		}
	} else {
		// если переданы типы сервисов,
		// проверяем, все ли они допустимые, и для
		// допустимых - отдаем список тендоров с выбранными типами работ
		for _, v := range *params.ServiceType {
			if _, e := s[string(v)]; e {
				s[string(v)] = true
			}
		}
	}

	// бизнес-логика
	v, err := server.serviceTenders.Get(l, o, s)
	if rErr := server.response(w, v, err); rErr != nil {
		log.Println(
			err, "ошибка при отправке ответа клиенту",
		)
	}
}

// Получить тендеры пользователя
// (GET /tenders/my)
func (s Server) GetUserTenders(w http.ResponseWriter, r *http.Request, params GetUserTendersParams) {
	// изначально задаем
	// стандартные значения для всех параметров
	var (
		l int32  = 5
		o int32  = 0
		u string = ""
	)

	// проверяем, если есть значения от
	// пользователя, то обновляем значения параметров
	if params.Limit != nil {
		l = *params.Limit
	}
	if params.Offset != nil {
		o = *params.Offset
	}
	if params.Username != nil {
		u = *params.Username
	}

	// бизнес-логика
	v, err := s.serviceTenders.GetByUser(l, o, u)
	if rErr := s.response(w, v, err); rErr != nil {
		log.Println(
			err, "ошибка при отправке ответа клиенту",
		)
	}
}

// Создание нового тендера
// (POST /tenders/new)
func (s Server) CreateTender(w http.ResponseWriter, r *http.Request) {
	var (
		b CreateTenderJSONBody
		d = json.NewDecoder(r.Body)
	)

	// читаем данные из
	// запроса пользователя
	if err := d.Decode(&b); err != nil {
		if rErr := s.response(w, nil, err); rErr != nil {
			log.Println(
				err, "ошибка при отправке ответа клиенту",
			)
		}
	}

	// бизнес логика
	v, err := s.serviceTenders.New(b.Name, b.Description, string(b.ServiceType), b.OrganizationId, b.CreatorUsername)
	if rErr := s.response(w, v, err); rErr != nil {
		log.Println(
			err, "ошибка при отправке ответа клиенту",
		)
	}
}

// Редактирование тендера
// (PATCH /tenders/{tenderId}/edit)
func (s Server) EditTender(w http.ResponseWriter, r *http.Request, tenderId TenderId, params EditTenderParams) {
	var (
		b = EditTenderJSONBody{}
		d = json.NewDecoder(r.Body)

		en, ed, ek string
	)

	if err := d.Decode(&b); err != nil {
		if rErr := s.response(w, nil, err); rErr != nil {
			log.Println(
				err, "ошибка при отправке ответа клиенту",
			)
		}
	}

	if b.Name != nil {
		en = *b.Name
	}
	if b.Description != nil {
		ed = *b.Description
	}
	if b.ServiceType != nil {
		ek = string(*b.ServiceType)
	}

	v, err := s.serviceTenders.Update(
		tenderId, params.Username, en, ed, ek,
	)
	if rErr := s.response(w, v, err); rErr != nil {
		log.Println(
			err, "ошибка при отправке ответа клиенту",
		)
	}
}

// Откат версии тендера
// (PUT /tenders/{tenderId}/rollback/{version})
func (s Server) RollbackTender(w http.ResponseWriter, r *http.Request, tenderId TenderId, version int32, params RollbackTenderParams) {
	v, err := s.serviceTenders.Rollback(tenderId, params.Username, version)
	if rErr := s.response(w, v, err); rErr != nil {
		log.Println(
			err, "ошибка при отправке ответа клиенту",
		)
	}
}

// Получение текущего статуса тендера
// (GET /tenders/{tenderId}/status)
func (s Server) GetTenderStatus(w http.ResponseWriter, r *http.Request, tenderId TenderId, params GetTenderStatusParams) {
	var u string
	if params.Username != nil {
		u = *params.Username
	}
	v, err := s.serviceTenders.Status(tenderId, u)
	if rErr := s.response(w, v, err); rErr != nil {
		log.Println(
			err, "ошибка при отправке ответа клиенту",
		)
	}
}

// Изменение статуса тендера
// (PUT /tenders/{tenderId}/status)
func (s Server) UpdateTenderStatus(w http.ResponseWriter, r *http.Request, tenderId TenderId, params UpdateTenderStatusParams) {
	v, err := s.serviceTenders.UpdateStatus(
		tenderId, params.Username, string(params.Status),
	)
	if rErr := s.response(w, v, err); rErr != nil {
		log.Println(
			err, "ошибка при отправке ответа клиенту",
		)
	}
}
