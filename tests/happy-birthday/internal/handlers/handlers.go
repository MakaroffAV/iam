// ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++ //

package handlers

// ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++ //

import (
	"database/sql"
	"happy-birthday/internal/dbc"
	"happy-birthday/internal/models"
	"html/template"
	"log"
	"net/http"
	"time"

	"github.com/google/uuid"
)

// ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++ //

var (

	// dbConn defines the instance of
	// the mysql  database connection
	dbConn *sql.DB
)

// ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++ //

// statusPage  is  function to define
// params of the status HTML template
type statusPage struct {
	Title   string
	Header  string
	Message template.HTML
}

// ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++ //

// init is function to initialize
// the      handlers      package
func init() {

	// create  instance  of  the
	// mysql database connection
	c, cErr := dbc.Connection()
	if cErr != nil {
		log.Fatalf(
			`
			creation instance of
			the  mysql  database
			connection    failed`,
		)
	}

	dbConn = c

}

// ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++ //

// checkSession is function to
// check if user's  session is
// still   alive   and  exists
func checkSession(r *http.Request) bool {

	var t string

	for _, v := range r.Cookies() {
		if v.Name == "sessionToken" {
			t = v.Value
		}
	}

	if len(t) == 0 {
		return false
	}

	c, cErr := models.UserSession{SToken: t}.Check(dbConn)
	if cErr != nil {
		return false
	}

	return c

}

// ------------------------------------------------------------------------ //

// showStatusPage is function to build
// and show  status page HTML template
func showStatusPage(w http.ResponseWriter, data *statusPage) {

	// build  and  send the
	// HTML        template
	template.Must(template.ParseFiles("./templates/status.html")).Execute(w, data)

}

// ------------------------------------------------------------------------ //

// createSession is function to
// create  new   user   session
func createSession(w http.ResponseWriter, r *http.Request, userId int) {

	var (

		// t     defines     the
		// new   session   token
		t = uuid.New().String()
	)

	// create new user's session
	cErr := models.UserSession{UserId: userId, SToken: t}.Create(dbConn)
	if cErr != nil {

		// build     and    send
		// the   HTML   template
		showStatusPage(
			w,
			&statusPage{
				Title:  "Ошибка",
				Header: "Ошибка при создании сессии",
				Message: template.HTML(
					"Вернитесь на страницу <a href='/front/signin'>авторизации</a>",
				),
			},
		)
		return

	}

	// set up  session token and
	// redirect to the main page
	http.SetCookie(
		w,
		&http.Cookie{
			Value:   t,
			Path:    "/",
			Name:    "sessionToken",
			Expires: time.Now().Add(time.Hour * 2),
		},
	)
	http.Redirect(w, r, "/front/index", http.StatusFound)

}

// ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++ //
