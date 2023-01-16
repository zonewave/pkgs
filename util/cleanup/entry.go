// Package cleanup provides a single point for registering clean up functions.
// This is similar to Finalizer, except that the clean up functions are
// guaranteed to be called if the process terminates normally.
//
// Usage:
//
// In my_package.go
//
//	cleanup.Register(func(){
//	  // Arbitrary clean up function, most likely close goroutine, etc.
//	})
//
// In main.go
//
//	func main() {
//	  flag.Parse()
//	  defer cleanup.Run()
//	}
package cleanup

import (
	"github.com/zonewave/pkgs/log"
	"reflect"
	"sync"

	"github.com/cockroachdb/errors"
)

// entry global cleanup entry
var entry Entry

// Register adds a function to the global cleanup queue.
func Register(container interface{}) {
	entry.Register(container)
}

// Run runs all the global cleanup functions registered.
func Run() {
	entry.Run()
}

// Entry cleanup Entry
type Entry struct {
	mu   sync.Mutex
	fns  []func()
	once sync.Once
}

// Run runs all the cleanup functions registered.
func (entry *Entry) Run() {
	log.Infof("cleanup: performing %d cleanups", len(entry.fns))

	entry.once.Do(func() {
		for _, f := range entry.fns {
			f()
		}
	})

	log.Infof("cleanup: all done")
}

// Register adds a function to the cleanup queue.
func (entry *Entry) Register(container interface{}) {
	v := reflect.Indirect(reflect.ValueOf(container))
	var err error
	switch v.Kind() {
	case reflect.Func:
		err = entry.RegisterFunc(v.Interface())
	case reflect.Struct:
		err = entry.RegisterStruct(v.Interface())
	default:
		panic("cleanup: unsupported type")
	}
	if err != nil {
		panic(err)
	}
}

// RegisterStruct add registered interface
func (entry *Entry) RegisterStruct(ctor interface{}) error {
	cValue := reflect.Indirect(reflect.ValueOf(ctor))
	if cValue.Kind() != reflect.Struct {
		return errors.New("RegisterStruct receive a struct or ptr to struct")
	}
	for i := 0; i < cValue.NumField(); i++ {
		field := cValue.Field(i)
		if field.IsZero() {
			continue
		}
		method := field.MethodByName("Close")
		if !method.IsValid() {
			method = field.MethodByName("Flush")
		}
		if method.IsValid() {
			if err := entry.RegisterFunc(method.Interface()); err != nil {
				log.WithError(err).WithField("fieldName", field.Type().Name()).Error("register func failed")
			}
		}
	}
	return nil
}

// RegisterFunc receive func() or func() error
func (entry *Entry) RegisterFunc(fn interface{}) error {
	fType := reflect.TypeOf(fn)
	if fType.Kind() != reflect.Func {
		return errors.New("cleanup: unsupported type")
	}
	if fType.NumIn() > 0 {
		return errors.New("RegisterFunc receive func() or func() error")
	}

	if f, ok := fn.(func()); ok {
		entry.register(f)
		return nil
	}
	if f, ok := fn.(func() error); ok {
		entry.register(func() {
			if err := f(); err != nil {
				log.WithError(err).Error("register func() failed")
			}
		})
		return nil
	}
	return errors.New("RegisterFunc receive func() or func() error")
}

func (entry *Entry) register(fns ...func()) {
	entry.mu.Lock()
	defer entry.mu.Unlock()
	for _, fn := range fns {
		entry.fns = append(entry.fns, fn)
	}
}
