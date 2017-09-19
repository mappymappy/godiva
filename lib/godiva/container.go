// Godiva Provide Simple Dependency Injection (DI) Container
package godiva

import (
	"fmt"
	"sync"
)

var mutex sync.Mutex
var cachem sync.Mutex

// Container
type Container struct {

	// Anonymous functions for object factory
	factories map[string]func(c *Container) (interface{}, error)

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
	mutex.Lock()
	defer mutex.Unlock()
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
	cachem.Lock()
	c.cacheStorage[key] = Object
	cachem.Unlock()

	return Object, nil
}
