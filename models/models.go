package models

import (
	"fmt"
	"math"
	"strings"
)

// ======== Types =============

type Tuple struct {
	X, Y, Z, W float64
}

type Color struct {
	red, green, blue float64
}

type Canvas [][]Color

// ======= Utils ==============

func equals(a, b float64) bool {
	eps := 0.00001
	return math.Abs(a-b) < eps
}

// ============================

func NewCanvas(w, h int) Canvas {
	var canvas [][]Color = make([][]Color, h)
	defaultColor := NewColor(0, 0, 0)
	for i := 0; i < h; i++ {
		row := make([]Color, w)
		for j := 0; j < w; j++ {
			row[j] = defaultColor
		}
		canvas[i] = row
	}
	return canvas
}

func (c *Canvas) SetEveryPixel(color Color) {
	canvas := *c
	x, y := c.GetWidthAndHeight()
	for i := 0; i < y; i++ {
		for j := 0; j < x; j++ {
			canvas[i][j] = color
		}
	}
}

func (c *Canvas) Get(x, y int64) Color {
	canvas := *c
	return canvas[y][x]
}

func (c *Canvas) WritePixel(x, y int64, color Color) {
	canvas := *c
	canvas[y][x] = color
}

func (c *Canvas) GetWidthAndHeight() (int, int) {
	canvas := *c
	return len(canvas[0]), len(canvas)
}

func (c *Canvas) ToPPM() string {
	canvas := *c
	content := strings.Builder{}
	lineMaxLen := 70

	w, h := c.GetWidthAndHeight()
	header := fmt.Sprintf("P3\n%d %d\n255\n", w, h)
	content.WriteString(header)

	// TODO: Fix this ugliest piece of code ever
	for _, row := range canvas {
		sb := strings.Builder{}
		lineLength := 0
		for _, pixel := range row {
			r := scale(pixel.red)
			rStr := fmt.Sprintf("%d ", r)
			if lineMaxLen-1 < len(rStr)+lineLength {
				content.WriteString(strings.TrimSuffix(sb.String(), " "))
				content.WriteString("\n")
				sb.Reset()
				lineLength = 0
			}
			lineLength = lineLength + len(rStr)
			sb.WriteString(rStr)
			g := scale(pixel.green)
			gStr := fmt.Sprintf("%d ", g)
			if lineMaxLen-1 < len(gStr)+lineLength {
				content.WriteString(strings.TrimSuffix(sb.String(), " "))
				content.WriteString("\n")
				sb.Reset()
				lineLength = 0
			}
			sb.WriteString(gStr)
			lineLength = lineLength + len(gStr)
			b := scale(pixel.blue)
			bStr := fmt.Sprintf("%d ", b)
			if lineMaxLen-1 < len(bStr)+lineLength {
				content.WriteString(strings.TrimSuffix(sb.String(), " "))
				content.WriteString("\n")
				sb.Reset()
				lineLength = 0
			}
			sb.WriteString(bStr)
			lineLength = lineLength + len(bStr)
		}
		s := strings.TrimSuffix(sb.String(), " ")
		content.WriteString(s)
		content.WriteString("\n")
		lineLength = 0

	}
	content.WriteString("\n")
	return content.String()
}

func scale(value float64) int64 {
	// Assuming the value given is between 0..1
	if value < 0.0 {
		return 0.0
	}
	scaled := math.Ceil(value * 255.0)
	return int64(math.Min(scaled, 255.0))
}

func NewColor(red, green, blue float64) Color {
	return Color{red, green, blue}
}

func NewPoint(x, y, z float64) Tuple {
	return Tuple{x, y, z, 1.0}
}

func NewVector(x, y, z float64) Tuple {
	return Tuple{x, y, z, 0.0}
}

func (t Tuple) Equals(b Tuple) bool {
	return equals(t.X, b.X) && equals(t.Y, b.Y) && equals(t.Z, b.Z) // TODO: w as well ?
}

func Add(a, b Tuple) Tuple {
	x := a.X + b.X
	y := a.Y + b.Y
	z := a.Z + b.Z
	w := a.W + b.W

	return Tuple{x, y, z, w}
}

func Substract(a, b Tuple) Tuple {
	x := a.X - b.X
	y := a.Y - b.Y
	z := a.Z - b.Z
	w := a.W - b.W

	return Tuple{x, y, z, w}
}

func Negate(a Tuple) Tuple {
	return Tuple{-a.X, -a.Y, -a.Z, -a.W}
}

func Multiply(a Tuple, i float64) Tuple {
	return Tuple{i * a.X, i * a.Y, i * a.Z, i * a.W}
}

func Magnitude(a Tuple) float64 {
	return math.Sqrt(a.X*a.X + a.Y*a.Y + a.Z*a.Z + a.W*a.W)
}

func Normalize(a Tuple) Tuple {
	mag := Magnitude(a)
	return Tuple{a.X / mag, a.Y / mag, a.Z / mag, a.W / mag}
}

func Dot(a, b Tuple) float64 {
	return a.X*b.X + a.Y*b.Y + a.Z*b.Z + a.W*b.W
}

func Cross(a, b Tuple) Tuple {
	return NewVector(a.Y*b.Z-a.Z*b.Y,
		a.Z*b.X-a.X*b.Z,
		a.X*b.Y-a.Y*b.X)
}

func (c Color) Equals(b Color) bool {
	return equals(c.red, b.red) && equals(c.green, b.green) && equals(c.blue, b.blue)
}

func (c Color) Add(b Color) Color {
	return Color{c.red + b.red, c.green + b.green, c.blue + b.blue}
}

func (c Color) Sub(b Color) Color {
	return Color{c.red - b.red, c.green - b.green, c.blue - b.blue}
}

func (c Color) Times(x float64) Color {
	return Color{c.red * x, c.green * x, c.blue * x}
}

func (c Color) Blend(b Color) Color {
	return Color{c.red * b.red, c.green * b.green, c.blue * b.blue}
}
