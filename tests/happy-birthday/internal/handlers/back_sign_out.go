// ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++ //

package handlers

// ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++ //

import (
	"happy-birthday/internal/models"
	"html/template"
	"net/http"
)

// ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++ //

func BackSignOut(w http.ResponseWriter, r *http.Request) {

	var s = ""

	if r.Method != http.MethodGet {

		showStatusPage(
			w,
			&statusPage{
				Title:  "Ошибка",
				Header: "Недопустимый метод",
				Message: template.HTML(
					"Вернитесь на <a href='/front/index'>главную страницу</a>",
				),
			},
		)
		return
	}

	for _, v := range r.Cookies() {
		if v.Name == "sessionToken" {
			s = v.Value
		}
	}
	if len(s) == 0 {
		showStatusPage(
			w,
			&statusPage{
				Title:  "Ошибка",
				Header: "Внутренняя ошибка",
				Message: template.HTML(
					"Вернитесь на <a href='/front/index'>главную страницу</a>",
				),
			},
		)
		return
	}

	cErr := models.UserSession{}.Close(dbConn, s)
	if cErr != nil {
		showStatusPage(
			w,
			&statusPage{
				Title:  "Ошибка",
				Header: "Внутренняя ошибка",
				Message: template.HTML(
					"Вернитесь на <a href='/front/index'>главную страницу</a>",
				),
			},
		)
		return
	}

	http.Redirect(w, r, "/front/signin", http.StatusFound)

}

// ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++ //
