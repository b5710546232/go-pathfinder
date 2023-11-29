package pathfinder

import (
	"math"

	"github.com/b5710546232/pathfinder/internal/collections"
	"github.com/b5710546232/pathfinder/pathfinder/directions"
	"github.com/b5710546232/pathfinder/pathfinder/heuristics"
	"github.com/b5710546232/pathfinder/pathfinder/model"
	"github.com/b5710546232/pathfinder/pathfinder/utils"
)

type AStarPathFinder struct {
	heuristicFunc heuristics.Heuristic
	dirs          []directions.Direction
	grid          [][]model.Node
}

func NewAStarPathFinder(options ...func(*AStarPathFinder)) *AStarPathFinder {
	// default value
	finder := &AStarPathFinder{
		grid:          [][]model.Node{},
		dirs:          directions.DIRS_8,
		heuristicFunc: heuristics.ManhattanDistanceHeuristic,
	}
	for _, o := range options {
		o(finder)
	}
	return finder
}

func WithHeuristicFunc(heuristic heuristics.Heuristic) func(*AStarPathFinder) {
	return func(svr *AStarPathFinder) {
		svr.heuristicFunc = heuristic
	}
}

func WithDirs(dirs []directions.Direction) func(*AStarPathFinder) {
	return func(svr *AStarPathFinder) {
		svr.dirs = dirs
	}
}

func WithGrid(grid [][]model.Node) func(*AStarPathFinder) {
	return func(svr *AStarPathFinder) {
		svr.grid = grid
	}
}

func (a AStarPathFinder) Search(start model.Node, end model.Node) []model.PathNode {
	grid := a.grid
	rows, cols := len(grid), len(grid[0])

	pq := collections.NewMinHeap(rows * cols)
	pq.Push(start, 0)
	parents := make([]int, rows*cols)    // i is x, j is y, parents[i*cols+j] is parent of (i, j)
	visited := make([]bool, rows*cols)   // i is x, j is y, visited[i*cols+j] is whether (i, j) is visited
	gScore := make([]float64, rows*cols) // i is x, j is y, gScore[i*cols+j] is gScore of (i, j)
	for i := range gScore {
		gScore[i] = math.MaxFloat64
	}
	gScore[start.Y*cols+start.X] = 0

	for pq.Len() > 0 {
		current := pq.Pop()
		if current == end {
			// # Reconstruct the shortest path
			return utils.ReconstructPath(rows, cols, current, start, parents)
		}

		for _, dir := range a.dirs {
			nextRow := current.Y + dir[1]
			nextCol := current.X + dir[0]
			if utils.IsOutOfBounds(grid, nextCol, nextRow) {
				continue
			}
			next := grid[nextRow][nextCol]
			nextIndex := next.Y*cols + next.X
			if !next.IsWalkable() || visited[nextIndex] {
				continue
			}

			tentativeGScore := gScore[current.Y*cols+current.X] + a.heuristicFunc(float64(current.X), float64(current.Y), float64(next.X), float64(next.Y))

			if tentativeGScore < gScore[nextIndex] {
				parents[nextIndex] = current.Y*cols + current.X
				gScore[nextIndex] = tentativeGScore
				pq.Push(next, tentativeGScore+a.heuristicFunc(float64(next.X), float64(next.Y), float64(end.X), float64(end.Y)))
			}
		}
		visited[current.Y*cols+current.X] = true
	}

	return []model.PathNode{} // No path found
}
