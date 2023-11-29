package heuristics

import "math"

type Heuristic func(x1, y1, x2, y2 float64) float64

func ManhattanDistanceHeuristic(x1, y1, x2, y2 float64) float64 {
	dx := math.Abs(x2 - x1)
	dy := math.Abs(y2 - y1)
	return dx + dy
}

func EuclideanDistanceHeuristic(x1, y1, x2, y2 float64) float64 {
	dx := x2 - x1
	dy := y2 - y1
	return math.Sqrt(dx*dx + dy*dy)
}
