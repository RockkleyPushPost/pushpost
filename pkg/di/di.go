package di

import (
	"errors"
	"github.com/gofiber/fiber/v2"
	"log"
	"reflect"
	"strings"
)

type DI struct {
	dependencies map[reflect.Type]interface{}
	app          *fiber.App
}

func NewDI(app *fiber.App) *DI {
	return &DI{
		dependencies: make(map[reflect.Type]interface{}),
		app:          app,
	}
}

func (di *DI) Register(dependency interface{}) {
	di.dependencies[reflect.TypeOf(dependency)] = dependency
	t := reflect.TypeOf(dependency)
	di.dependencies[t] = dependency
	log.Printf("Registered dependency type: %v", t)

}

func (di *DI) Bind(target interface{}) error {
	val := reflect.ValueOf(target).Elem()
	typ := val.Type()

	log.Printf("Binding target type: %T", target)
	for k := range di.dependencies {
		log.Printf("Available dependency: %v", k)
	}
	for i := 0; i < val.NumField(); i++ {
		field := val.Field(i)
		fieldType := typ.Field(i)
		if tag, ok := fieldType.Tag.Lookup("bind"); ok {
			// Look for exact type match first
			log.Printf("Looking to bind field type: %v with tag: %v", field.Type(), tag)

			if dep, exists := di.dependencies[field.Type()]; exists {
				field.Set(reflect.ValueOf(dep))
				continue
			}

			// If not found, try to find by tag
			for depType, dep := range di.dependencies {
				if depType.String() == tag {
					field.Set(reflect.ValueOf(dep))
					break
				}
			}
		}
	}
	return nil
}

func (di *DI) RegisterRoutes(routes interface{}, pathPrefix string) error {
	val := reflect.ValueOf(routes)
	typ := val.Type()

	// Ensure the input is a struct
	if val.Kind() != reflect.Struct {
		return errors.New("routes must be a struct")
	}

	// Iterate over the struct fields
	for j := 0; j < val.NumField(); j++ {
		field := val.Field(j)
		fieldType := typ.Field(j)

		// Check if the field has a `method` tag
		if method, ok := fieldType.Tag.Lookup("method"); ok {
			// Construct the full path
			path := pathPrefix + "/" + fieldType.Name

			// Register the route using Fiber's routing methods
			handler := field.Interface().(fiber.Handler)
			switch method {
			case "GET":
				di.app.Get(path, handler)
			case "POST":
				di.app.Post(path, handler)
			case "PUT":
				di.app.Put(path, handler)
			case "DELETE":
				di.app.Delete(path, handler)
			case "PATCH":
				di.app.Patch(path, handler)
			case "HEAD":
				di.app.Head(path, handler)
			case "OPTIONS":
				di.app.Options(path, handler)
			case "CONNECT":
				di.app.Connect(path, handler)
			case "TRACE":
				di.app.Trace(path, handler)
			default:
				return errors.New("unsupported HTTP method: " + method)
			}
		}
	}

	return nil
}

// parseRouteTag parses the route tag into method and path
func parseRouteTag(tag string) (string, string) {
	// Example tag: "GET /users"
	parts := strings.Split(tag, " ")
	if len(parts) != 2 {
		return "", ""
	}
	return parts[0], parts[1]
}
