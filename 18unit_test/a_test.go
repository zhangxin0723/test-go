package main

import "testing"

func TestAdd(t *testing.T) {
	res := add(1, 2)
	if res != 3 {
		t.Errorf("add(1, 2) = %d; want 3", res)
	}
}
