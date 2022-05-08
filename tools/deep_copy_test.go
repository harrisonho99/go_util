package tools

import (
	"testing"

	"github.com/hotsnow199/go_util/util"
)

func TestCopy(t *testing.T) {
	// copy struct
	var p = P{A: 1, B: 2, C: 3, D: "d", E: []int{1, 2, 3}}
	var q Q

	util.CheckTestError(t, DeepCopy(p, &q))

	q.D = "mutated"
	q.E[0] = 10
	q.E[1] = 20
	q.E[2] = 30

	for i := range p.E {
		if p.E[i] == q.E[i] {
			t.Error("struct not copied correctly")
		}
	}
	t.Logf("p : %#v\n", p)
	t.Logf("q : %#v\n", q)

	// copy map
	var a map[string]int = map[string]int{"a": 1, "b": 2, "c": 3}
	var b map[string]int
	DeepCopy(a, &b)
	b["a"] = 10
	b["b"] = 20
	b["c"] = 30
	t.Logf("a : %#v\n", a)
	t.Logf("b : %#v\n", b)

	for key := range a {
		if a[key] == b[key] {
			t.Error("map not copied correctly")
		}
	}

	// copy slice
	var d, e [](map[string]int)
	d = append(d, map[string]int{"a": 1, "b": 2, "c": 3})
	util.CheckTestError(t, DeepCopy(d, &e))
	e[0]["a"] = 10
	e[0]["b"] = 20
	e[0]["c"] = 30

	for i := range d {
		for key := range d[i] {
			if d[i][key] == e[i][key] {
				t.Error("slice not copied correctly")
			}
		}
	}

	t.Logf("d : %#v\n", d)
	t.Logf("e : %#v\n", e)

}
