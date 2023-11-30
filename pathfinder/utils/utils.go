package utils

import (
	"github.com/b5710546232/go-pathfinder/pathfinder/model"
)

func ReconstructPath(rows int, cols int, current model.Node, start model.Node, parents []int, result []model.PathNode) []model.PathNode {
	i := 0
	for current.X != start.X || current.Y != start.Y {
		c := model.PathNode{X: current.X, Y: current.Y}
		result[i] = c
		current = model.Node{X: parents[current.Y*cols+current.X] % cols, Y: parents[current.Y*cols+current.X] / cols}
		i++
	}
	result[i] = model.PathNode{X: start.X, Y: start.Y}
	i++
	path := result[:i]
	for i := 0; i < len(path)/2; i++ {
		path[i], path[len(path)-i-1] = path[len(path)-i-1], path[i]
	}
	return path
}

func IsOutOfBounds(grid [][]model.Node, x int, y int) bool {
	return x < 0 || x >= len(grid[0]) || y < 0 || y >= len(grid)
}
