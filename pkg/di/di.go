package di

import (
	"errors"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"log"
	"pushpost/pkg/middleware"
	"reflect"
	"strings"
	"sync"
)

type DI struct {
	mu           sync.RWMutex
	dependencies map[reflect.Type]interface{}
	app          *fiber.App
	jwtSecret    string
}

func NewDI(app *fiber.App, jwtSecret string) *DI {

	return &DI{
		dependencies: make(map[reflect.Type]interface{}),
		jwtSecret:    jwtSecret,
		app:          app,
	}
}

func (di *DI) Register(dependencies ...interface{}) error {
	di.mu.Lock()
	defer di.mu.Unlock()
	for _, dependency := range dependencies {
		if dependency == nil {

			return errors.New("cannot register nil dependency")
		}
		t := reflect.TypeOf(dependency)
		di.dependencies[t] = dependency
		log.Printf("Registered dependency type: %v", t)
	}

	return nil
}

func (di *DI) Bind(targets ...interface{}) error {
	for _, target := range targets {
		if err := di.bindTarget(target); err != nil {
			return err
		}
	}

	return nil
}

func (di *DI) bindTarget(target interface{}) error {
	val := reflect.ValueOf(target).Elem()
	typ := val.Type()

	for i := 0; i < val.NumField(); i++ {
		if err := di.bindField(val.Field(i), typ.Field(i)); err != nil {

			return err
		}
	}

	return nil
}

func (di *DI) bindField(field reflect.Value, fieldType reflect.StructField) error {
	tag, ok := fieldType.Tag.Lookup("bind")

	if !ok {

		return nil
	}

	if dep, exists := di.dependencies[field.Type()]; exists {
		field.Set(reflect.ValueOf(dep))
		log.Printf("Bound dependency using tag: %s", tag)

		return nil
	}

	for depType, dep := range di.dependencies {
		if field.Type().Kind() == reflect.Interface && depType.Implements(field.Type()) {
			field.Set(reflect.ValueOf(dep))
			log.Printf("Bound interface dependency using tag: %s", tag)

			return nil
		}
	}

	return fmt.Errorf("dependency not found for field %s with tag %s", fieldType.Name, tag)
}

func (di *DI) RegisterRoutes(routes interface{}, pathPrefix string) error {
	val := reflect.ValueOf(routes)
	typ := val.Type()

	if val.Kind() != reflect.Struct {

		return errors.New("routes must be a struct")
	}

	for j := 0; j < val.NumField(); j++ {
		field := val.Field(j)
		fieldType := typ.Field(j)

		// Check if the field has a `method` and 'secure' tags
		method, hasMethod := fieldType.Tag.Lookup("method")
		secure, hasSecure := fieldType.Tag.Lookup("secure")

		if hasMethod {
			path := pathPrefix + "/" + fieldType.Name
			handler := field.Interface().(fiber.Handler)

			var handlers []fiber.Handler
			if hasSecure && secure == "true" {
				handlers = append(handlers, middleware.AuthJWTMiddleware(di.jwtSecret))
			}
			handlers = append(handlers, handler)

			switch method {
			case "GET":
				di.app.Get(path, handlers...)
			case "POST":
				di.app.Post(path, handlers...)
			case "PUT":
				di.app.Put(path, handlers...)
			case "DELETE":
				di.app.Delete(path, handlers...)
			case "PATCH":
				di.app.Patch(path, handlers...)
			case "HEAD":
				di.app.Head(path, handlers...)
			case "OPTIONS":
				di.app.Options(path, handlers...)
			case "CONNECT":
				di.app.Connect(path, handlers...)
			case "TRACE":
				di.app.Trace(path, handlers...)
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
