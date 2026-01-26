package middleware

import (
	"bytes"
	"errors"
	"io"
	"maps"
	"net/http"
	"net/url"
	"reflect"

	"src/presentation/http/common"

	"github.com/getsentry/sentry-go"
	"github.com/gofiber/fiber/v2"
	"github.com/leandroluk/gox/util"
	"github.com/leandroluk/gox/validate"
)

var DefaultErrorToStatus = common.NewErrorToStatus().
	Map(fiber.StatusBadRequest, &validate.ValidationError{})

func ErrorHandler(aditionalErrorToStatus ...*common.ErrorToStatus) fiber.ErrorHandler {
	errorToStatus := DefaultErrorToStatus
	for _, other := range aditionalErrorToStatus {
		maps.Copy(*errorToStatus, *other)
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

			if customStatus, ok := (*errorToStatus)[t.Name()]; ok {
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
