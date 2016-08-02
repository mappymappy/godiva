// Godiva Provide Simple Dependency Injection (DI) Container
package godiva

import (
	"fmt"
)

// Container
type Container struct {

	// Anonymous functions for object factory
	factories    map[string]func(c *Container) (interface{}, error)

	// Created Object Cache
	cacheStorage map[string]interface{}
}

type factoryFunction func(c *Container) (interface{}, error)

// CreateContainer
func CreateContainer() *Container {
	return &Container{
		map[string]func(c *Container) (interface{}, error){},
		map[string]interface{}{},
	}
}

// Register FactoryFunction
func (c *Container) Register(key string, factory factoryFunction) {
	c.factories[key] = factory
}

// Create Object By Key
func (c *Container) Create(key string, ignoreCache bool) (interface{}, error) {

	if _, exist := c.factories[key]; !exist {
		return nil, fmt.Errorf("key:%s called unregistered resource", key)
	}

	if c.cacheStorage[key] != nil && !ignoreCache {
		return c.cacheStorage[key], nil
	}

	factory := c.factories[key]
	Object, err := factory(c)
	if err != nil {
		return nil, err
	}
	c.cacheStorage[key] = Object

	return Object, nil
}
