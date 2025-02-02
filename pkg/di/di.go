package di

import (
	"errors"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"log"
	"reflect"
	"strings"
	"sync"
)

type DI struct {
	sync.RWMutex
	dependencies map[reflect.Type]interface{}
	app          *fiber.App
}

func NewDI(app *fiber.App) *DI {
	return &DI{
		dependencies: make(map[reflect.Type]interface{}),
		app:          app,
	}
}

func (di *DI) Register(dependencies ...interface{}) error {
	di.Lock()
	defer di.Unlock()
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
