package mailer

import "context"

// Sender represents an email sending service
type Sender interface {
	// Send sends an email
	Send(ctx context.Context, email Email) error

	// SendBatch sends multiple emails
	SendBatch(ctx context.Context, emails []Email) error
}
