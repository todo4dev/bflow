// port/mailer/email.go
package mailer

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
