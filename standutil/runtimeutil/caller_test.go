package runtimeutil

import (
	"github.com/stretchr/testify/require"
	"strings"
	"testing"
)

func TestCallerFuncPos(t *testing.T) {
	require.True(t, strings.Contains(moo(), "TestCallerFuncPos:"))
}

func TestCaller(t *testing.T) {
	require.True(t, strings.Contains(foo(1), "TestCaller"))
	require.True(t, strings.Contains(bar(2), "TestCaller"))
	require.Empty(t, foo(100))

	file, line, funcName := Caller(0)
	require.True(t, strings.Contains(file, "caller_test.go"))
	require.True(t, strings.Contains(funcName, ".TestCaller"))
	require.True(t, line != 0)
}

func TestCallerFuncName(t *testing.T) {
	require.True(t, strings.Contains(baz(), "TestCallerFuncName"))
}

func TestFuncName(t *testing.T) {
	require.True(t, strings.Contains(FuncName(bar), "runtimeutil.bar"))
	require.True(t, strings.Contains(FuncName(1), "1"))
}

func foo(skip int) string {
	_, _, s := Caller(skip)
	return s
}

func bar(skip int) string {
	return foo(skip)
}

func moo() string {
	return CallerFuncPos(1)
}

func baz() string {
	return CallerFuncName(1)
}
