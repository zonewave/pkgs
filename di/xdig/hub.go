package xdig

import (
	"github.com/cockroachdb/errors"
	"go.uber.org/dig"
)

// New return a New instance
func New(opts ...dig.Option) *Hub {
	hub := &Hub{
		container: dig.New(opts...),
	}

	return hub
}

// Hub is a wrapper of dig
type Hub struct {
	container *dig.Container
	provides  []*Provided
}

// Invoke invokes a function, supplying arguments from the container.
func (hub *Hub) Invoke(function interface{}, opts ...dig.InvokeOption) error {
	return hub.container.Invoke(function, opts...)
}

// Provide adds a constructor to the container.
func (hub *Hub) Provide(hubProvides ...HubProvide) error {
	provides := Groups(hubProvides...).Provided()
	for _, p := range provides {
		if err := hub.container.Provide(p.Actor, p.Opts...); err != nil {
			return errors.WithStack(err)
		}
		hub.provides = append(hub.provides, p)
	}
	return nil
}

// HubProvideGet return all provides
func (hub *Hub) HubProvideGet() HubProvide {
	hubProvide := make([]HubProvide, 0, len(hub.provides))
	for _, p := range hub.provides {
		hubProvide = append(hubProvide, p)

	}
	return Groups(hubProvide...)
}
