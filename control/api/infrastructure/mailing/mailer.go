// application/port/mailing/mailer.go
package mailing

type Mailer interface {
	Send(to string, subject string, html string) error
}
