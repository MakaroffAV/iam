// ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++ //

package handlers

// ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++ //

import (
	"html/template"
	"net/http"
)

// ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++ //

// FrontSignUp is function to
// generate HTML template for
// the   user's  registration
func FrontSignUp(w http.ResponseWriter, r *http.Request) {

	// check  the   method  of
	// the   client's  request
	if r.Method != http.MethodGet {

		// build    and   send
		// the  HTML  template
		showStatusPage(w, &statusPage{
			Title:  "Ошибка",
			Header: "Недопустимый метод",
			Message: template.HTML(
				"Вернитесь на страницу <a href='/front/signup'>регистрации</a>",
			),
		})
		return
	}

	// build   and   send  the
	// HTML           template
	template.Must(template.ParseFiles("./templates/signup.html")).Execute(w, nil)

}

// ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++ //
