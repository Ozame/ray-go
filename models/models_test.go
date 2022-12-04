package models

import (
	"fmt"
	"math"
	"strings"
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

func TestCanvasWidthAndHeight(t *testing.T) {
	c := NewCanvas(2, 5)
	w, h := c.GetWidthAndHeight()
	if w != 2 || h != 5 {
		t.Errorf("Got %d, %d, %d, %d", w, h, 2, 5)
	}
}

func TestWriteAndGetPixel(t *testing.T) {
	canvas := NewCanvas(4, 3)
	red := NewColor(1, 0, 0)
	canvas.WritePixel(3, 2, red)
	coloredPixel := canvas.Get(3, 2)
	if coloredPixel != red {
		t.Errorf("Got %v, expected %v", coloredPixel, red)
	}

}

func TestCanvasToPPM(t *testing.T) {
	canvas := NewCanvas(5, 3)
	c1 := NewColor(1.5, 0, 0)
	c2 := NewColor(0, 0.5, 0)
	c3 := NewColor(-0.5, 0, 1)
	canvas.WritePixel(0, 0, c1)
	canvas.WritePixel(2, 1, c2)
	canvas.WritePixel(4, 2, c3)

	ppm := canvas.ToPPM()
	lines := strings.Split(ppm, "\n")

	if lines[0] != "P3" || lines[1] != "5 3" || lines[2] != "255" {
		t.Errorf("Header incorrect, got %v", lines)
	}

	expectedLine4 := "255 0 0 0 0 0 0 0 0 0 0 0 0 0 0"
	expectedLine5 := "0 0 0 0 0 0 0 128 0 0 0 0 0 0 0"
	expectedLine6 := "0 0 0 0 0 0 0 0 0 0 0 0 0 0 255"
	if lines[3] != expectedLine4 || lines[4] != expectedLine5 || lines[5] != expectedLine6 {
		t.Errorf("Content incorrect, got \n%v", ppm)
	}

}

func TestPpmLineLength(t *testing.T) {
	c := NewCanvas(10, 2)
	color := NewColor(1, 0.8, 0.6)
	c.SetEveryPixel(color)

	ppm := c.ToPPM()
	lines := strings.Split(ppm, "\n")

	expectedLine4 := "255 204 153 255 204 153 255 204 153 255 204 153 255 204 153 255 204"
	expectedLine5 := "153 255 204 153 255 204 153 255 204 153 255 204 153"
	expectedLine6 := "255 204 153 255 204 153 255 204 153 255 204 153 255 204 153 255 204"
	expectedLine7 := "153 255 204 153 255 204 153 255 204 153 255 204 153"
	if lines[3] != expectedLine4 {
		t.Errorf("Line incorrect, got \n%v, expected \n%v", lines[3], expectedLine4)
	}

	if lines[4] != expectedLine5 {
		t.Errorf("Line incorrect, got \n%v, expected \n%v", lines[4], expectedLine5)
	}

	if lines[5] != expectedLine6 {
		t.Errorf("Line incorrect, got \n%v, expected \n%v", lines[5], expectedLine6)
	}

	if lines[6] != expectedLine7 {
		t.Errorf("Line incorrect, got \n%v, expected \n%v", lines[6], expectedLine7)
	}

	if lines[7] != "" {
		t.Errorf("File should end in a newline, got %v", lines[6])
	}

}
