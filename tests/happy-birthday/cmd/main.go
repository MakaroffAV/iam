// ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++ //

package main

// ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++ //

import (
	"happy-birthday/internal/app"
	"log"
)

// ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++ //

// main is function to define
// the  program's  main  loop
func main() {

	if rErr := app.Run(); rErr != nil {
		log.Fatalf(
			`
			the server unexpectedly
			shut down with error %s`, rErr.Error(),
		)
	}

}

// ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++ //
