package emailh

import (
	"github.com/tyronechan294/go-comm/configh"
	"gopkg.in/gomail.v2"
)

func New(smtp configh.Smtp) *gomail.Dialer {
	return gomail.NewDialer(smtp.Host, smtp.Port, smtp.Username, smtp.Password)
}

func Send(dialer gomail.Dialer, subject string, html string, from string, to ...string) error {

	m := gomail.NewMessage()
	m.SetHeader("From", from)
	m.SetHeader(`To`, to...)
	m.SetHeader(`Subject`, subject) //主题
	m.SetBody(`text/html`, html)

	return dialer.DialAndSend(m)
}
