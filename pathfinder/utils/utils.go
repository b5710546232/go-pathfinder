package utils

import (
	"github.com/b5710546232/go-pathfinder/pathfinder/model"
)

func ReconstructPath(rows int, cols int, current model.Node, start model.Node, parents []int, result []model.PathNode) []model.PathNode {
	i := len(result) - 1

	for left, right := i, len(result)-1; left < right; left, right = left+1, right-1 {
		result[left], result[right] = result[right], result[left]
	}

	var idx int
	for current.X != start.X || current.Y != start.Y {
		result[i] = model.PathNode{X: current.X, Y: current.Y}
		idx = current.Y*cols + current.X
		current = model.Node{X: parents[idx] % cols, Y: parents[idx] / cols}
		i--
	}
	result[i] = model.PathNode{X: start.X, Y: start.Y}

	return result[i:]
}
