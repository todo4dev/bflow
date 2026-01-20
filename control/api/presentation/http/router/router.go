package router

import (
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/leandroluk/gox/di"
	"github.com/leandroluk/gox/env"
	"github.com/leandroluk/gox/oas"
	"github.com/leandroluk/gox/oas/types"
)

type SpecEntry struct {
	Path   string
	Method string
	Spec   func(*oas.Operation)
}

type Router struct {
	app      *fiber.App
	fiber    fiber.Router
	path     string
	specs    []*SpecEntry
	children []*Router
}

type GroupRouter = Router

type GroupDefinition struct {
	Path    string
	Factory func(*GroupRouter)
}

type RouteDefinition struct {
	handlers  []fiber.Handler
	operation func(*oas.Operation)
}

func Group(path string, factory func(*GroupRouter)) GroupDefinition {
	return GroupDefinition{Path: path, Factory: factory}
}

func Route(handlers ...fiber.Handler) *RouteDefinition {
	return &RouteDefinition{handlers: handlers}
}

func (rd *RouteDefinition) Operation(spec func(*oas.Operation)) *RouteDefinition {
	rd.operation = spec
	return rd
}

func Wrapper(app *fiber.App) *Router {
	c := Config{
		Port:                env.Get("SERVER_PORT", 3000),
		BasePath:            env.Get("SERVER_BASE_PATH", "/"),
		EnableSwagger:       env.Get("SERVER_ENABLE_SWAGGER", "true") == "true",
		SwaggerPath:         env.Get("SERVER_SWAGGER_PATH", "/docs"),
		SwaggerTitle:        env.Get("SERVER_SWAGGER_TITLE", "Bflow - Control Plane API"),
		SwaggerDescription:  env.Get("SERVER_SWAGGER_DESCRIPTION", "API for Control Plane of Bflow solution"),
		SwaggerContactName:  env.Get("SERVER_SWAGGER_CONTACT_NAME", "Leandro Santiago Gomes"),
		SwaggerContactURL:   env.Get("SERVER_SWAGGER_CONTACT_URL", "https://github.com/leandroluk"),
		SwaggerContactEmail: env.Get("SERVER_SWAGGER_CONTACT_EMAIL", "leandroluk@gmail.com"),
		SwaggerLicenseName:  env.Get("SERVER_SWAGGER_LICENSE_NAME", "MIT"),
		SwaggerLicenseURL:   env.Get("SERVER_SWAGGER_LICENSE_URL", "https://opensource.org/licenses/MIT"),
		SwaggerVersion:      env.Get("SERVER_SWAGGER_VERSION", "1.0.0"),
	}
	if _, err := ConfigSchema.Validate(c); err != nil {
		panic(err)
	}

	r := Router{app: app, fiber: app, path: ""}

	di.RegisterAs[*Config](func() *Config { return &c })
	di.RegisterAs[*Router](func() *Router { return &r })

	return &r
}

func (r *Router) Group(def GroupDefinition) *Router {
	fullPath := r.path + def.Path
	if !strings.HasPrefix(fullPath, "/") {
		fullPath = "/" + fullPath
	}

	child := &Router{
		app:   r.app,
		fiber: r.fiber.Group(def.Path),
		path:  fullPath,
	}
	r.children = append(r.children, child)

	if def.Factory != nil {
		def.Factory(child)
	}

	return r
}

func (r *Router) addSpec(subPath, method string, spec func(*oas.Operation)) {
	if spec == nil {
		return
	}

	full := r.path + subPath
	full = strings.ReplaceAll(full, "//", "/")

	r.specs = append(r.specs, &SpecEntry{
		Path:   full,
		Method: method,
		Spec:   spec,
	})
}

func (r *Router) Get(path string, route *RouteDefinition) *Router {
	r.fiber.Get(path, route.handlers...)
	r.addSpec(path, "GET", route.operation)
	return r
}

func (r *Router) Post(path string, route *RouteDefinition) *Router {
	r.fiber.Post(path, route.handlers...)
	r.addSpec(path, "POST", route.operation)
	return r
}

func (r *Router) Put(path string, route *RouteDefinition) *Router {
	r.fiber.Put(path, route.handlers...)
	r.addSpec(path, "PUT", route.operation)
	return r
}

func (r *Router) Patch(path string, route *RouteDefinition) *Router {
	r.fiber.Patch(path, route.handlers...)
	r.addSpec(path, "PATCH", route.operation)
	return r
}

func (r *Router) Delete(path string, route *RouteDefinition) *Router {
	r.fiber.Delete(path, route.handlers...)
	r.addSpec(path, "DELETE", route.operation)
	return r
}

func (r *Router) Options(path string, route *RouteDefinition) *Router {
	r.fiber.Options(path, route.handlers...)
	r.addSpec(path, "OPTIONS", route.operation)
	return r
}

func (r *Router) Head(path string, route *RouteDefinition) *Router {
	r.fiber.Head(path, route.handlers...)
	r.addSpec(path, "HEAD", route.operation)
	return r
}

func (r *Router) collectSpecs() []*SpecEntry {
	all := make([]*SpecEntry, 0, len(r.specs))
	all = append(all, r.specs...)
	for _, child := range r.children {
		all = append(all, child.collectSpecs()...)
	}
	return all
}

func (r *Router) GenerateOAS() *oas.Document {
	doc := oas.New()

	doc.Info(func(i *types.Info) {
		i.Title(env.Get("OPENAPI_TITLE", "API Title"))
		i.Version(env.Get("OPENAPI_VERSION", "1.0.0"))
		i.Description(env.Get("OPENAPI_DESCRIPTION", "API Description"))
		i.Contact().
			Name(env.Get("OPENAPI_CONTACT_NAME", "")).
			Email(env.Get("OPENAPI_CONTACT_EMAIL", "")).
			URL(env.Get("OPENAPI_CONTACT_URL", ""))

		i.License().
			Name(env.Get("OPENAPI_LICENSE_NAME", "MIT")).
			URL(env.Get("OPENAPI_LICENSE_URL", ""))
	})

	specs := r.collectSpecs()

	// Group specs by path
	specsByPath := make(map[string][]*SpecEntry)
	for _, spec := range specs {
		specsByPath[spec.Path] = append(specsByPath[spec.Path], spec)
	}

	for path, entries := range specsByPath {
		doc.Path(path, func(p *oas.Path) {
			for _, spec := range entries {
				method := strings.ToUpper(spec.Method)
				opCallback := spec.Spec

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

func (r *Router) Listen(addr string) error {
	return r.app.Listen(addr)
}
