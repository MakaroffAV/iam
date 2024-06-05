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

// BackBirthdayNew  is function
// to create new 'birthday' man
func BackBirthdayNew(w http.ResponseWriter, r *http.Request) {

	var (

		// s defines the user's
		// session        token
		s = ""

		// n defines the form's
		// param with name=name
		n = ""

		// b defines the form's
		// param with name=bday
		b = ""

		// d defines the form's
		// param with name=notify_before
		d = ""
	)

	// check  the   method  of
	// the   client's  request
	if r.Method != http.MethodPost {

		// build    and   send
		// the  HTML  template
		showStatusPage(w, &statusPage{
			Title:  "Ошибка",
			Header: "Недопустимый метод",
			Message: template.HTML(
				"Вернитесь на <a href='/front/index'>главную страницу</a>",
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

	// retrieve   form  params
	n = r.FormValue("name")
	b = r.FormValue("bday")
	d = r.FormValue("notify_before")
	for _, v := range r.Cookies() {
		if v.Name == "sessionToken" {
			s = v.Value
		}
	}
	if len(n) == 0 || len(b) == 0 || len(d) == 0 || len(s) == 0 {

		// log    the   error
		log.Print(
			`
			parsed  'birthday man'
			params are not allowed
			`,
		)

		// build   and   send
		// the  HTML template
		showStatusPage(
			w,
			&statusPage{
				Title:  "Ошибка",
				Header: "Недопустимое значение",
				Message: template.HTML(
					"Вернитесь на <a href='/front/index'>главную страницу</a>",
				),
			},
		)
		return
	}

	// convert notify before to int
	x, xErr := strconv.Atoi(d)
	if xErr != nil {

		// log    the   error
		log.Print(
			`
			converting  notify before
			from string to int failed
			`,
		)

		// build   and   send
		// the  HTML template
		showStatusPage(
			w,
			&statusPage{
				Title:  "Ошибка",
				Header: "Недопустимое значение",
				Message: template.HTML(
					"Вернитесь на <a href='/front/index'>главную страницу</a>",
				),
			},
		)
		return
	}

	// create new 'birthday man'
	cErr := models.Birthday{Name: n, BDay: b, NotifyBefore: x}.Create(dbConn, s)
	if cErr != nil {

		// log    the   error
		log.Print(
			`
			writing  new  birthday
			man to database failed
			`,
		)

		// build   and   send
		// the  HTML template
		showStatusPage(
			w,
			&statusPage{
				Title:  "Ошибка",
				Header: "Ошибка на сервере",
				Message: template.HTML(
					"Вернитесь на <a href='/front/index'>главную страницу</a>",
				),
			},
		)
		return
	}

	http.Redirect(w, r, "/front/index", http.StatusFound)

}

// ------------------------------------------------------------------------ //

func BackBirthdaySubscribe(w http.ResponseWriter, r *http.Request) {

	i, iErr := strconv.Atoi(r.URL.Query().Get("id"))
	if iErr != nil {
		showStatusPage(
			w,
			&statusPage{
				Title:  "Ошибка",
				Header: "Ошибка на сервере",
				Message: template.HTML(
					"Вернитесь на <a href='/front/index'>главную страницу</a>",
				),
			},
		)
		return
	}

	uErr := models.Birthday{Id: i, Subscribed: 1}.UpdateSubscription(dbConn)
	if uErr != nil {
		showStatusPage(
			w,
			&statusPage{
				Title:  "Ошибка",
				Header: "Ошибка на сервере",
				Message: template.HTML(
					"Вернитесь на <a href='/front/index'>главную страницу</a>",
				),
			},
		)
		return
	}

	http.Redirect(w, r, "/front/index", http.StatusFound)

}

// ------------------------------------------------------------------------ //

func BackBirthdayUnSubscribe(w http.ResponseWriter, r *http.Request) {

	i, iErr := strconv.Atoi(r.URL.Query().Get("id"))
	if iErr != nil {
		showStatusPage(
			w,
			&statusPage{
				Title:  "Ошибка",
				Header: "Ошибка на сервере",
				Message: template.HTML(
					"Вернитесь на <a href='/front/index'>главную страницу</a>",
				),
			},
		)
		return
	}

	uErr := models.Birthday{Id: i, Subscribed: 0}.UpdateSubscription(dbConn)
	if uErr != nil {
		showStatusPage(
			w,
			&statusPage{
				Title:  "Ошибка",
				Header: "Ошибка на сервере",
				Message: template.HTML(
					"Вернитесь на <a href='/front/index'>главную страницу</a>",
				),
			},
		)
		return
	}

	http.Redirect(w, r, "/front/index", http.StatusFound)

}

// ------------------------------------------------------------------------ //

func BackBirthdayEdit(w http.ResponseWriter, r *http.Request) {

	var s = ""

	if r.Method != http.MethodPost {
		log.Print(r.Method)
		showStatusPage(w, &statusPage{
			Title:  "Ошибка",
			Header: "Недопустимый метод",
			Message: template.HTML(
				"Вернитесь на <a href='/front/index'>главную страницу</a>",
			),
		})
		return
	}

	i, iErr := strconv.Atoi(r.URL.Query().Get("id"))
	if iErr != nil {
		log.Print(iErr)
		showStatusPage(
			w,
			&statusPage{
				Title:  "Ошибка",
				Header: "Ошибка на сервере",
				Message: template.HTML(
					"Вернитесь на <a href='/front/index'>главную страницу</a>",
				),
			},
		)
		return
	}

	n := r.FormValue("name")
	b := r.FormValue("bday")
	d := r.FormValue("notify_before")
	for _, v := range r.Cookies() {
		if v.Name == "sessionToken" {
			s = v.Value
		}
	}

	x, xErr := strconv.Atoi(d)
	if xErr != nil {
		log.Print(xErr)
		showStatusPage(
			w,
			&statusPage{
				Title:  "Ошибка",
				Header: "Ошибка на сервере",
				Message: template.HTML(
					"Вернитесь на <a href='/front/index'>главную страницу</a>",
				),
			},
		)
		return
	}

	uErr := models.Birthday{Id: i, BDay: b, Name: n, NotifyBefore: x}.Update(dbConn, s)
	if uErr != nil {
		log.Print(uErr)
		showStatusPage(
			w,
			&statusPage{
				Title:  "Ошибка",
				Header: "Ошибка на сервере",
				Message: template.HTML(
					"Вернитесь на <a href='/front/index'>главную страницу</a>",
				),
			},
		)
		return
	}

	http.Redirect(w, r, "/front/index", http.StatusFound)

}

// ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++ //
