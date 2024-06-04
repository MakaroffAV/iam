// ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++ //

package app

// ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++ //

import (
	"happy-birthday/internal/handlers"
	"net/http"
)

// ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++ //

// route is data structure to
// define web server's routes
// and       its     handlers
type route struct {
	path    string
	handler func(w http.ResponseWriter, r *http.Request)
}

// ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++ //

// routes defines web server's
// routes   and  its  handlers
var routes = []route{
	{
		path:    "/back/signin",
		handler: handlers.BackSignIn,
	},
	{
		path:    "/back/signup",
		handler: handlers.BackSignUp,
	},
	{
		path:    "/back/birthday/new",
		handler: handlers.BackBirthdayNew,
	},
	{
		path:    "/back/sigin/confirm",
		handler: handlers.BackSignInConfirm,
	},

	{
		path:    "/front/index",
		handler: handlers.FrontIndex,
	},
	{
		path:    "/front/signin",
		handler: handlers.FrontSignIn,
	},
	{
		path:    "/front/signup",
		handler: handlers.FrontSignUp,
	},
	{
		path:    "/front/signin/confirm",
		handler: handlers.FrontSignInConfirm,
	},
}

// ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++ //

// Run is function to define
// the  program  entry point
func Run() error {

	// set  up   web  server's
	// routes and its handlers
	for _, r := range routes {
		http.HandleFunc(r.path, r.handler)
	}

	// initialize the web  app
	return http.ListenAndServe(":8080", nil)

}

// ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++ //
