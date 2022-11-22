package main

import "testing"

func TestEquals(t *testing.T) {

	a, b := 1.002, 1.002
	if !equals(a, b) {
		t.Errorf("a and b were not equal")
	}

	c, d := 1.002, 1.002000001
	if !equals(c, d) {
		t.Errorf("c and d were not equal")
	}

	e, f := 33.0, 45.2
	if equals(e, f) {
		t.Errorf("e and f were the same")
	}
}
