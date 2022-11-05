package models

import "math"

type Tuple struct {
	X, Y, Z, W float64
}

func NewPoint(x, y, z float64) Tuple {
	return Tuple{x, y, z, 1.0}
}

func NewVector(x, y, z float64) Tuple {
	return Tuple{x, y, z, 0.0}
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
