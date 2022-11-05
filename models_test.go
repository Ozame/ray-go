package main

import (
	"fmt"
	"testing"
)

func TestNewPoint(t *testing.T) {
	p := NewPoint(1, 2, 3)

	if p.x != 1 || p.y != 2 || p.z != 3 {
		t.Errorf("the point contents did not match")
	}

	if p.w != 1 {
		t.Errorf("Expected w to be %f, but was %f", 0.0, p.w)
	}
}

func TestNewVector(t *testing.T) {
	p := NewVector(1, 2, 3)

	if p.x != 1 || p.y != 2 || p.z != 3 {
		t.Errorf("the point contents did not match")
	}

	if p.w != 0 {
		t.Errorf("Expected w to be %f, but was %f", 0.0, p.w)
	}
}

func TestAdd(t *testing.T) {

	var tests = []struct {
		a, b, want Tuple
	}{
		{NewVector(1, 2, 3), NewVector(1, 2, -5), NewVector(2, 4, -2)},
		{NewVector(1, 2, 3), NewPoint(1, 2, -5), NewPoint(2, 4, -2)},
		{NewVector(0, 0, 0), NewPoint(1, 2, -5), NewPoint(1, 2, -5)},
		{NewPoint(0, 0, 0), NewPoint(1, 2, -5), Tuple{1, 2, -5, 2}},
	}

	for _, tt := range tests {

		testname := fmt.Sprintf("%v, %v", tt.a, tt.b)
		t.Run(testname, func(t *testing.T) {
			ans := Add(tt.a, tt.b)
			if ans != tt.want {
				t.Errorf("Got %v, wanted %v", ans, tt.want)
			}
		})

	}

}

func TestSubstract(t *testing.T) {

	var tests = []struct {
		a, b, want Tuple
	}{
		{NewVector(1, 2, 3), NewVector(1, 2, -5), NewVector(0, 0, 8)},
		{NewVector(1, 2, 3), NewPoint(1, 2, -5), Tuple{0, 0, 8, -1}},
		{NewVector(0, 0, 0), NewPoint(1, 2, -5), Tuple{-1, -2, 5, -1}},
		{NewPoint(0, 0, 0), NewPoint(1, 2, -5), Tuple{-1, -2, 5, 0}},
	}

	for _, tt := range tests {

		testname := fmt.Sprintf("%v, %v", tt.a, tt.b)
		t.Run(testname, func(t *testing.T) {
			ans := Substract(tt.a, tt.b)
			if ans != tt.want {
				t.Errorf("Got %v, wanted %v", ans, tt.want)
			}
		})

	}

}
