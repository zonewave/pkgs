package di

import (
	"reflect"

	"github.com/ybzhanghx/pkgs/di/xreflect"
	"github.com/ybzhanghx/pkgs/util/cleanup"
	"github.com/ybzhanghx/pkgs/werr"
	"go.uber.org/dig"
)

// New return a New instance
func New(opts ...HubOption) (*Hub, error) {
	hub := &Hub{
		cleanup:  cleanup.Entry{},
		provides: make([]Provided, 0),
	}
	for _, opt := range opts {
		opt.Apply(hub)
	}

	hub.container = dig.New(
		dig.DryRun(hub.validate),
	)
	for i, p := range hub.provides {
		if err := hub.usingProvide(p); err != nil {
			return nil, werr.Errorf("error after options[%d] were applied: %v", i, err)
		}
	}

	return hub, nil
}

// Hub is a directed acyclic graph of types and their dependencies.
// extend dig.Container
type Hub struct {
	container *dig.Container
	cleanup   cleanup.Entry
	provides  []Provided
	validate  bool
}

// Cleanup runs all the cleanup functions registered in the hub.
func (hub *Hub) Cleanup() {
	hub.cleanup.Run()
}

// Provided is a single constructor provided to di.
type Provided struct {
	// Constructor provided to di. This may be an di.Annotated.
	Target interface{}

	// Stack trace of where this Provided was made.
	Stack xreflect.Stack

	// IsSupply is true when the Target constructor was emitted by di.Supply.
	IsSupply bool
}

func (hub *Hub) usingProvide(p Provided) error {
	constructor := p.Target

	if _, ok := constructor.(HubOption); ok {
		return werr.Errorf("di.Option should be passed to di.New directly, "+
			"not to di.Provide: di.Provide received %v from:\n%+v",
			constructor, p.Stack)
	}

	if ann, ok := constructor.(Annotated); ok {
		var opts []dig.ProvideOption
		switch {
		case len(ann.Group) > 0 && len(ann.Name) > 0:
			return werr.Errorf(
				"di.Annotated may specify only one of Name or Group: received %v from:\n%+v",
				ann, p.Stack)
		case len(ann.Name) > 0:
			opts = append(opts, dig.Name(ann.Name))
		case len(ann.Group) > 0:
			opts = append(opts, dig.Group(ann.Group))

		}

		// 注册初始化函数
		if err := hub.Provide(ann.Target, opts...); err != nil {
			return werr.Errorf("di.Provide(%v) from:\n%+vFailed: %v", ann, p.Stack, err)
		}

		// 注册清理函数
		if ann.Close != nil {
			hub.cleanup.Register(ann.Close)
		}

		return nil
	}

	if reflect.TypeOf(constructor).Kind() == reflect.Func {
		ft := reflect.ValueOf(constructor).Type()

		for i := 0; i < ft.NumOut(); i++ {
			t := ft.Out(i)

			if t == reflect.TypeOf(Annotated{}) {
				return werr.Errorf(
					"di.Annotated should be passed to di.Provide directly, "+
						"it should not be returned by the constructor: "+
						"di.Provide received %v from:\n%+v",
					xreflect.FuncName(constructor), p.Stack)
			}
		}
	}

	if err := hub.Provide(constructor); err != nil {
		return werr.Errorf("di.Provide(%v) from:\n%+vFailed: %v", xreflect.FuncName(constructor), p.Stack, err)
	}

	return nil
}

// Invoke runs the given function after instantiating its dependencies.
//
// Any arguments that the function has are treated as its dependencies. The
// dependencies are instantiated in an unspecified order along with any
// dependencies that they might have.
//
// The function may return an error to indicate failure. The error will be
// returned to the caller as-is.
func (hub *Hub) Invoke(function interface{}, opts ...InvokeOption) error {
	return hub.container.Invoke(function, opts...)
}

// Provide teaches the container how to build values of one or more types and
// expresses their dependencies.
//
// The first argument of Provide is a function that accepts zero or more
// parameters and returns one or more results. The function may optionally
// return an error to indicate that it failed to build the value. This
// function will be treated as the constructor for all the types it returns.
// This function will be called AT MOST ONCE when a type produced by it, or a
// type that consumes this function's output, is requested via Invoke. If the
// same types are requested multiple times, the previously produced value will
// be reused.
//
// In addition to accepting constructors that accept dependencies as separate
// arguments and produce results as separate return values, Provide also
// accepts constructors that specify dependencies as dig.In structs and/or
// specify results as dig.Out structs.
func (hub *Hub) Provide(constructor interface{}, opts ...ProvideOption) error {
	return hub.container.Provide(constructor, opts...)
}

// GetProvidedSlice  return provide slice
func (hub *Hub) GetProvidedSlice() []Provided {
	return hub.provides
}
