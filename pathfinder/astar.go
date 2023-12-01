package pathfinder

import (
	"math"

	"github.com/b5710546232/go-pathfinder/internal/collections"
	"github.com/b5710546232/go-pathfinder/pathfinder/directions"
	"github.com/b5710546232/go-pathfinder/pathfinder/heuristics"
	"github.com/b5710546232/go-pathfinder/pathfinder/model"
	"github.com/b5710546232/go-pathfinder/pathfinder/utils"
)

type AStarPathFinder struct {
	heuristicFunc heuristics.Heuristic
	dirs          []directions.Direction
	grid          [][]model.Node
	parents       []int
	visited       []bool
	gScore        []float64
	pq            collections.MinHeap
	result        []model.PathNode
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
	return func(p *AStarPathFinder) {
		p.heuristicFunc = heuristic
	}
}

func WithDirs(dirs []directions.Direction) func(*AStarPathFinder) {
	return func(p *AStarPathFinder) {
		p.dirs = dirs
	}
}

func WithGrid(grid [][]model.Node) func(*AStarPathFinder) {
	rows, cols := len(grid), len(grid[0])

	pq := collections.NewMinHeap(rows * cols)
	parents := make([]int, rows*cols)    // i is x, j is y, parents[i*cols+j] is parent of (i, j)
	visited := make([]bool, rows*cols)   // i is x, j is y, visited[i*cols+j] is whether (i, j) is visited
	gScore := make([]float64, rows*cols) // i is x, j is y, gScore[i*cols+j] is gScore of (i, j)
	result := make([]model.PathNode, rows*cols)
	return func(p *AStarPathFinder) {
		p.grid = grid
		p.parents = parents
		p.visited = visited
		p.gScore = gScore
		p.pq = pq
		p.result = result
	}
}

func (a AStarPathFinder) Search(start model.Node, end model.Node) []model.PathNode {

	grid := a.grid
	rows, cols := len(grid), len(grid[0])
	parents := a.parents
	visited := a.visited
	gScore := a.gScore
	// reset
	for i := 0; i < rows*cols; i++ {
		parents[i] = 0
		visited[i] = false
	}

	pq := a.pq
	pq.Reset()
	pq.Push(start, 0)

	for i := range gScore {
		gScore[i] = math.MaxFloat64
	}
	gScore[start.Y*cols+start.X] = 0

	for pq.Len() > 0 {
		current := pq.Pop()
		if current == end {
			// # Reconstruct the shortest path
			return utils.ReconstructPath(rows, cols, current, start, parents, a.result)
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
