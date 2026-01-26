package common

import (
	"reflect"
)

type ErrorToStatus map[string]int

func NewErrorToStatus() *ErrorToStatus {
	return &ErrorToStatus{}
}

func (s *ErrorToStatus) Map(status int, errs ...error) *ErrorToStatus {
	for _, err := range errs {
		(*s)[reflect.TypeOf(err).Name()] = status
	}
	return s
}
