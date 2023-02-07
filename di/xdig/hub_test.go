package xdig

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"go.uber.org/dig"
)

func TestHub_New(t *testing.T) {
	t.Run("basic di", func(t *testing.T) {
		hub := New()
		err := hub.Provide(module1)
		require.NoError(t, err)

		require.Len(t, hub.HubProvideGet(), 4)
		require.Contains(t, hub.HubProvideGet().String(), "NewE")
		require.NoError(t, err)
		require.NotNil(t, hub)

		err = hub.Invoke(func(c C) {
			assert.Equal(t, c.Number, int64(100))
		})
		require.NoError(t, err)

		err = hub.Invoke(func(d D) {
			assert.Equal(t, d.C.Number, int64(100))
		})
		require.NoError(t, err)

		err = hub.Invoke(func(f1 F1) {
			assert.Equal(t, f1.Name, "f1")
		})
		require.NoError(t, err)
	})

	t.Run("name", func(t *testing.T) {
		hub := New()
		err := hub.Provide(module2)
		require.NoError(t, err)

		require.NotNil(t, hub)

		// invoke now
		err = hub.Invoke(func(k1 K1) { assert.Equal(t, k1.J.Name, "j1") })
		require.NoError(t, err)
		err = hub.Invoke(func(k2 K2) { assert.Equal(t, k2.J.Name, "j2") })
		require.NoError(t, err)

	})

	t.Run("missing deps", func(t *testing.T) {
		hub := New()
		err := hub.Provide(moduleFailed)
		require.NoError(t, err)

		err = hub.Invoke(func(g G) {})
		assert.Error(t, err)
		assert.True(t, strings.Contains(err.Error(), "missing dependencies"))
	})
	t.Run("provide failed", func(t *testing.T) {
		hub := New()
		err := hub.Provide(NewProvide(1))
		require.Error(t, err)
	})
}

var module1 = Groups(
	NewProvide(func() A {
		return A{Name: "a"}
	}),
	NewProvide(func(a A) B {
		return B{
			A:   a,
			Age: 20,
		}
	}),
	NewProvide(func(a A, b B) C {
		return C{
			A:      a,
			B:      b,
			Number: 100,
		}
	}),
	NewProvide(NewE),
)

var module2 = Groups(
	NewProvide(func() J {
		return J{Name: "j1"}
	}, dig.Name("j1"),
	),
	NewProvide(func() J {
		return J{Name: "j2"}
	}, dig.Name("j2"),
	),
)

var moduleFailed = Groups(
	NewProvide(func(g G) H {
		return H{
			G:    g,
			Name: "h",
		}
	}),
)

type A struct {
	Name string
}

type B struct {
	A   A
	Age int
}

type C struct {
	A      A
	B      B
	Number int64
}

type D struct {
	dig.In

	A A
	B B
	C C
}

func NewE() E {
	return E{
		F1: F1{
			Name: "f1",
		},
		F2: F2{
			Name: "f2",
		},
	}
}

type E struct {
	dig.Out

	F1 F1
	F2 F2
}

type F1 struct {
	Name string
}

type F2 struct {
	Name string
}

type G struct {
	Name string
}

type H struct {
	G    G
	Name string
}

type J struct {
	Name string
}

type K0 struct {
	Closed bool
}

func (k *K0) Close() {
	k.Closed = true
}

type K1 struct {
	dig.In
	J J `name:"j1"`
}

type K2 struct {
	dig.In
	J J `name:"j2"`
}
