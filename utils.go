package main

import "math"

func equals(a, b float64) bool {
	eps := 0.00001
	return math.Abs(a-b) < eps
}
