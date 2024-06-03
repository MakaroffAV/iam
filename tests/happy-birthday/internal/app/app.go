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
		path:    "/back/signup",
		handler: handlers.BackSignUp,
	},

	{
		path:    "/front/signin",
		handler: handlers.FrontSignIn,
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

	// initialize the  web app
	return http.ListenAndServe(":8080", nil)

}

// ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++ //
