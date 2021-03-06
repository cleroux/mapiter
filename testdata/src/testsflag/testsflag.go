package testsflag

import "fmt"

type kv struct {
	k string
	v bool
}

func testsFlag(m map[string]bool) (s []kv, err error) {

	s = make([]kv, 0)

	for k, v := range m { // want `iterating over a map`
		kv := kv{k: k, v: v}
		s = append(s, kv)
	}

	for i, v := range s { // no error
		fmt.Printf("%d: %s", i, v.k)
	}

	return s, nil
}
