package router

import (
	"reflect"

	"github.com/leandroluk/gox/oas"
	"github.com/leandroluk/gox/validate"
)

func ErrorSchema[E error]() func(m *oas.MediaType) {
	errType := reflect.TypeFor[E]()
	for errType.Kind() == reflect.Pointer {
		errType = errType.Elem()
	}
	code := errType.Name()
	return func(m *oas.MediaType) {
		m.Schema(func(s *oas.Schema) {
			s.Object().
				Required("code", func(p *oas.Schema) { p.String().Example(code) }).
				Required("message", func(p *oas.Schema) { p.String().Example((*new(E)).Error()) })
		})
	}
}

func InPath(o *oas.Operation, name string, fn func(s *oas.Schema)) *oas.Operation {
	return o.Parameter(func(p *oas.Parameter) {
		p.Name(name).In("path").Required(true).Schema(fn)
	})
}

func BodyJson(o *oas.Operation, fn func(s *oas.Schema)) *oas.Operation {
	return o.RequestBody(func(rb *oas.RequestBody) {
		rb.Json(func(m *oas.MediaType) { m.Schema(fn) })
	})
}

func ResponseValidationError(o *oas.Operation) *oas.Operation {
	return o.Response("400", func(r *oas.Response) {
		r.Description("Validation error").Json(ErrorSchema[validate.ValidationError]())
	})
}

func ResponseIssue[E error](o *oas.Operation, code string, optionalDescription ...string) *oas.Operation {
	description := (*new(E)).Error()
	if len(optionalDescription) > 0 {
		description = optionalDescription[0]
	}
	return o.Response(code, func(r *oas.Response) {
		r.Description(description).Json(ErrorSchema[E]())
	})
}

func ResponseStatus(o *oas.Operation, code string, description string) *oas.Operation {
	return o.Response(code, func(r *oas.Response) {
		r.Description(description)
	})
}
