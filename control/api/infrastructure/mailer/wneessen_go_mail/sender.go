// infrastructure/mailer/wneessen_go_mail/sender.go
package wneessen_go_mail

import (
	"bytes"
	"context"
	"src/port/mailer"

	"github.com/wneessen/go-mail"
)

type Sender struct {
	client      *mail.Client
	fromAddress string
}

var _ mailer.Sender = (*Sender)(nil)

func NewSender(rawConfig Config) (*Sender, error) {
	config, err := ConfigSchema.Validate(rawConfig)
	if err != nil {
		return nil, err
	}

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

	return &Sender{
		client:      client,
		fromAddress: config.FromAddress,
	}, nil
}

func (s *Sender) toGoMailMsg(email mailer.Email) (*mail.Msg, error) {
	m := mail.NewMsg()
	if err := m.From(s.fromAddress); err != nil {
		return nil, err
	}

	// Override from if provided and different?
	// The interface has "From", but usually the sender is fixed or authenticated.
	// Letting the interface "From" take precedence if permitted by SMTP server.
	if email.From != "" {
		if err := m.From(email.From); err != nil {
			return nil, err
		}
	}

	if err := m.To(email.To...); err != nil {
		return nil, err
	}
	if len(email.CC) > 0 {
		if err := m.Cc(email.CC...); err != nil {
			return nil, err
		}
	}
	if len(email.BCC) > 0 {
		if err := m.Bcc(email.BCC...); err != nil {
			return nil, err
		}
	}

	m.Subject(email.Subject)
	m.SetBodyString(mail.TypeTextPlain, email.Body)
	if email.HTMLBody != "" {
		m.SetBodyString(mail.TypeTextHTML, email.HTMLBody)
	}

	return m, nil
}

func (s *Sender) Send(ctx context.Context, email mailer.Email) error {
	msg, err := s.toGoMailMsg(email)
	if err != nil {
		return err
	}

	for _, att := range email.Attachments {
		msg.AttachReader(att.Filename, bytes.NewReader(att.Content))
	}

	return s.client.DialAndSendWithContext(ctx, msg)
}

func (s *Sender) SendBatch(ctx context.Context, emails []mailer.Email) error {
	msgs := make([]*mail.Msg, 0, len(emails))
	for _, email := range emails {
		msg, err := s.toGoMailMsg(email)
		if err != nil {
			return err
		}
		for _, att := range email.Attachments {
			msg.AttachReader(att.Filename, bytes.NewReader(att.Content))
		}
		msgs = append(msgs, msg)
	}

	return s.client.DialAndSendWithContext(ctx, msgs...)
}
