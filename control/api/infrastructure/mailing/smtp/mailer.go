// infrastructure/mailing/smtp/mailer.go
package smtp

import (
	"bytes"
	"context"
	"fmt"
	"html/template"
	"os"
	"path/filepath"
	"sync"

	"src/port/mailing"

	"github.com/wneessen/go-mail"
)

type Mailer struct {
	client       *mail.Client
	fromAddress  string
	templatePath string
	templates    map[string]*template.Template
	mu           sync.RWMutex
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
		client:       client,
		fromAddress:  config.FromAddress,
		templatePath: config.TemplatePath,
		templates:    make(map[string]*template.Template),
	}

	if err := mailer.loadTemplates(); err != nil {
		return nil, err
	}

	return mailer, nil
}

func (s *Mailer) loadTemplates() error {
	entries, err := os.ReadDir(s.templatePath)
	if err != nil {
		return fmt.Errorf("failed to read template directory: %w", err)
	}

	for _, entry := range entries {
		if entry.IsDir() {
			continue
		}

		name := entry.Name()
		fullPath := filepath.Join(s.templatePath, name)

		content, err := os.ReadFile(fullPath)
		if err != nil {
			return fmt.Errorf("failed to read template file %s: %w", fullPath, err)
		}

		t, err := template.New(name).Parse(string(content))
		if err != nil {
			return fmt.Errorf("failed to parse template %s: %w", name, err)
		}

		s.templates[name] = t
	}

	return nil
}

func (s *Mailer) getTemplate(name string) (*template.Template, error) {
	s.mu.RLock()
	tmpl, ok := s.templates[name]
	s.mu.RUnlock()
	if !ok {
		return nil, fmt.Errorf("template not found: %s", name)
	}
	return tmpl, nil
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

	if email.Body != "" {
		m.SetBodyString(mail.TypeTextPlain, email.Body)
	}

	if email.Template != "" {
		tmpl, err := s.getTemplate(email.Template)
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

func (s *Mailer) SendBatch(ctx context.Context, emails []mailing.Email) error {
	msgs := make([]*mail.Msg, 0, len(emails))
	for _, email := range emails {
		msg, err := s.toGoMailMsg(email)
		if err != nil {
			return err
		}
		msgs = append(msgs, msg)
	}
	return s.client.DialAndSendWithContext(ctx, msgs...)
}
