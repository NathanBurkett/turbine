package turbine

import "reflect"

const SINGLETON = 1
const FACTORY = 2

type BindingMap map[string]Binding

type Binding struct {
	_          interface{}
	Name       string
	Resolution interface{}
	BindType   uint16
}

type Singleton interface{}

type Factory func(c *Container) interface{}

func NewBinding(name string, resolution interface{}) Binding {
	b := Binding{
		Name:       name,
		Resolution: resolution,
		BindType:   SINGLETON,
	}

	if reflect.TypeOf(resolution).Kind() == reflect.Func {
		b.BindType = FACTORY
		return b
	}

	return b
}

func (b Binding) IsSingleton() bool {
	return b.BindType == SINGLETON
}

func (b Binding) IsFactory() bool {
	return b.BindType == FACTORY
}
