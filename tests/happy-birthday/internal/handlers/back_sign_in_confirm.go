// ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++ //

package handlers

// ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++ //

import (
	"fmt"
	"happy-birthday/internal/models"
	"html/template"
	"log"
	"net/http"
)

// ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++ //

// BackSignInConfirm is function to
// confirm  user sign in  procedure
func BackSignInConfirm(w http.ResponseWriter, r *http.Request) {

	var (

		// t defines the value of
		// the signInToken cookie
		t = ""

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
					"Вернитесь на страницу <a href='/front/signin'>авторизации</a>",
				),
			},
		)
		return
	}

	// get values  from  the form
	c = r.FormValue("code")
	for _, v := range r.Cookies() {
		if v.Name == "signInToken" {
			t = v.Value
		}
	}
	if len(c) == 0 || len(t) == 0 {

		// log      the     error
		log.Print(
			`
			one or  many  sign in
			credentials are empty
			`,
		)

		// build    and      send
		// the    HTML   template
		showStatusPage(
			w,
			&statusPage{
				Title:  "Ошибка",
				Header: "Недопустимые параметры",
				Message: template.HTML(
					"Вернитесь на страницу <a href='/front/signin'>авторизации</a>",
				),
			},
		)
		return
	}

	// check  one-time  code  and
	// retrieve     user's     id
	i, iErr := models.User{}.UserIdBySignInToken(dbConn, c, t)
	if iErr != nil || i < 0 {

		fmt.Println("here1")

		// log      the     error
		log.Printf(
			`
			creation  new user's session
			after user  sign  in  failed
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
					"Вернитесь на страницу <a href='/front/signin'>авторизации</a>",
				),
			},
		)
		return
	}

	createSession(w, r, i)

}

// ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++ //
