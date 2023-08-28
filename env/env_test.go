//go:build unit

package env

import "testing"

func TestGetenv(t *testing.T) {
	tc := []struct {
		description string
		key         string
		val         string
		fallback    string
		exp         string
	}{
		{
			description: "existing env var should return var",
			key: "0",
			val: "zero",
			fallback: "",
			exp: "zero",
		},
		{
			description: "non-existent env var should return fallback",
			key: "",
			val: "",
			fallback: "one",
			exp: "one",
		},
	}

	for i, c := range tc {
		if c.key != "" {
			t.Setenv(c.key, c.val)
		}
		
		v := Getenv(c.key, c.fallback)
		if c.val != "" && c.val != v {
			t.Errorf("[%d] %s: expected val %s; received %s", i, c.description, c.val, v)
		}
		if c.fallback != "" && c.fallback != v {
			t.Errorf("[%d] %s: expected fallback %s; received %s", i, c.description, c.fallback, v)	
		}
	}
}
