package main

type Tuple struct {
	x, y, z, w float64
}

func NewPoint(x, y, z float64) Tuple {
	return Tuple{x, y, z, 1.0}
}

func NewVector(x, y, z float64) Tuple {
	return Tuple{x, y, z, 0.0}
}

func Add(a, b Tuple) Tuple {
	x := a.x + b.x
	y := a.y + b.y
	z := a.z + b.z
	w := a.w + b.w

	return Tuple{x, y, z, w}
}

func Substract(a, b Tuple) Tuple {
	x := a.x - b.x
	y := a.y - b.y
	z := a.z - b.z
	w := a.w - b.w

	return Tuple{x, y, z, w}
}

