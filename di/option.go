package di

import (
	"fmt"
	"strings"

	"github.com/ybzhanghx/pkgs/di/xreflect"
	"go.uber.org/dig"
)

// Option configures a Hub. It's included for future functionality;
// currently, there are no concrete implementations.
type Option = dig.Option

// A ProvideOption modifies the default behavior of Provide.
type ProvideOption = dig.ProvideOption

// An InvokeOption modifies the default behavior of Invoke. It's included for
// future functionality; currently, there are no concrete implementations.
type InvokeOption = dig.InvokeOption

// HubOption is an option configures an Hub using the functional options paradigm
// popularized by Rob Pike. If you're unfamiliar with this style, see
// https://commandcenter.blogspot.com/2014/01/self-referential-functions-and-design.html.
type HubOption interface {
	fmt.Stringer
	Apply(hub *Hub)
}

// Options converts a collection of Options into a single Option. This allows
// packages to bundle sophisticated functionality into easy-to-use Fx modules.
// For example, a logging package might export a simple option like this:
//
//  package logging
//
//  var Module = fx.Provide(func() *log.Logger {
//    return log.New(os.Stdout, "", 0)
//  })
//
// A shared all-in-one microservice package could then use Options to bundle
// logging with similar metrics, tracing, and gRPC modules:
//
//  package server
//
//  var Module = fx.Options(
//    logging.Module,
//    metrics.Module,
//    tracing.Module,
//    grpc.Module,
//  )
//
// Since this all-in-one module1 has a minimal API surface, it's easy to add
// new functionality to it without breaking existing users. Individual
// applications can take advantage of all this functionality with only one
// line of code:
//
//  app := di.New(server.Module)
//
// Use this pattern sparingly, since it limits the user's ability to customize
// their application.
func Options(opts ...HubOption) HubOption {
	return optionGroup(opts)
}

type optionGroup []HubOption

func (og optionGroup) Apply(hub *Hub) {
	for _, opt := range og {
		opt.Apply(hub)
	}
}

func (og optionGroup) String() string {
	items := make([]string, len(og))
	for i, opt := range og {
		items[i] = fmt.Sprint(opt)
	}
	return fmt.Sprintf("di.Options(%s)", strings.Join(items, ", "))
}

// Annotation define annotated and apply to Provided
// Annotation provides instantiated values for dependency injection as if
// they had been provided using a constructor that simply returns them.
// The most specific type of each value (as determined by reflection) is used.
//
//	type K0 struct {
//		Closed bool
//	}
//
//	func (k *K0) Close() {
//		k.Closed = true
//	}
//
//	k0 := &K0{}
//	var module = Annotation(func(ann *Annotated) {
//		ann.Target = func() *K0 { return k0 }
//		ann.Close = k0.Close
//	})
//
//	hub, err := di.New(module)
func Annotation(f func(ann *Annotated)) HubOption {
	annotation := &Annotated{}
	f(annotation)

	return provideOption{
		Targets: []interface{}{*annotation},
		Stack:   xreflect.CallerStack(1, 0),
	}
}

// Provide registers any number of constructor functions, teaching the
// application how to instantiate various types. The supplied constructor
// function(s) may depend on other types available in the application, must
// return one or more objects, and may return an error. For example:
//
//  // Constructs type *C, depends on *A and *B.
//  func(*A, *B) *C
//
//  // Constructs type *C, depends on *A and *B, and indicates failure by
//  // returning an error.
//  func(*A, *B) (*C, error)
//
//  // Constructs types *B and *C, depends on *A, and can fail.
//  func(*A) (*B, *C, error)
func Provide(constructors ...interface{}) HubOption {
	// check no nil or error
	for _, value := range constructors {
		switch value.(type) {
		case nil:
			panic("untyped nil passed to di.Provide")
		case error:
			panic("error value passed to di.Provide")
		}
	}

	return provideOption{
		Targets: constructors,
		Stack:   xreflect.CallerStack(1, 0),
	}
}

type provideOption struct {
	Targets []interface{}
	Stack   xreflect.Stack
}

func (o provideOption) Apply(hub *Hub) {
	for _, target := range o.Targets {
		hub.provides = append(hub.provides, Provided{
			Target: target,
			Stack:  o.Stack,
		})
	}
}

func (o provideOption) String() string {
	items := make([]string, len(o.Targets))
	for i, c := range o.Targets {
		items[i] = xreflect.FuncName(c)
	}
	return fmt.Sprintf("fx.Provide(%s)", strings.Join(items, ", "))
}

// ValidateHub validates that supplied graph would run and is not missing any dependencies.
func ValidateHub(v bool) HubOption {
	return validateOption{validate: v}
}

type validateOption struct {
	validate bool
}

func (o validateOption) Apply(hub *Hub) {
	hub.validate = o.validate
}

func (o validateOption) String() string {
	return fmt.Sprintf("fx.validate(%v)", o.validate)
}
