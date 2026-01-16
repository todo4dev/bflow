// domain/shared/issue.go
package shared

import (
	"reflect"

	"github.com/leandroluk/go/mut"
)

type DomainIssue[T any] struct {
	Code    mut.Mut[string] `json:"code"`
	Message mut.Mut[string] `json:"message"`
	Cause   mut.Mut[error]  `json:"-"`
}

var _ error = (*DomainIssue[any])(nil)
var _ interface{ Unwrap() []error } = (*DomainIssue[any])(nil)

func (i *DomainIssue[T]) Name() string {
	var zero T
	t := reflect.TypeOf(zero)
	if t == nil {
		return ""
	}
	if t.Kind() == reflect.Pointer {
		t = t.Elem()
	}
	return t.Name()
}

func (i *DomainIssue[T]) Error() string {
	if i.Cause.Dirty() && i.Cause.Get() != nil {
		return i.Cause.Get().Error()
	}
	return i.Message.Get()
}

func (i *DomainIssue[T]) Unwrap() []error {
	var errorList []error
	if i.Cause.Dirty() {
		errorList = append(errorList, i.Cause.Get())
	}
	return errorList
}
