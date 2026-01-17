// port/mailer/attachment.go
package mailer

// Attachment represents an attachment
type Attachment struct {
	Filename    string
	ContentType string
	Content     []byte
}
