// ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++ //

package handlers

// ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++ //

import (
	"html/template"
	"net/http"
)

// ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++ //

// BackSignIn is  function to
// create new session for the
// user by the  email address
func BackSignIn(w http.ResponseWriter, r *http.Request) {

	// check the  method of
	// the client's request
	if r.Method == http.MethodPost {
		w.WriteHeader(http.StatusMethodNotAllowed)
		w.Write()
	}

	w.Write([]byte("here"))

}

// ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++ //

// FrontSignIn is function to
// generate HTML template for
// the user's  authentication
func FrontSignIn(w http.ResponseWriter, r *http.Request) {

	// check the  method of
	// the client's request
	if r.Method != http.MethodGet {
		w.Write([]byte("Неверный"))
		return
	}

	t, tErr := template.ParseFiles("./templates/signin.html")
	if tErr != nil {

	}

	t.Execute(w, nil)

}

// ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++ //
