package container

import (
	"errors"
	"sync"
)

// Container holds dependencies with strict type safety.
type Container struct {
	services sync.Map
}

func NewContainer() *Container {
	return &Container{}
}

// Register stores a dependency of a specific concrete type.
func (c *Container) Register(service interface{}) {
	if service == nil {

		panic("register: nil service")
	}
	c.services.Store(service, service)
}

// Resolve retrieves a dependency by type.
func (c *Container) Resolve(service interface{}) (interface{}, error) {
	if service == nil {

		return nil, errors.New("service cannot be nil")
	}
	// Look for the exact type
	if resolved, ok := c.services.Load(service); ok {

		return resolved, nil
	}

	return nil, errors.New("service not found: " + service.(string))
}

// MustResolve retrieves a service and panics if not found.
func (c *Container) MustResolve(service interface{}) interface{} {
	resolved, err := c.Resolve(service)

	if err != nil {

		panic(err)
	}

	return resolved
}
