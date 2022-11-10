package xdig

import (
	"fmt"
	"github.com/zonewave/pkgs/standutil/runtimeutil"
	"strings"

	"go.uber.org/dig"
)

// HubProvide is a wrapper of dig.ProvideOption
type HubProvide interface {
	fmt.Stringer
	Provided() []*Provided
}

// Groups is a helper function to group multiple HubProvide into a single HubProvide.
func Groups(opts ...HubProvide) HubProvide {
	return ProvidedGroup(opts)
}

// ProvidedGroup is a group of HubProvide.
type ProvidedGroup []HubProvide

// Provided return a Provided
func (og ProvidedGroup) Provided() []*Provided {
	provides := make([]*Provided, 0, len(og))
	for _, p := range og {
		provides = append(provides, p.Provided()...)
	}
	return provides
}

// String return a string
func (og ProvidedGroup) String() string {
	items := make([]string, len(og))
	for i, opt := range og {
		items[i] = fmt.Sprint(opt)
	}
	return fmt.Sprintf("di.Options(%s)", strings.Join(items, ", "))
}

// Provided is a wrapper of dig.ProvideOption
type Provided struct {
	Actor interface{}
	Opts  []dig.ProvideOption
}

// NewProvide return a new Provided
func NewProvide(actor interface{}, Opts ...dig.ProvideOption) *Provided {
	return &Provided{
		Actor: actor,
		Opts:  Opts,
	}
}

// Provided return a Provided
func (p *Provided) Provided() []*Provided {
	return []*Provided{p}
}

// String return a string
func (p *Provided) String() string {
	return fmt.Sprintf("Provide(%s)", runtimeutil.FuncName(p.Actor))
}
