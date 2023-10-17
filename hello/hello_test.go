//go:build unit

package hello

import (
	"testing"
)

func TestHello(t *testing.T) {
	tc := []struct {
		desc string
		exp  string
	}{
		{
			desc: "hello() should return 'Hello, World!'",
			exp:  "Hello, World!",
		},
	}

	for i, c := range tc {
		s := Hello()
		if s != c.exp {
			t.Errorf("[%d] %s: expected %s, received %s", i, c.desc, c.exp, s)
		}
	}
}
