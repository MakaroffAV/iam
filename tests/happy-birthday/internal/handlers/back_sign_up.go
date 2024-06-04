// ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++ //

package handlers

// ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++ //

import (
	"happy-birthday/internal/models"
	"html/template"
	"log"
	"net/http"
)

// ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++ //

// BackSignUp is function
// to  create  new   user
func BackSignUp(w http.ResponseWriter, r *http.Request) {

	var (

		// t defines the value of
		// the signUpToken cookie
		t = ""

		// n defines the value of
		// user's            name
		n = ""

		// c defines the value of
		// user's  one-time  code
		c = ""
	)

	// check   the    method   of
	// the    client's    request
	if r.Method != http.MethodPost {

		// build     and     send
		// the    HTML   template
		showStatusPage(
			w,
			&statusPage{
				Title:  "Ошибка",
				Header: "Недопустимый метод",
				Message: template.HTML(
					"Вернитесь на страницу <a href='/front/signup'>регистрации</a>",
				),
			},
		)
		return
	}

	// get values  from the  form
	n = r.FormValue("name")
	c = r.FormValue("code")
	for _, v := range r.Cookies() {
		if v.Name == "signUpToken" {
			t = v.Value
		}
	}

	// check if  required  values
	// are        not       empty
	if len(n) == 0 || len(c) == 0 || len(t) == 0 {

		// log      the     error
		log.Print(
			`
			one  or  many sign up
			credentials are empty
			`,
		)

		// build     and     send
		// the    HTML   template
		showStatusPage(
			w,
			&statusPage{
				Title:  "Ошибка",
				Header: "Недопустимые параметры",
				Message: template.HTML(
					"Вернитесь на страницу <a href='/front/signup'>регистрации</a>",
				),
			},
		)
		return
	}

	// create  new user by passed
	// name,  code,   and   token
	i, iErr := models.User{Name: n}.CreateBySignUpToken(dbConn, c, t)
	if iErr != nil || i < 0 {

		// log      the     error
		log.Print(
			`
			creation new user's session
			after user  sign  up failed
			`,
		)

		// build     and     send
		// the    HTML   template
		showStatusPage(
			w,
			&statusPage{
				Title:  "Ошибка",
				Header: "Внутренняя ошибка сервера",
				Message: template.HTML(
					"Вернитесь на страницу <a href='/front/signup'>регистрации</a>",
				),
			},
		)
		return
	}

	createSession(w, r, i)

}

// ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++ //
