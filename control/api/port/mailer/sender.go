package mailer

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
	Body        string
	HTMLBody    string
	Attachments []Attachment
}

// Sender represents an email sending service
type Sender interface {
	// Send sends an email
	Send(ctx context.Context, email Email) error

	// SendBatch sends multiple emails
	SendBatch(ctx context.Context, emails []Email) error
}
