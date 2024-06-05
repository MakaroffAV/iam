// ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++ //

package handlers

// ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++ //

import (
	"happy-birthday/internal/models"
	"html/template"
	"log"
	"net/http"
	"strconv"
)

// ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++ //

func FrontBirthdayEdit(w http.ResponseWriter, r *http.Request) {

	var s = ""

	if r.Method != http.MethodGet {
		showStatusPage(w, &statusPage{
			Title:  "Ошибка",
			Header: "Недопустимый метод",
			Message: template.HTML(
				"Вернитесь на <a href='/front/'>главную страницу</a>",
			),
		})
		return
	}

	for _, v := range r.Cookies() {
		if v.Name == "sessionToken" {
			s = v.Value
		}
	}

	i, iErr := strconv.Atoi(r.URL.Query().Get("id"))
	if iErr != nil {
		log.Print(iErr)
		showStatusPage(w, &statusPage{
			Title:  "Ошибка",
			Header: "Внутренняя ошибка сервера",
			Message: template.HTML(
				"Вернитесь на <a href='/front/'>главную страницу</a>",
			),
		})
		return
	}

	b, bErr := models.Birthday{Id: i}.GetById(dbConn, s)
	if bErr != nil {
		log.Print(bErr)
		showStatusPage(w, &statusPage{
			Title:  "Ошибка",
			Header: "Внутренняя ошибка сервера",
			Message: template.HTML(
				"Вернитесь на <a href='/front/'>главную страницу</a>",
			),
		})
		return
	}

	// build  and  send  the
	// HTML         template
	template.Must(template.ParseFiles("./templates/birthday_edit.html")).Execute(w, b)

}

// ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++ //
