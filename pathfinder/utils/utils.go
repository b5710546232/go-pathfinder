package utils

import "github.com/b5710546232/go-pathfinder/pathfinder/model"

func ReconstructPath(rows int, cols int, current model.Node, start model.Node, parents []int) []model.PathNode {
	path := make([]model.PathNode, 0, rows*cols)
	for current.X != start.X || current.Y != start.Y {
		c := model.PathNode{X: current.X, Y: current.Y}
		path = append(path, c)
		current = model.Node{X: parents[current.Y*cols+current.X] % cols, Y: parents[current.Y*cols+current.X] / cols}
	}
	path = append(path, model.PathNode{X: start.X, Y: start.Y})
	for i := 0; i < len(path)/2; i++ {
		path[i], path[len(path)-i-1] = path[len(path)-i-1], path[i]
	}

	return path
}

func IsOutOfBounds(grid [][]model.Node, x int, y int) bool {
	return x < 0 || x >= len(grid[0]) || y < 0 || y >= len(grid)
}
