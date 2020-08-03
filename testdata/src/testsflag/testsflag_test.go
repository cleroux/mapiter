package testsflag

import "testing"

func TestTestsFlag(t *testing.T) {

	m := make(map[string]bool)
	m["a"] = true
	m["b"] = true

	_, _ = testsFlag(m)

	for k, v := range m { // no error when tests flag is false
		_ = k
		_ = v
	}
}
