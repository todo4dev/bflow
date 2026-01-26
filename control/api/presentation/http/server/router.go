// presentation/http/server/router.go
package server

import (
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/leandroluk/gox/oas"
)

type Router struct {
	fiber    fiber.Router
	path     string
	specs    []*Spec
	children []*Router
	Config   *Config
}

func (r *Router) Group(def Group) *Router {
	fullPath := r.path + def.Path
	if !strings.HasPrefix(fullPath, "/") {
		fullPath = "/" + fullPath
	}

	child := &Router{
		fiber:  r.fiber.Group(def.Path),
		path:   fullPath,
		Config: r.Config,
	}
	r.children = append(r.children, child)

	if def.Factory != nil {
		def.Factory(child)
	}

	return r
}

func (r *Router) addSpec(path, method string, spec func(*oas.Operation)) {
	if spec == nil {
		return
	}

	full := strings.ReplaceAll(r.path+path, "//", "/")
	r.specs = append(r.specs, &Spec{
		Path:        PARAM_REGEX.ReplaceAllString(full, "{$1}"),
		Method:      method,
		OperationFn: spec,
	})
}

func (r *Router) Use(handlers ...interface{}) *Router {
	r.fiber.Use(handlers...)
	return r
}

func (r *Router) wrapHandlers(handlers []Handler) []fiber.Handler {
	fiberHandlers := make([]fiber.Handler, len(handlers))
	for i, h := range handlers {
		fiberHandlers[i] = func(c *fiber.Ctx) error {
			return h(&Context{Ctx: c, Config: r.Config, Router: r})
		}
	}
	return fiberHandlers
}

func (r *Router) Get(path string, route *Route) *Router {
	r.fiber.Get(path, r.wrapHandlers(route.handlers)...)
	r.addSpec(path, "GET", route.operation)
	return r
}

func (r *Router) Post(path string, route *Route) *Router {
	r.fiber.Post(path, r.wrapHandlers(route.handlers)...)
	r.addSpec(path, "POST", route.operation)
	return r
}

func (r *Router) Put(path string, route *Route) *Router {
	r.fiber.Put(path, r.wrapHandlers(route.handlers)...)
	r.addSpec(path, "PUT", route.operation)
	return r
}

func (r *Router) Patch(path string, route *Route) *Router {
	r.fiber.Patch(path, r.wrapHandlers(route.handlers)...)
	r.addSpec(path, "PATCH", route.operation)
	return r
}

func (r *Router) Delete(path string, route *Route) *Router {
	r.fiber.Delete(path, r.wrapHandlers(route.handlers)...)
	r.addSpec(path, "DELETE", route.operation)
	return r
}

func (r *Router) Options(path string, route *Route) *Router {
	r.fiber.Options(path, r.wrapHandlers(route.handlers)...)
	r.addSpec(path, "OPTIONS", route.operation)
	return r
}

func (r *Router) Head(path string, route *Route) *Router {
	r.fiber.Head(path, r.wrapHandlers(route.handlers)...)
	r.addSpec(path, "HEAD", route.operation)
	return r
}

func (r *Router) collectSpecs() []*Spec {
	all := make([]*Spec, 0, len(r.specs))
	all = append(all, r.specs...)
	for _, child := range r.children {
		all = append(all, child.collectSpecs()...)
	}
	return all
}

func (r *Router) GenerateJSON() *oas.Document {
	doc := oas.New()
	doc.Info(func(i *oas.Info) {
		i.Title(r.Config.OpenAPITitle)
		i.Version(r.Config.OpenAPIVersion)
		i.Description(r.Config.OpenAPIDescription)
		i.Contact().
			Name(r.Config.OpenAPIContactName).
			Email(r.Config.OpenAPIContactEmail).
			URL(r.Config.OpenAPIContactURL)
		i.License().
			Name(r.Config.OpenAPILicenseName).
			URL(r.Config.OpenAPILicenseURL)
	})

	specsByPath := make(map[string][]*Spec)
	for _, spec := range r.collectSpecs() {
		specsByPath[spec.Path] = append(specsByPath[spec.Path], spec)
	}

	for path, entries := range specsByPath {
		doc.Path(path, func(p *oas.Path) {
			for _, spec := range entries {
				method := strings.ToUpper(spec.Method)
				opCallback := spec.OperationFn

				switch method {
				case "GET":
					p.Get(opCallback)
				case "POST":
					p.Post(opCallback)
				case "PUT":
					p.Put(opCallback)
				case "PATCH":
					p.Patch(opCallback)
				case "DELETE":
					p.Delete(opCallback)
				case "OPTIONS":
					p.Options(opCallback)
				case "HEAD":
					p.Head(opCallback)
				}
			}
		})
	}

	return doc
}
