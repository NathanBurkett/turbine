package turbine

import (
	"fmt"
	"sync"
)

type Container struct {
	_   interface{}
	mux sync.RWMutex

	d      map[string]interface{}
	strict bool
}

// Construct a new container with set strictness
func New(strict bool, dict map[string]interface{}) *Container {
	return &Container{
		strict: strict,
		d:      dict,
	}
}

// Determine if container is strict
func (c *Container) IsStrict() bool {
	return c.strict
}

// Determine if item exists in container by name
func (c *Container) Has(name string) (ok bool) {
	c.mux.RLock()
	_, ok = c.d[name]
	c.mux.RUnlock()

	return ok
}

// Set item in the container by name
// If the container is strict, attempting to set multiple items
// with the same name will result in an error w/o a set operation
func (c *Container) Set(name string, item interface{}) (err error) {
	if c.d == nil {
		c.d = make(map[string]interface{})
	}

	c.mux.Lock()
	err = c.handleSet(name, item)
	c.mux.Unlock()

	return err
}

func (c *Container) handleSet(name string, item interface{}) error {
	if c.strict && c.d[name] != nil {
		return fmt.Errorf("%s already exists", name)
	}

	c.d[name] = item

	return nil
}

// Get item out of container by name. If the item was not previously bound
// into the container, 'ok' will be false
func (c *Container) Get(name string) (item interface{}, ok bool) {
	c.mux.RLock()
	item, ok = c.d[name]
	c.mux.RUnlock()

	return item, ok
}
