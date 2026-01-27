// port/mailing/mailer.go
package mailing

import "context"

// Attachment represents an attachment
type Attachment struct {
	Filename    string
	ContentType string
	Content     []byte
}

// Email represents an email message
type Email struct {
	From        string
	To          []string
	CC          []string
	BCC         []string
	Subject     string
	Text        string
	Html        string
	Attachments []Attachment
}

// Mailer represents an email sending service
type Mailer interface {
	// Send sends an email
	Send(ctx context.Context, email Email) error
}
