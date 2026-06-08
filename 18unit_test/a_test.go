package main

import "testing"

func TestAdd(t *testing.T) {

	t.Run("add1", func(t *testing.T) {
		res := add(1, 2)
		if res != 3 {
			t.Errorf("add(1, 2) = %d; want 3", res)
		}
	})

	t.Run("add2", func(t *testing.T) {
		res := add(2, 3)
		if res != 5 {
			t.Errorf("add(2, 3) = %d; want 5", res)
		}
	})

	cases := []struct {
		name      string
		a, b, res int
	}{
		{name: "add3", a: 3, b: 4, res: 7},
		{name: "add4", a: 4, b: 5, res: 9},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			res := add(c.a, c.b)
			if res != c.res {
				t.Errorf("add(%d, %d) = %d; want %d", c.a, c.b, res, c.res)
			}
		})
	}
}
