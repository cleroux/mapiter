package basic

import "testing"

func TestBasic(t *testing.T) {

	m := make(map[string]bool)
	m["a"] = true
	m["b"] = true

	_, _ = basic(m)

	// Iterating over a map is ignored in tests by default.
	for k, v := range m { // want `iterating over a map`
		_ = k
		_ = v
	}
}
