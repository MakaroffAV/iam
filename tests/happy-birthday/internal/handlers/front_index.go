// ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++ //

package handlers

// ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++ //

import (
	"happy-birthday/internal/models"
	"html/template"
	"net/http"
)

// ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++ //

// FrontIndex is  function to
// generate HTML template for
// the    index    app   page
func FrontIndex(w http.ResponseWriter, r *http.Request) {

	var s = ""

	// check  the   method  of
	// the   client's  request
	if r.Method != http.MethodGet {

		// build    and   send
		// the  HTML  template
		showStatusPage(w, &statusPage{
			Title:  "Ошибка",
			Header: "Недопустимый метод",
			Message: template.HTML(
				"Вернитесь на <a href='/front/'>главную страницу</a>",
			),
		})
		return
	}

	// check if user session's
	// token exists  and alive
	if !checkSession(r) {

		// build    and   send
		// the  HTML  template
		showStatusPage(w, &statusPage{
			Title:  "Ошибка",
			Header: "Сессия закрыта",
			Message: template.HTML(
				"Вернитесь на страницу <a href='/front/signin'>авторизации</a>",
			),
		})
		return
	}

	for _, v := range r.Cookies() {
		if v.Name == "sessionToken" {
			s = v.Value
		}
	}

	b, bErr := models.Birthday{}.GetAll(dbConn, s)
	if bErr != nil {

		// build    and   send
		// the  HTML  template
		showStatusPage(w, &statusPage{
			Title:  "Ошибка",
			Header: "Внутренняя ошибка сервера",
			Message: template.HTML(
				"Вернитесь на <a href='/front/'>главную страницу</a>",
			),
		})
		return
	}

	// build   and   send  the
	// HTML           template
	template.Must(template.ParseFiles("./templates/index.html")).Execute(
		w,
		struct {
			Bdays []models.Birthday
		}{
			Bdays: b,
		},
	)

}

// ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++ //
