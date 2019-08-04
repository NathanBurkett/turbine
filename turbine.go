package turbine

import (
	"fmt"
	"sync"
)

type Container struct {
	_   interface{}
	mux sync.RWMutex

	dict   map[string]Binding
	strict bool
}

// Construct a new container with set strictness
func New(strict bool, dict map[string]Binding) *Container {
	return &Container{
		strict: strict,
		dict:      dict,
	}
}

// Determine if container is strict
func (c *Container) IsStrict() bool {
	return c.strict
}

// Determine if item exists in container by name
func (c *Container) Has(name string) (ok bool) {
	c.mux.RLock()
	_, ok = c.dict[name]
	c.mux.RUnlock()

	return ok
}

// Set item in the container by name
// If the container is strict, attempting to set multiple items
// with the same name will result in an error with no item bound
// into the container
func (c *Container) Set(b Binding) (err error) {
	if c.dict == nil {
		c.dict = make(map[string]Binding)
	}

	c.mux.Lock()
	err = c.handleSet(b)
	c.mux.Unlock()

	return err
}

func (c *Container) handleSet(b Binding) error {
	if c.strict && c.dict[b.Name] != b {
		return fmt.Errorf("%s already exists", b.Name)
	}

	c.dict[b.Name] = b

	return nil
}

// Get item out of container by name. If the item was not previously bound
// into the container, 'ok' will be false
func (c *Container) Get(name string) (item interface{}, ok bool) {
	var b Binding

	c.mux.RLock()
	b, ok = c.dict[name]
	if ok {
		item = c.handleResolution(b)
	}
	c.mux.RUnlock()

	return item, ok
}

func (c *Container) handleResolution(b Binding) (item interface{}) {
	if b.BindType == SINGLETON {
		return b.Resolution
	}

	fn := b.Resolution.(Factory)

	return fn(c)
}
