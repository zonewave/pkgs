package xdig

import (
	"fmt"
	"github.com/zonewave/pkgs/util/runtimeutil"
	"strings"

	"go.uber.org/dig"
)

type HubProvide interface {
	fmt.Stringer
	Provided() []*Provided
}

func Groups(opts ...HubProvide) HubProvide {
	return ProvidedGroup(opts)
}

type ProvidedGroup []HubProvide

func (og ProvidedGroup) Provided() []*Provided {
	provides := make([]*Provided, 0, len(og))
	for _, p := range og {
		provides = append(provides, p.Provided()...)
	}
	return provides
}

func (og ProvidedGroup) String() string {
	items := make([]string, len(og))
	for i, opt := range og {
		items[i] = fmt.Sprint(opt)
	}
	return fmt.Sprintf("di.Options(%s)", strings.Join(items, ", "))
}

type Provided struct {
	Actor interface{}
	Opts  []dig.ProvideOption
}

func NewProvide(actor interface{}, Opts ...dig.ProvideOption) *Provided {
	return &Provided{
		Actor: actor,
		Opts:  Opts,
	}
}

func (p *Provided) Provided() []*Provided {
	return []*Provided{p}
}

func (o *Provided) String() string {
	return fmt.Sprintf("Provide(%s)", runtimeutil.FuncName(o.Actor))
}
