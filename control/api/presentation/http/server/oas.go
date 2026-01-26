// presentation/http/server/oas.go
package server

import (
	"reflect"
	"strconv"

	"github.com/leandroluk/gox/meta"
	"github.com/leandroluk/gox/oas"
)

func ErrorSchemaOf(err error) func(m *oas.MediaType) {
	errType := reflect.TypeOf(err)
	for errType.Kind() == reflect.Pointer {
		errType = errType.Elem()
	}
	name := errType.Name()
	return func(m *oas.MediaType) {
		m.Schema(func(s *oas.Schema) {
			s.Object().
				Required("name", func(p *oas.Schema) { p.String().Example(name) }).
				Required("message", func(p *oas.Schema) {
					example := ""
					if err != nil {
						val := reflect.ValueOf(err)
						if val.Kind() != reflect.Pointer || !val.IsNil() {
							example = err.Error()
						}
					}
					p.String().Example(example)
				})
		})
	}
}

func ErrorSchemaAs[E error]() func(m *oas.MediaType) {
	var err E
	return ErrorSchemaOf(err)
}

func InPath(o *oas.Operation, name string, fn func(s *oas.Schema)) *oas.Operation {
	return o.Parameter(func(p *oas.Parameter) {
		p.Name(name).In("path").Required(true).Schema(fn)
	})
}
func InQuery(o *oas.Operation, name string, fn func(s *oas.Schema)) *oas.Operation {
	return o.Parameter(func(p *oas.Parameter) {
		p.Name(name).In("query").Required(true).Schema(fn)
	})
}

func BodyJson(o *oas.Operation, fn func(s *oas.Schema)) *oas.Operation {
	return o.RequestBody(func(rb *oas.RequestBody) {
		rb.Json(func(m *oas.MediaType) { m.Schema(fn) })
	})
}

func ResponseIssueOf(o *oas.Operation, err error, code int, optionalDescription ...string) *oas.Operation {
	description := ""
	if err != nil {
		val := reflect.ValueOf(err)
		if val.Kind() != reflect.Ptr || !val.IsNil() {
			description = err.Error()
		}
	}
	if len(optionalDescription) > 0 {
		description = optionalDescription[0]
	}
	return o.Response(strconv.Itoa(code), func(r *oas.Response) {
		r.Description(description).Json(ErrorSchemaOf(err))
	})
}

func ResponseIssueAs[E error](o *oas.Operation, code int, optionalDescription ...string) *oas.Operation {
	var err E
	return ResponseIssueOf(o, err, code, optionalDescription...)
}

func ResponseStatus(o *oas.Operation, code int, description string, optionalSchema ...func(m *oas.Schema)) *oas.Operation {
	return o.Response(strconv.Itoa(code), func(r *oas.Response) {
		r.Description(description)
		if len(optionalSchema) > 0 {
			r.Json(func(m *oas.MediaType) { m.Schema(optionalSchema[0]) })
		}
	})
}

func SchemaAs[T any](optionalFn ...func(s *oas.Schema)) func(s *oas.Schema) {
	return func(s *oas.Schema) {
		t := reflect.TypeFor[T]()
		buildSchema(s, t)
		if len(optionalFn) > 0 && optionalFn[0] != nil {
			optionalFn[0](s)
		}
	}
}

func buildSchema(s *oas.Schema, t reflect.Type) {
	for t.Kind() == reflect.Pointer {
		t = t.Elem()
	}

	metadata := meta.GetObjectMetadataByType(t)
	if metadata != nil && metadata.Example != nil {
		s.Example(metadata.Example)
	}

	switch t.Kind() {
	case reflect.String:
		s.String()
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64,
		reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		s.Integer()
	case reflect.Float32, reflect.Float64:
		s.Number()
	case reflect.Bool:
		s.Boolean()
	case reflect.Slice, reflect.Array:
		s.Array().Items(func(item *oas.Schema) {
			buildSchema(item, t.Elem())
		})
	case reflect.Map:
		s.Object()
	case reflect.Struct:
		s.Object()
		if t.Name() == "Time" && t.PkgPath() == "time" {
			s.String().Format("date-time")
			return
		}

		for i := 0; i < t.NumField(); i++ {
			field := t.Field(i)
			if !field.IsExported() {
				continue
			}

			tag := field.Tag.Get("json")
			if tag == "-" {
				continue
			}
			name := field.Name
			if tag != "" {
				parts := split(tag, ",")
				name = parts[0]
			}

			s.Property(name, func(prop *oas.Schema) {
				buildSchema(prop, field.Type)
				if metadata != nil {
					if fm, ok := metadata.Fields[field.Name]; ok {
						if fm.Example != nil {
							prop.Example(fm.Example)
						}
						if fm.Description != "" {
							prop.Description(fm.Description)
						}
						if fm.Nullable {
							prop.Nullable()
						}
					}
				}
			})
		}
	}
}

func split(s, sep string) []string {
	var result []string
	start := 0
	for i := 0; i < len(s); i++ {
		if i+len(sep) <= len(s) && s[i:i+len(sep)] == sep {
			result = append(result, s[start:i])
			start = i + len(sep)
			i += len(sep) - 1
		}
	}
	result = append(result, s[start:])
	return result
}
