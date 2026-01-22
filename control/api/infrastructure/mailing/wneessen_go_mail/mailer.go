// infrastructure/mailing/wneessen_go_mail/mailer.go
package wneessen_go_mail

import (
	"bytes"
	"context"
	"html/template"
	"os"
	"path/filepath"
	"src/port/mailing"

	"github.com/wneessen/go-mail"
)

type Mailer struct {
	client       *mail.Client
	fromAddress  string
	templatePath string
}

var _ mailing.Mailer = (*Mailer)(nil)

func NewSender(config *Config) (*Mailer, error) {
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

	return &Mailer{
		client:       client,
		fromAddress:  config.FromAddress,
		templatePath: config.TemplatePath,
	}, nil
}

func (s *Mailer) toGoMailMsg(email mailing.Email) (*mail.Msg, error) {
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

	if email.Template != "" {
		fullPath := filepath.Join(s.templatePath, email.Template)
		content, err := os.ReadFile(fullPath)
		if err != nil {
			return nil, err
		}

		tmpl, err := template.New(email.Template).Parse(string(content))
		if err != nil {
			return nil, err
		}

		var body bytes.Buffer
		if err := tmpl.Execute(&body, email.Variables); err != nil {
			return nil, err
		}
		m.SetBodyString(mail.TypeTextHTML, body.String())
	} else if email.HTMLBody != "" {
		m.SetBodyString(mail.TypeTextHTML, email.HTMLBody)
	}

	return m, nil
}

func (s *Mailer) Send(ctx context.Context, email mailing.Email) error {
	msg, err := s.toGoMailMsg(email)
	if err != nil {
		return err
	}

	for _, att := range email.Attachments {
		msg.AttachReader(att.Filename, bytes.NewReader(att.Content))
	}

	return s.client.DialAndSendWithContext(ctx, msg)
}

func (s *Mailer) SendBatch(ctx context.Context, emails []mailing.Email) error {
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
