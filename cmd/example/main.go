package main

import (
	"fmt"

	"github.com/b5710546232/go-pathfinder/pathfinder"
	"github.com/b5710546232/go-pathfinder/pathfinder/directions"
	"github.com/b5710546232/go-pathfinder/pathfinder/heuristics"
	"github.com/b5710546232/go-pathfinder/pathfinder/model"
)

func main() {
	grid := [][]int{
		{1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1},
		{1, 3, 0, 0, 1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1},
		{1, 1, 1, 0, 1, 0, 1, 0, 1, 0, 1, 1, 1, 1, 1, 1, 1, 1, 1, 0, 1},
		{1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 0, 1, 1, 0, 0, 1, 1, 0, 0, 0, 1},
		{1, 0, 1, 0, 1, 1, 1, 1, 1, 1, 0, 1, 1, 1, 0, 1, 1, 0, 1, 1, 1},
		{1, 0, 1, 0, 0, 0, 1, 0, 1, 0, 0, 1, 1, 0, 0, 1, 1, 0, 0, 0, 1},
		{1, 0, 1, 1, 1, 0, 1, 0, 1, 0, 1, 1, 1, 1, 1, 1, 1, 1, 1, 0, 1},
		{1, 0, 0, 0, 1, 0, 1, 0, 1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1, 0, 1},
		{1, 1, 1, 0, 1, 0, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 0, 1, 0, 1},
		{1, 0, 0, 0, 1, 0, 1, 0, 1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1, 0, 1},
		{1, 0, 1, 1, 1, 0, 1, 0, 1, 0, 1, 1, 1, 1, 1, 1, 1, 1, 1, 0, 1},
		{1, 0, 0, 0, 0, 0, 1, 0, 1, 0, 0, 0, 0, 1, 0, 0, 0, 0, 0, 0, 1},
		{1, 0, 1, 1, 1, 1, 1, 0, 1, 0, 0, 1, 1, 1, 0, 1, 1, 1, 1, 1, 1},
		{1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 4, 1},
		{1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1},
	}

	start := model.NewNode(0, 0)
	end := model.NewNode(0, 0)
	gridNode := make([][]model.Node, len(grid))
	//  3 is start, 4 is end
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[i]); j++ {
			n := model.NewNode(j, i)
			if grid[i][j] == 3 {
				start = n
			} else if grid[i][j] == 4 {
				end = n
			} else if grid[i][j] == 1 {
				n.SetWalkable(false)
			}
			gridNode[i] = append(gridNode[i], n)
		}
	}
	// print grid node
	for i := 0; i < len(gridNode); i++ {
		for j := 0; j < len(gridNode[i]); j++ {
			if start.Y == i && start.X == j {
				fmt.Print("S")
			} else if end.Y == i && end.X == j {
				fmt.Print("E")
			} else if !gridNode[i][j].IsWalkable() {
				fmt.Print("#")
			} else {
				fmt.Print(" ")
			}
		}
		fmt.Println()
	}
	dirs := directions.DIRS_8
	finder := pathfinder.NewAStarPathFinder(
		pathfinder.WithDirs(dirs),
		pathfinder.WithHeuristicFunc(heuristics.ManhattanDistanceHeuristic),
		pathfinder.WithGrid(gridNode),
	)
	paths := finder.Search(start, end)
	for _, p := range paths {
		grid[p.Y][p.X] = 2
	}
	fmt.Printf("\n---- result ----\n\n")
	// print path
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[i]); j++ {
			if start.Y == i && start.X == j {
				fmt.Print("S")
			} else if end.Y == i && end.X == j {
				fmt.Print("E")
			} else if grid[i][j] == 1 {
				fmt.Print("#")
			} else if grid[i][j] == 2 {
				fmt.Print(".")
			} else {
				fmt.Print(" ")
			}
		}
		fmt.Println()
	}
}
