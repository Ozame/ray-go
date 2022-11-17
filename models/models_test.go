package models

import (
	"fmt"
	"math"
	"testing"
)

func TestNewPoint(t *testing.T) {
	p := NewPoint(1, 2, 3)

	if p.X != 1 || p.Y != 2 || p.Z != 3 {
		t.Errorf("the point contents did not match")
	}

	if p.W != 1 {
		t.Errorf("Expected w to be %f, but was %f", 0.0, p.W)
	}
}

func TestNewVector(t *testing.T) {
	p := NewVector(1, 2, 3)

	if p.X != 1 || p.Y != 2 || p.Z != 3 {
		t.Errorf("the point contents did not match")
	}

	if p.W != 0 {
		t.Errorf("Expected w to be %f, but was %f", 0.0, p.W)
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

func TestNegate(t *testing.T) {
	a := NewPoint(4, 20, 3)
	res := Negate(a)
	if res.X != -a.X || res.Y != -a.Y || res.Z != -a.Z || res.W != -a.W {
		t.Errorf("Unexpected result, got %v", res)
	}
}

func TestMultiply(t *testing.T) {
	a := Tuple{1, -2, 3, -4}

	scalarResult := Multiply(a, 3.5)
	scalarExp := Tuple{3.5, -7, 10.5, -14}
	if scalarResult != scalarExp {
		t.Errorf("Expected %v, got %v", scalarExp, scalarResult)
	}

	fractionResult := Multiply(a, 0.5)
	fractionExp := Tuple{0.5, -1, 1.5, -2}

	if fractionResult != fractionExp {
		t.Errorf("Expected %v, got %v", fractionExp, fractionResult)
	}
}

func TestMagnitude(t *testing.T) {
	var tests = []struct {
		tuple Tuple
		want  float64
	}{
		{NewVector(1, 0, 0), 1},
		{NewVector(0, 1, 0), 1},
		{NewVector(0, 0, 1), 1},
		{NewVector(1, 2, 3), math.Sqrt(14)},
		{NewVector(-1, -2, -3), math.Sqrt(14)},
	}

	for _, tt := range tests {

		testname := fmt.Sprintf("%v", tt)
		t.Run(testname, func(t *testing.T) {
			ans := Magnitude(tt.tuple)
			if ans != tt.want {
				t.Errorf("Got %v, wanted %v", ans, tt.want)
			}
		})

	}

}

func TestNormalize(t *testing.T) {
	var tests = []struct {
		tuple, want Tuple
	}{
		{NewVector(4, 0, 0), NewVector(1, 0, 0)},
		{NewVector(1, 2, 3), NewVector(1/math.Sqrt(14), 2/math.Sqrt(14), 3/math.Sqrt(14))},
	}

	for _, tt := range tests {

		testname := fmt.Sprintf("%v", tt)
		t.Run(testname, func(t *testing.T) {
			ans := Normalize(tt.tuple)
			if ans != tt.want {
				t.Errorf("Got %v, wanted %v", ans, tt.want)
			}
		})

	}

}

func TestDot(t *testing.T) {
	a, b := NewVector(1, 2, 3), NewVector(2, 3, 4)
	res := Dot(a, b)

	if res != 20 {
		t.Errorf("Expected 20, got %f", res)
	}

}

func TestCross(t *testing.T) {
	a, b := NewVector(1, 2, 3), NewVector(2, 3, 4)

	res1 := Cross(a, b)
	res2 := Cross(b, a)
	exp1 := NewVector(-1, 2, -1)
	exp2 := NewVector(1, -2, 1)

	if res1 != exp1 {
		t.Errorf("Got %v, expected %v", res1, exp1)
	}
	if res2 != exp2 {
		t.Errorf("Got %v, expected %v", res2, exp2)
	}

}

func TestNewColor(t *testing.T) {
	r, g, b := -0.5, 0.4, 1.7
	c := NewColor(r, g, b)

	if c.red != r || c.green != g || c.blue != b {
		t.Errorf("Got %v, expected %f, %f, %f", c, r, g, b)
	}
}

func TestColorEquals(t *testing.T) {
	var tests = []struct {
		c1, c2 Color
		want   bool
	}{
		{NewColor(1, 2, 3), NewColor(1, 2, 3), true},
		{NewColor(1, 2, 3), NewColor(1, 2, 5), false},
		{NewColor(1, 2, 3), NewColor(1, 1, 3), false},
		{NewColor(1, 2, 3), NewColor(3, 2, 3), false},
		{NewColor(1, 2, 3), NewColor(1, 2, -3), false},
	}

	for _, tt := range tests {

		testname := fmt.Sprintf("%v", tt)
		t.Run(testname, func(t *testing.T) {
			ans := tt.c1.Equals(tt.c2)
			if ans != tt.want {
				t.Errorf("Got %v, wanted %v", ans, tt.want)
			}
		})

	}

}

func TestAddColors(t *testing.T) {
	c := NewColor(0.9, 0.6, 0.75)
	c2 := NewColor(0.7, 0.1, 0.25)

	want := NewColor(1.6, 0.7, 1.0)
	res := c.Add(c2)

	if !res.Equals(want) {
		t.Errorf("Got %v, expected %v", res, want)
	}
}

func TestSubColors(t *testing.T) {
	c := NewColor(0.9, 0.6, 0.75)
	c2 := NewColor(0.7, 0.1, 0.25)

	want := NewColor(0.2, 0.5, 0.5)
	res := c.Sub(c2)

	if !res.Equals(want) {
		t.Errorf("Got %v, expected %v", res, want)
	}
}

func TestColorTimesScalar(t *testing.T) {
	c := NewColor(0.9, 0.6, 0.80)

	want := NewColor(2.25, 1.5, 2.0)
	res := c.Times(2.5)

	if !res.Equals(want) {
		t.Errorf("Got %v, expected %v", res, want)
	}
}

func TestBlendColors(t *testing.T) {
	c := NewColor(1, 2, 3)
	c2 := NewColor(3, 2, 1)

	want := NewColor(3, 4, 3)
	res := c.Blend(c2)

	if !res.Equals(want) {
		t.Errorf("Got %v, expected %v", res, want)
	}
}

func TestNewCanvas(t *testing.T) {
	c := NewCanvas(10, 20)
	defaultColor := NewColor(0, 0, 0)

	for i := 0; i < 20; i++ {
		for j := 0; j < 10; j++ {
			if !c[i][j].Equals(defaultColor) {
				t.Errorf("Got %v, expected %v", c[i][j], defaultColor)
			}
		}
	}
}
