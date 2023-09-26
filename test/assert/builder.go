package assert

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

// Builder is a wrapper of testify/assert
type Builder struct {
	ok             bool
	Assert         *assert.Assertions
	t              *testing.T
	StopWhenFailed bool
}

// NewBuilderWithSuite  create a new Builder with testify/suite
func NewBuilderWithSuite(s *suite.Suite, opts ...BuildOption) *Builder {
	opts = append([]BuildOption{WithTestifySuite(s)}, opts...)
	return NewBuilder(opts...)
}

// NewBuilder create a new Builder
func NewBuilder(opts ...BuildOption) *Builder {
	ret := &Builder{ok: true}
	for _, opt := range opts {
		opt(ret)
	}
	return ret
}

// Apply apply a assert function
func (b *Builder) Apply(assert func(option *Option) bool, opts ...WithOption) *Builder {
	option := NewOption(opts...)
	if option.SetGlobalStopWhenFail {
		b.StopWhenFailed = true
	}

	b.ok = assert(option)
	if !b.ok {
		if b.StopWhenFailed || option.StopWhenFail {
			b.t.FailNow()
		}
	}
	return b
}

// Equal assert expect and actual is equal
func (b *Builder) Equal(expect, actual any, opts ...WithOption) *Builder {
	return b.Apply(func(opt *Option) bool {
		return b.Assert.Equal(expect, actual, opt.MsgAndArgs...)
	}, opts...)
}

// Ok return the result of assert
func (b *Builder) Ok() bool {
	return b.ok
}

// True assert actual is true
func (b *Builder) True(actual bool, opts ...WithOption) *Builder {
	return b.Apply(func(opt *Option) bool {
		return b.Assert.True(actual, opt.MsgAndArgs...)
	}, opts...)
}

// False assert actual is false
func (b *Builder) False(actual bool, opts ...WithOption) *Builder {
	return b.Apply(func(opt *Option) bool {
		return b.Assert.False(actual, opt.MsgAndArgs...)
	}, opts...)
}

// Len assert length of a is equal to length
func (b *Builder) Len(a any, length int, opts ...WithOption) *Builder {
	return b.Apply(func(opt *Option) bool {
		return b.Assert.Len(a, length, opt.MsgAndArgs...)
	}, opts...)
}

// Sizer is a interface for size
type Sizer interface {
	Size() int
}

// Sizer assert length of a is equal to length
func (b *Builder) Sizer(a Sizer, length int, opts ...WithOption) *Builder {
	return b.Equal(a.Size(), length, opts...)
}

// NoError assert err is nil
func (b *Builder) NoError(err error, opts ...WithOption) *Builder {
	return b.Apply(func(option *Option) bool {
		return b.Assert.NoError(err, option.MsgAndArgs)
	}, opts...)
}

// NotEmpty assert obj is not empty
func (b *Builder) NotEmpty(obj any, opts ...WithOption) *Builder {
	return b.Apply(func(option *Option) bool {
		return b.Assert.NotEmpty(obj, option.MsgAndArgs)
	}, opts...)
}

// Contains assert f contains contains
func (b *Builder) Contains(f interface{}, contains interface{}, opts ...WithOption) *Builder {
	return b.Apply(func(option *Option) bool {
		return b.Assert.Contains(f, contains, option.MsgAndArgs)
	}, opts...)
}

// Cond assert fn is true
func (b *Builder) Cond(fn func() bool, opts ...WithOption) *Builder {
	return b.Apply(func(option *Option) bool {
		return b.Assert.True(fn(), option.MsgAndArgs)
	}, opts...)
}

// AssertResult assert the result
func (b *Builder) AssertResult(opts ...WithOption) {
	b.True(b.ok, opts...)
	return
}
