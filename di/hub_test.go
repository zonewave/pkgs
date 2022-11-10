package di

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"go.uber.org/dig"
)

func TestHub_New(t *testing.T) {
	t.Run("basic di", func(t *testing.T) {
		hub, err := New(module1, ValidateHub(true))
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

	t.Run("annotation", func(t *testing.T) {
		k0 := &K0{}
		hub, err := New(module2, Annotation(func(ann *Annotated) {
			ann.Target = func() *K0 { return k0 }
			ann.Close = k0.Close
		}))
		require.NoError(t, err)
		require.NotNil(t, hub)

		// invoke now
		err = hub.Invoke(func(k1 K1) { assert.Equal(t, k1.J.Name, "j1") })
		require.NoError(t, err)
		err = hub.Invoke(func(k2 K2) { assert.Equal(t, k2.J.Name, "j2") })
		require.NoError(t, err)

		// cleanup and check it
		hub.Cleanup()
		require.True(t, k0.Closed)
	})

	t.Run("missing deps", func(t *testing.T) {
		hub, err := New(moduleFailed, ValidateHub(false))
		require.NoError(t, err) // won't check missing H in G, only check it after invoke
		require.NotNil(t, hub)

		err = hub.Invoke(func(g G) {})
		assert.Error(t, err)
		assert.True(t, strings.Contains(err.Error(), "missing dependencies"))
	})
}

var module1 = Options(
	Provide(func() A {
		return A{Name: "a"}
	}),
	Provide(func(a A) B {
		return B{
			A:   a,
			Age: 20,
		}
	}),
	Provide(func(a A, b B) C {
		return C{
			A:      a,
			B:      b,
			Number: 100,
		}
	}),
	Provide(NewE),
)

var module2 = Options(
	Provide(Annotated{
		Name: "j1",
		Target: func() J {
			return J{Name: "j1"}
		},
	}),
	Provide(Annotated{
		Name: "j2",
		Target: func() J {
			return J{Name: "j2"}
		},
	}),
)

var moduleFailed = Options(
	Provide(func(g G) H {
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
