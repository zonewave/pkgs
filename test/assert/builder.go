package assert

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type Builder struct {
	ok             bool
	Assert         *assert.Assertions
	t              *testing.T
	StopWhenFailed bool
}

func NewBuilderWithSuite(s *suite.Suite, opts ...BuildOption) *Builder {
	opts = append([]BuildOption{WithTestifySuite(s)}, opts...)
	return NewBuilder(opts...)
}
func NewBuilder(opts ...BuildOption) *Builder {
	ret := &Builder{ok: true}
	for _, opt := range opts {
		opt(ret)
	}
	return ret
}

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

func (b *Builder) Equal(expect, actual any, opts ...WithOption) *Builder {
	return b.Apply(func(opt *Option) bool {
		return b.Assert.Equal(expect, actual, opt.MsgAndArgs...)
	}, opts...)
}

func (b *Builder) Ok() bool {
	return b.ok
}
func (b *Builder) True(actual bool, opts ...WithOption) *Builder {
	return b.Apply(func(opt *Option) bool {
		return b.Assert.True(actual, opt.MsgAndArgs...)
	}, opts...)
}
func (b *Builder) False(actual bool, opts ...WithOption) *Builder {
	return b.Apply(func(opt *Option) bool {
		return b.Assert.False(actual, opt.MsgAndArgs...)
	}, opts...)
}
func (b *Builder) Len(a any, length int, opts ...WithOption) *Builder {
	return b.Apply(func(opt *Option) bool {
		return b.Assert.Len(a, length, opt.MsgAndArgs...)
	}, opts...)
}

type Sizer interface {
	Size() int
}

func (b *Builder) Sizer(a Sizer, length int, opts ...WithOption) *Builder {
	return b.Equal(a.Size(), length, opts...)
}

func (b *Builder) NoError(err error, opts ...WithOption) *Builder {
	return b.Apply(func(option *Option) bool {
		return b.Assert.NoError(err, option.MsgAndArgs)
	}, opts...)
}
func (b *Builder) NotEmpty(obj any, opts ...WithOption) *Builder {
	return b.Apply(func(option *Option) bool {
		return b.Assert.NotEmpty(obj, option.MsgAndArgs)
	}, opts...)
}

func (b *Builder) Contains(f interface{}, contains interface{}, opts ...WithOption) *Builder {
	return b.Apply(func(option *Option) bool {
		return b.Assert.Contains(f, contains, option.MsgAndArgs)
	}, opts...)
}
func (b *Builder) Cond(fn func() bool, opts ...WithOption) *Builder {
	return b.Apply(func(option *Option) bool {
		return b.Assert.True(fn(), option.MsgAndArgs)
	}, opts...)
}

func (b *Builder) AssertResult(opts ...WithOption) {
	b.True(b.ok, opts...)
	return
}
