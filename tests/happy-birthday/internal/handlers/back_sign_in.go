// ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++ //

package handlers

// ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++ //

import (
	"fmt"
	"happy-birthday/internal/models"
	"happy-birthday/pkg/mail"
	"html/template"
	"log"
	"net/http"
	"time"

	"github.com/google/uuid"
)

// ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++ //

// userExists   is    function  to
// check if user already signed up
func userSignedUp(email string, w http.ResponseWriter) bool {

	// check     if     user
	// has  been  signed  up
	e, eErr := models.User{Mail: email}.Exists(dbConn)
	if eErr != nil {

		// log   the   error
		log.Printf(
			`
			checking  if   user  has
			been signed up failed %s
			`,
			eErr.Error(),
		)

		// build   and  send
		// the HTML template
		showStatusPage(
			w,
			&statusPage{
				Title:  "Ошибка",
				Header: "Произошла внутренняя ошибка",
				Message: template.HTML(
					"Вернитесь на страницу <a href='/front/signin'>авторизации</a>",
				),
			},
		)
	}

	return e

}

// ------------------------------------------------------------------------ //

// prepareUserSignIn is function to
// send one-time code to  the  user
// mail   and   redirect   to   the
// /front/signin/confirm      route
func prepareUserSignIn(email string, w http.ResponseWriter, r *http.Request) {

	var (

		// c defines the one-time
		// sign      in      code
		c = time.Now().Unix()

		// t defines the  sign in
		// user's           token
		t = uuid.New().String()
	)

	// write sign up  code to the
	// user_code  database  table
	iErr := models.UserCode{
		Mail:  email,
		Code:  c,
		Token: t,
	}.Insert(dbConn)
	if iErr != nil {

		// log      the     error
		log.Printf(
			`
			inserting sign in code
			to the  database table
			failed              %s
			`,
			iErr.Error(),
		)

		// build  and  write  the
		// HTML          template
		showStatusPage(
			w,
			&statusPage{
				Title:  "Ошибка",
				Header: "Произошла внутренняя ошибка",
				Message: template.HTML(
					"Вернитесь на страницу <a href='/front/signin'>авторизации</a>",
				),
			},
		)
		return
	}

	// send sing in  code  to the
	// passed    email    address
	sErr := mail.Send(
		email,
		"Подтвердите вход",
		fmt.Sprintf("Ваш код для авторизации: %d", c),
	)
	if sErr != nil {

		// log      the     error
		log.Printf(
			`
			sending  sign in  code
			to the  user failed %s
			`,
			sErr.Error(),
		)

		// build   and  send
		// the HTML template
		showStatusPage(w, &statusPage{
			Title:   "Ошибка",
			Header:  "Произошла ошибка",
			Message: template.HTML("Вернитесь на страницу <a href='/front/signin'>авторизации</a>"),
		})
		return
	}

	// set up sign up cookies
	// and  redirect   to the
	// signup            page
	http.SetCookie(
		w,
		&http.Cookie{
			Value:   t,
			Path:    "/",
			Name:    "signInToken",
			Expires: time.Now().Add(time.Hour),
		},
	)
	http.Redirect(w, r, "/front/signin/confirm", http.StatusFound)

}

// ------------------------------------------------------------------------ //

// prepareUserSignUp is function
// to send one-time code  to the
// user  mail  and  redirect  to
// the    /front/signup    route
func prepareUserSignUp(email string, w http.ResponseWriter, r *http.Request) {

	var (

		// c defines the one-time
		// sign      in      code
		c = time.Now().Unix()

		// t defines the  sign in
		// user's           token
		t = uuid.New().String()
	)

	// write sign up  code to the
	// user_code  database  table
	iErr := models.UserCode{
		Mail:  email,
		Code:  c,
		Token: t,
	}.Insert(dbConn)
	if iErr != nil {

		// log      the     error
		log.Printf(
			`
			inserting sign up code
			to the  database table
			failed              %s
			`,
			iErr.Error(),
		)

		// build  and  write  the
		// HTML          template
		showStatusPage(
			w,
			&statusPage{
				Title:  "Ошибка",
				Header: "Произошла внутренняя ошибка",
				Message: template.HTML(
					"Вернитесь на страницу <a href='/front/signin'>авторизации</a>",
				),
			},
		)
		return
	}

	// send sign up  code  to the
	// passed    email    address
	sErr := mail.Send(
		email,
		"Завершите регистрацию",
		fmt.Sprintf("Ваш код для завершения регистрации: %d", c),
	)
	if sErr != nil {

		// log      the     error
		log.Printf(
			`
			sending  sign up  code
			to the  user failed %s
			`,
			sErr.Error(),
		)

		// build   and  send
		// the HTML template
		showStatusPage(w, &statusPage{
			Title:   "Ошибка",
			Header:  "Произошла ошибка",
			Message: template.HTML("Вернитесь на страницу <a href='/front/signin'>авторизации</a>"),
		})
		return
	}

	// set up sign up cookies
	// and  redirect   to the
	// signup            page
	http.SetCookie(
		w,
		&http.Cookie{
			Value:   t,
			Path:    "/",
			Name:    "signUpToken",
			Expires: time.Now().Add(time.Hour),
		},
	)
	http.Redirect(w, r, "/front/signup", http.StatusFound)

}

// ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++ //

// BackSignIn is  function to
// create new session for the
// user by the  email address
func BackSignIn(w http.ResponseWriter, r *http.Request) {

	// check  the  method  of
	// the  client's  request
	if r.Method != http.MethodPost {

		// build   and   send
		// the  HTML template
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

	// check  if  form  value
	// is      not      empty
	m := r.FormValue("email")
	if len(m) == 0 {

		// log    the   error
		log.Print(
			`
			parsed user email from
			the HTML form is empty
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
					"Вернитесь на страницу <a href='/front/signin'>авторизации</a>",
				),
			},
		)
		return
	}

	// check if user has been
	// already   signed    up
	if userSignedUp(m, w) {
		prepareUserSignIn(m, w, r)
	} else {
		prepareUserSignUp(m, w, r)
	}

}

// ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++ //
