package heuristics

import "math"

type Heuristic func(x1, y1, x2, y2 int) int

func ManhattanDistanceHeuristic(x1, y1, x2, y2 int) int {
	dx := abs(x1 - x2)
	dy := abs(y1 - y2)
	return dx + dy
}

func EuclideanDistanceHeuristic(x1, y1, x2, y2 int) int {
	dx := x2 - x1
	dy := y2 - y1
	return int(math.Round(math.Sqrt(float64(dx*dx + dy*dy))))
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
