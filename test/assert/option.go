package assert

import "github.com/stretchr/testify/suite"

type BuildOption func(*Builder)

func WithTestifySuite(s *suite.Suite) BuildOption {
	return func(a *Builder) {
		a.Assert = s.Assertions
		a.t = s.T()
	}
}

func WithGlobalStopBuildOption() BuildOption {
	return func(a *Builder) {
		a.StopWhenFailed = true
	}
}

func NewOption(opts ...WithOption) *Option {
	option := &Option{}
	for _, opt := range opts {
		opt(option)
	}
	return option
}

type Option struct {
	StopWhenFail          bool
	SetGlobalStopWhenFail bool
	Msg                   string
	Args                  []any
	MsgAndArgs            []any
}

type WithOption func(*Option)

func WithStop() WithOption {
	return func(opt *Option) {
		opt.StopWhenFail = true
	}
}

var WMsg = WithMsgAndArg

func WithMsgAndArg(msg string, args ...any) WithOption {
	return func(opt *Option) {
		if msg != "" {
			opt.MsgAndArgs = append([]any{msg}, args...)
		}
	}
}

func WithGlobalStop() WithOption {
	return func(opt *Option) {
		opt.SetGlobalStopWhenFail = true
	}
}
