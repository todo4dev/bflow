// presentation/http/router/error.go
package router

import (
	"bytes"
	"errors"
	"io"
	"maps"
	"net/http"
	"net/url"
	"reflect"

	"github.com/getsentry/sentry-go"
	"github.com/gofiber/fiber/v2"
	"github.com/leandroluk/gox/util"
	"github.com/leandroluk/gox/validate"
)

type ErrorMap map[string]int

func NewErrorMap() *ErrorMap {
	return &ErrorMap{}
}

func (m *ErrorMap) Map(e error, c int) *ErrorMap {
	(*m)[reflect.TypeOf(e).Name()] = c
	return m
}

func (m *ErrorMap) Status(e error, errs ...error) *ErrorMap {
	for _, err := range errs {
		m.Map(err, fiber.StatusInternalServerError)
	}
	return m
}

func (m *ErrorMap) Merge(others ...*ErrorMap) *ErrorMap {
	for _, other := range others {
		maps.Copy(*m, *other)
	}
	return m
}

var DefaultErrorMap = NewErrorMap().
	Map(&validate.ValidationError{}, fiber.StatusBadRequest)

func ErrorHandler(optionalErrorMap ...*ErrorMap) fiber.ErrorHandler {
	errorMap := DefaultErrorMap
	if len(optionalErrorMap) > 0 {
		errorMap = optionalErrorMap[0]
	}

	return func(ctx *fiber.Ctx, err error) error {
		code := fiber.StatusInternalServerError
		name := "InternalServerError"
		message := "Internal Server Error"

		var fiberErr *fiber.Error
		if errors.As(err, &fiberErr) {
			code = fiberErr.Code
			name = "FiberError"
			message = fiberErr.Message
		}

		for cursor := err; cursor != nil; cursor = errors.Unwrap(cursor) {
			t := reflect.TypeOf(cursor)
			if t.Kind() == reflect.Pointer {
				t = t.Elem()
			}

			if customStatus, ok := (*errorMap)[t.Name()]; ok {
				code = customStatus
				name = t.Name()
				message = cursor.Error()
				break
			}
		}

		if code == fiber.StatusInternalServerError {
			sentry.WithScope(func(scope *sentry.Scope) {
				scope.SetTag("method", ctx.Method())
				scope.SetTag("path", ctx.Path())
				scope.SetUser(sentry.User{IPAddress: ctx.IP()})
				scope.SetRequest(&http.Request{
					Method:     ctx.Method(),
					URL:        util.Must(url.Parse(ctx.OriginalURL())),
					RequestURI: ctx.OriginalURL(),
					RemoteAddr: ctx.IP(),
					Header:     http.Header(ctx.GetReqHeaders()),
					Body:       io.NopCloser(bytes.NewReader(ctx.Body())),
				})
				sentry.CaptureException(err)
			})
			println("[Error] 500:", err.Error())
		}

		return ctx.Status(code).JSON(fiber.Map{"name": name, "message": message})
	}
}
