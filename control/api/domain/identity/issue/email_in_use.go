// domain/identity/issue/email_in_use.go
package issue

import "fmt"

type EmailInUse struct{ Email string }

func (e EmailInUse) Error() string { return fmt.Sprintf("email '%s' already in use", e.Email) }
