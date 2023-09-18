package runtimeutil

import (
	"fmt"
	"reflect"
	"runtime"
	"strings"
)

// CallerFuncPos get caller's func name
func CallerFuncPos(skip int) string {
	_, line, name := Caller(skip + 1)
	return fmt.Sprintf("%s:%d", name, line)
}

// CallerFuncName get caller's func name
func CallerFuncName(skip int) string {
	_, _, name := Caller(skip + 1)
	arr := strings.Split(name, "/")
	start := 0
	if len(arr) >= 3 {
		start = 2
	}
	return strings.Join(arr[start:], ".")
}

// Caller file, file line, function name
func Caller(skip int) (file string, line int, functionName string) {
	var (
		pc uintptr
		ok bool
	)
	pc, file, line, ok = runtime.Caller(skip + 1)
	if !ok {
		return
	}

	return file, line, runtime.FuncForPC(pc).Name()
}

// FuncName returns a funcs formatted name
func FuncName(fn interface{}) string {
	fnV := reflect.ValueOf(fn)
	if fnV.Kind() != reflect.Func {
		return fmt.Sprint(fn)
	}

	return runtime.FuncForPC(fnV.Pointer()).Name()
}
