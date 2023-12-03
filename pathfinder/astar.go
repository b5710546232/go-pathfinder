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
	dirs      []directions.Direction
	grid      [][]model.Node
	parents   []int
	visited   []bool
	gScore    []int
	pq        collections.MinHeap
	result    []model.PathNode
	heuristic func(int, int, int, int) int
}

func WithHeuristicFunc(heuristic func(int, int, int, int) int) func(*AStarPathFinder) {
	return func(p *AStarPathFinder) {
		p.heuristic = heuristic
	}
}

func NewAStarPathFinder(options ...func(*AStarPathFinder)) *AStarPathFinder {
	// default value
	finder := &AStarPathFinder{
		grid:      [][]model.Node{},
		dirs:      directions.DIRS_8,
		heuristic: heuristics.ManhattanDistanceHeuristic,
	}
	for _, o := range options {
		o(finder)
	}
	return finder
}

func WithDirs(dirs []directions.Direction) func(*AStarPathFinder) {
	return func(p *AStarPathFinder) {
		p.dirs = dirs
	}
}

func WithGrid(grid [][]model.Node) func(*AStarPathFinder) {
	rows, cols := len(grid), len(grid[0])

	pq := collections.NewMinHeap(rows * cols)
	parents := make([]int, rows*cols)  // i is x, j is y, parents[i*cols+j] is parent of (i, j)
	visited := make([]bool, rows*cols) // i is x, j is y, visited[i*cols+j] is whether (i, j) is visited
	gScore := make([]int, rows*cols)   // i is x, j is y, gScore[i*cols+j] is gScore of (i, j)
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

// Search is the main function of AStarPathFinder which will find the shortest path from start to end
// It will return a slice of PathNode which is the shortest path from start to end
// It use ManhattanDistanceHeuristic as default heuristic function
func (a AStarPathFinder) Search(start model.Node, end model.Node) []model.PathNode {
	grid := a.grid
	rows, cols := len(grid), len(grid[0])
	parents := a.parents
	visited := a.visited
	gScore := a.gScore
	heuristicFunc := a.heuristic
	// reset
	for i := 0; i < rows*cols; i++ {
		gScore[i] = math.MaxInt
		parents[i] = 0
		visited[i] = false
	}

	pq := a.pq
	pq.Reset()
	pq.Push(start, 0)

	gScore[start.Y*cols+start.X] = 0

	for pq.Len() > 0 {
		current := pq.Pop()
		if current.Equals(end) {
			// # Reconstruct the shortest path
			return utils.ReconstructPath(rows, cols, current, start, parents, a.result)
		}

		for _, dir := range a.dirs {
			nextRow := current.Y + dir[1]
			nextCol := current.X + dir[0]
			if nextRow < 0 || nextRow >= rows {
				continue
			}
			if nextCol < 0 || nextCol >= cols {
				continue
			}
			next := grid[nextRow][nextCol]
			if !next.IsWalkable() {
				continue
			}
			nextIndex := next.Y*cols + next.X
			if visited[nextIndex] {
				continue
			}

			tentativeGScore := gScore[current.Y*cols+current.X] + heuristicFunc((current.X), (current.Y), (next.X), (next.Y))

			if tentativeGScore < gScore[nextIndex] {
				parents[nextIndex] = current.Y*cols + current.X
				gScore[nextIndex] = tentativeGScore
				pq.Push(next, tentativeGScore+heuristicFunc((next.X), (next.Y), (end.X), (end.Y)))
			}
		}
		visited[current.Y*cols+current.X] = true
	}

	return []model.PathNode{} // No path found
}
