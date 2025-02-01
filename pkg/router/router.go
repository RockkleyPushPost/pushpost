package router

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
)

type Routable interface {
	GetMethod() string
	FullPath() string
}

type Router struct {
	mux *fiber.App
}

func NewRouter(mux *fiber.App) *Router {
	return &Router{mux: mux}
}

func (r *Router) RegisterHttpRoute(route Routable, handler fiber.Handler) error {
	method := route.GetMethod()
	path := route.FullPath()
	switch method {
	case fiber.MethodGet:
		r.mux.Get(path, handler)
	case fiber.MethodPost:
		r.mux.Post(path, handler)
	case fiber.MethodPut:
		r.mux.Put(path, handler)
	case fiber.MethodDelete:
		r.mux.Delete(path, handler)
	case fiber.MethodPatch:
		r.mux.Patch(path, handler)
	case fiber.MethodOptions:
		r.mux.Options(path, handler)
	case fiber.MethodHead:
		r.mux.Head(path, handler)
	case fiber.MethodConnect:
		r.mux.Connect(path, handler)
	case fiber.MethodTrace:
		r.mux.Trace(path, handler)
	default:
		return fmt.Errorf("cannot register route \"%s\" with method \"%s\"", path, method)
	}
	return nil
}
