// infrastructure/mailing/smtp/mailer.go
package smtp

import (
	"bytes"
	"context"

	"src/port/mailing"

	"github.com/wneessen/go-mail"
)

type Mailer struct {
	client      *mail.Client
	fromAddress string
}

var _ mailing.Mailer = (*Mailer)(nil)

func New(config *Config) (*Mailer, error) {
	client, err := mail.NewClient(
		config.Host,
		mail.WithPort(config.Port),
		mail.WithSMTPAuth(mail.SMTPAuthPlain),
		mail.WithUsername(config.Username),
		mail.WithPassword(config.Password),
	)
	if err != nil {
		return nil, err
	}

	mailer := &Mailer{
		client:      client,
		fromAddress: config.FromAddress,
	}

	return mailer, nil
}

func (s *Mailer) toGoMailMsg(email mailing.Email) (*mail.Msg, error) {
	m := mail.NewMsg()

	from := s.fromAddress
	if email.From != "" {
		from = email.From
	}
	if err := m.From(from); err != nil {
		return nil, err
	}

	if err := m.To(email.To...); err != nil {
		return nil, err
	}

	m.Subject(email.Subject)

	if email.Text != "" {
		m.SetBodyString(mail.TypeTextPlain, email.Text)
	}

	if email.Html != "" {
		m.SetBodyString(mail.TypeTextHTML, email.Html)
	}

	for _, att := range email.Attachments {
		m.AttachReader(att.Filename, bytes.NewReader(att.Content))
	}

	return m, nil
}

func (s *Mailer) Send(ctx context.Context, email mailing.Email) error {
	msg, err := s.toGoMailMsg(email)
	if err != nil {
		return err
	}
	return s.client.DialAndSendWithContext(ctx, msg)
}
