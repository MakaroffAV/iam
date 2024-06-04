// ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++ //

package mail

// ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++ //

import (
	"gopkg.in/gomail.v2"
)

// ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++ //

var (
	mailPort = 587
	mailHost = "smtp.yandex.ru"
	mailPass = "ugqwzrmxxwwyiboa"
	mailUser = "hp.mail.bot@yandex.ru"
)

// ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++ //

// Send is function to send
// email  to  the recipient
func Send(recipient string, theme string, mbody string) error {

	// create  email  message
	m := gomail.NewMessage()

	// set message's  headers
	m.SetHeader("To", recipient)
	m.SetHeader("From", mailUser)
	m.SetHeader("Subject", theme)

	// set   message's   body
	m.SetBody("text/plain", mbody)

	// create connection with
	// smtp  server and  send
	// the email to recipient
	return gomail.NewDialer(mailHost, mailPort, mailUser, mailPass).DialAndSend(m)

}

// ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++ //
