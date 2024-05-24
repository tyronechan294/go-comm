package emailh

import (
	"gopkg.in/gomail.v2"
)

func Send(dialer gomail.Dialer, subject string, html string, from string, to ...string) error {

	m := gomail.NewMessage()
	m.SetHeader("From", from)
	m.SetHeader(`To`, to...)
	m.SetHeader(`Subject`, subject) //主题
	m.SetBody(`text/html`, html)

	return dialer.DialAndSend(m)
}
