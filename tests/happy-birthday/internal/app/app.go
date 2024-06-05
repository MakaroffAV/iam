// ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++ //

package app

// ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++ //

import (
	"fmt"
	"happy-birthday/internal/dbc"
	"happy-birthday/internal/handlers"
	"happy-birthday/pkg/mail"
	"log"
	"net/http"
	"time"
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
		path:    "/back/signnout",
		handler: handlers.BackSignOut,
	},
	{
		path:    "/back/birthday/new",
		handler: handlers.BackBirthdayNew,
	},
	{
		path:    "/back/birthday/edit",
		handler: handlers.BackBirthdayEdit,
	},
	{
		path:    "/back/sigin/confirm",
		handler: handlers.BackSignInConfirm,
	},
	{
		path:    "/back/birthday/subscribe",
		handler: handlers.BackBirthdaySubscribe,
	},
	{
		path:    "/back/birthday/unsubscribe",
		handler: handlers.BackBirthdayUnSubscribe,
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
		path:    "/front/birthday/edit",
		handler: handlers.FrontBirthdayEdit,
	},
	{
		path:    "/front/signin/confirm",
		handler: handlers.FrontSignInConfirm,
	},
}

// ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++ //

func notificationsManager() {

	type notificationMetadata struct {
		Id                int
		Bday              string
		Mail              string
		Name              string
		Subscribed        int
		NotifyBefore      int
		DaysUntilBirthday int
	}

	var (
		s = `
		SELECT 
			id,
			bday,
			mail,
			name,
			subscribed,
			notify_before,
			days_until_birthday
		FROM (
			SELECT 
				b.id,
				b.bday,
				u.mail,
				b.name,
				b.subscribed,
				b.notify_before,
				DATEDIFF(
					CASE
						WHEN DATE(CONCAT(YEAR(CURDATE()),     '-', MONTH(b.bday), '-', DAY(b.bday))) >= CURDATE() 
						THEN DATE(CONCAT(YEAR(CURDATE()),     '-', MONTH(b.bday), '-', DAY(b.bday)))
						ELSE DATE(CONCAT(YEAR(CURDATE()) + 1, '-', MONTH(b.bday), '-', DAY(b.bday)))
					END, CURDATE()
				) AS days_until_birthday
			FROM 
				birthday AS b
			LEFT JOIN
				user     as u
			ON
				b.created_by = u.id 
		) AS subquery
		WHERE 
			days_until_birthday = notify_before AND subscribed = 1; 
		`

		m []notificationMetadata
	)

	c, cErr := dbc.Connection()
	if cErr != nil {
		log.Print(cErr)
	}
	defer c.Close()

	r, rErr := c.Query(s)
	if rErr != nil {
		log.Print(rErr)
	}
	defer r.Close()

	for r.Next() {
		var n notificationMetadata
		if sErr := r.Scan(
			&n.Id,
			&n.Bday,
			&n.Mail,
			&n.Name,
			&n.Subscribed,
			&n.NotifyBefore,
			&n.DaysUntilBirthday,
		); sErr != nil {
			log.Print(sErr)
			return
		}
		m = append(m, n)
	}

	for _, v := range m {
		sErr := mail.Send(
			v.Mail,
			"Уведомление о дне рождения",
			fmt.Sprintf(
				"%s будет праздновать день рождения (%s)! До дня рождения осталось %d дня (-ей)! Вы получили эту рассылку, тк хотели знать о празднике за %d дня(-ей)!",
				v.Name,
				v.Bday,
				v.DaysUntilBirthday,
				v.NotifyBefore,
			),
		)
		if sErr != nil {
			log.Print(sErr)
		}
	}

}

// ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++ //

// Run is function to define
// the  program  entry point
func Run() error {

	go func() {
		for {
			notificationsManager()
			time.Sleep(time.Hour * 24)
		}
	}()

	// set  up   web  server's
	// routes and its handlers
	for _, r := range routes {
		http.HandleFunc(r.path, r.handler)
	}

	// initialize the web  app
	return http.ListenAndServe(":8080", nil)

}

// ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++ //
