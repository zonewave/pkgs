package assert

import "github.com/stretchr/testify/suite"

// BuildOption is the option for Builder
type BuildOption func(*Builder)

// WithTestifySuite set testify/suite
func WithTestifySuite(s *suite.Suite) BuildOption {
	return func(a *Builder) {
		a.Assert = s.Assertions
		a.t = s.T()
	}
}

// WithGlobalStopBuildOption set global stop when failed
func WithGlobalStopBuildOption() BuildOption {
	return func(a *Builder) {
		a.StopWhenFailed = true
	}
}

// NewOption create a new Option
func NewOption(opts ...WithOption) *Option {
	option := &Option{}
	for _, opt := range opts {
		opt(option)
	}
	return option
}

// Option is the option for assert
type Option struct {
	StopWhenFail          bool
	SetGlobalStopWhenFail bool
	Msg                   string
	Args                  []any
	MsgAndArgs            []any
}

// WithOption is the option fn for assert
type WithOption func(*Option)

// WithStop set stop when failed
func WithStop() WithOption {
	return func(opt *Option) {
		opt.StopWhenFail = true
	}
}

// WMsg is WithMsg
var WMsg = WithMsgAndArg

// WithMsgAndArg set msg
func WithMsgAndArg(msg string, args ...any) WithOption {
	return func(opt *Option) {
		if msg != "" {
			opt.MsgAndArgs = append([]any{msg}, args...)
		}
	}
}

// WithGlobalStop set global stop when failed
func WithGlobalStop() WithOption {
	return func(opt *Option) {
		opt.SetGlobalStopWhenFail = true
	}
}
