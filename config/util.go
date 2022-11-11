package config

import (
	"github.com/spf13/afero"
	"os"
	"reflect"
)

// Check if File / Directory Exists
func exists(fs afero.Fs, path string) (bool, error) {
	_, err := fs.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, err
}

func checkObject(obj interface{}) error {
	objVal := reflect.ValueOf(obj)
	if objVal.Kind() == reflect.Ptr && objVal.Elem().Kind() == reflect.Struct {
		return nil
	}

	return InvalidConfigTypeError("should be a pointer to a struct")
}

// copyObject return inner type copy
func copyObject(ptr interface{}) interface{} {
	return reflect.New(reflect.Indirect(reflect.ValueOf(ptr)).Type()).Interface()
}
