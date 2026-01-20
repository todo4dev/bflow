// presentation/http/error_handler.go
package http

import (
	"errors"
	stdHttp "net/http"
	"net/url"
	"reflect"

	identity "src/domain/identity/issue"

	"github.com/getsentry/sentry-go"
	"github.com/gofiber/fiber/v2"
	v "github.com/leandroluk/gox/validate"
)

type errorMap map[string]int

func (m errorMap) Map(e error, c int) errorMap { m[reflect.TypeOf(e).Name()] = c; return m }

var typeNameMap = (errorMap{}).
	Map(v.ValidationError{}, fiber.StatusBadRequest).
	Map(identity.EmailInUse{}, fiber.StatusBadRequest)

func errorHandler(c *fiber.Ctx, err error) error {
	code := fiber.StatusInternalServerError
	message := "Internal Server Error"

	var e *fiber.Error
	if errors.As(err, &e) {
		code = e.Code
		message = e.Message
	}

	// Check general errors (Reflection based)
	var errCursor = err
	for errCursor != nil {
		t := reflect.TypeOf(errCursor)
		if t.Kind() == reflect.Pointer {
			t = t.Elem()
		}
		if status, ok := typeNameMap[t.Name()]; ok {
			code = status
			message = errCursor.Error()
			break
		}
		errCursor = errors.Unwrap(errCursor)
	}

	c.Set(fiber.HeaderContentType, fiber.MIMEApplicationJSON)

	if code == fiber.StatusInternalServerError {
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
	}

	return c.Status(code).JSON(fiber.Map{"error": true, "message": message})
}

func convertHeaders(headers map[string][]string) stdHttp.Header {
	h := make(stdHttp.Header)
	for k, v := range headers {
		h[k] = v
	}
	return h
}
