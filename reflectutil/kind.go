package reflectutil

import (
	"reflect"
)

// IsStructPtr returns true if obj is a pointer to a struct.
func IsStructPtr(obj interface{}) bool {
	objVal := reflect.ValueOf(obj)
	return objVal.Kind() == reflect.Ptr && objVal.Elem().Kind() == reflect.Struct
}
