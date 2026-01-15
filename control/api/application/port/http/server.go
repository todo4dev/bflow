// control/api/application/port/http/server.go
package http

import (
	"context"
	"net/http"
	"time"
)

type MethodEnum string

const (
	Method_Get     MethodEnum = "GET"
	Method_Post    MethodEnum = "POST"
	Method_Put     MethodEnum = "PUT"
	Method_Patch   MethodEnum = "PATCH"
	Method_Delete  MethodEnum = "DELETE"
	Method_Head    MethodEnum = "HEAD"
	Method_Options MethodEnum = "OPTIONS"
)

type HttpRequest struct {
	Context context.Context

	Method MethodEnum
	Path   string

	Headers map[string][]string
	Params  map[string]string
	Query   map[string][]string

	Body []byte
}

type CookieOption struct {
	Expires  time.Time
	Path     string
	Domain   string
	Secure   bool
	HttpOnly bool
	SameSite http.SameSite
}

type HttpResponse interface {
	Status(code int)
	Header(key string, values ...string)
	Cookie(name, value string, options ...CookieOption)

	Write(bytes []byte) error
	JSON(value any, optionalStatusCode ...int) error
	Redirect(url string, statusCode int)

	File(filePath string, optionalDownloadName ...string) error
	Bytes(contentType string, data []byte, optionalDownloadName ...string) error
	Param(name string) (string, bool)
	Query(name string) ([]string, bool)
}

type HttpContext struct {
	Request  *HttpRequest
	Response HttpResponse
}

type HandlerFunc func(ctx *HttpContext) error
type GuardFunc func(ctx *HttpContext) error
type InterceptorFunc func(ctx *HttpContext, next HandlerFunc) error

type Chain interface {
	UseGuards(guards ...GuardFunc)
	UseInterceptors(interceptors ...InterceptorFunc)
}

type Server interface {
	Chain

	Group(prefix string) Group

	Route(method MethodEnum, path string, handler HandlerFunc) Route

	Start() error
	Stop(ctx context.Context) error
}

type Group interface {
	Chain

	Group(prefix string) Group
	Route(method MethodEnum, path string, handler HandlerFunc) Route
}

type Route interface {
	Chain
}
