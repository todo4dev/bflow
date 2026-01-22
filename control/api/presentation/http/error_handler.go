// presentation/http/error_handler.go
package http

import (
	"errors"
	"maps"
	stdHttp "net/http"
	"net/url"
	"reflect"

	identity "src/domain/issue"

	"github.com/getsentry/sentry-go"
	"github.com/gofiber/fiber/v2"
	"github.com/leandroluk/gox/validate"
)

type errorMap map[string]int

func (m errorMap) Map(e error, c int) errorMap {
	m[reflect.TypeOf(e).Name()] = c
	return m
}

var typeNameMap = (errorMap{}).
	Map(validate.ValidationError{}, fiber.StatusBadRequest).
	Map(&identity.AccountEmailInUse{}, fiber.StatusBadRequest).
	Map(&identity.AccountInvalidOTP{}, fiber.StatusConflict).
	Map(&identity.AccountNotFound{}, fiber.StatusNotFound).
	Map(&identity.AccountAlreadyActivated{}, fiber.StatusNotAcceptable)

func errorHandler(c *fiber.Ctx, err error) error {
	status := fiber.StatusInternalServerError
	code := "StatusInternalServerError"
	message := "Internal Server Error"

	var e *fiber.Error
	if errors.As(err, &e) {
		status = e.Code
		message = e.Message
	}

	// Check general errors (Reflection based)
	var errCursor = err
	for errCursor != nil {
		t := reflect.TypeOf(errCursor)
		if t.Kind() == reflect.Pointer {
			t = t.Elem()
		}
		if s, ok := typeNameMap[t.Name()]; ok {
			status, code, message = s, t.Name(), errCursor.Error()
			break
		}
		errCursor = errors.Unwrap(errCursor)
	}

	c.Set(fiber.HeaderContentType, fiber.MIMEApplicationJSON)

	if status == fiber.StatusInternalServerError {
		sentry.WithScope(func(scope *sentry.Scope) {
			scope.SetTag("method", c.Method())
			scope.SetTag("path", c.Path())
			scope.SetUser(sentry.User{IPAddress: c.IP()})

			u, _ := url.Parse(c.OriginalURL())

			scope.SetRequest(&stdHttp.Request{ // Use standard http.Request strict for compatibility if needed, or just set extra
				Method:     c.Method(),
				URL:        u,
				RequestURI: c.OriginalURL(),
				RemoteAddr: c.IP(),
				Header:     convertHeaders(c.GetReqHeaders()),
			})
			sentry.CaptureException(err)
		})
		println(err.Error())
	}

	return c.Status(status).JSON(fiber.Map{"code": code, "message": message})
}

func convertHeaders(headers map[string][]string) stdHttp.Header {
	h := make(stdHttp.Header)
	maps.Copy(h, headers)
	return h
}
