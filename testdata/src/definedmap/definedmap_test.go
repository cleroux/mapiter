package basic

import "testing"

func TestBasic(t *testing.T) {

	m := make(mapStringBool)
	m["a"] = true
	m["b"] = true

	_, _ = basic(m)

	// Iterating over a map analysis is included in tests by default.
	for k, v := range m { // want `iterating over a map`
		_ = k
		_ = v
	}
}
