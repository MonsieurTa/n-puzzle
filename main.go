package main

import (
	"github.com/MonsieurTa/n-puzzle/algo"
	"github.com/MonsieurTa/n-puzzle/utils"
)

var start = [][]int{
	// {1, 8, 6},
	// {2, 0, 7},
	// {3, 4, 5},
	{1, 2, 3},
	{8, 6, 4},
	{0, 7, 5},
}

func main() {
	var a algo.Algo
	board := utils.SnailArray(len(start))
	a.Init(len(board))
	x, y := algo.GetRootPos(start)
	root := algo.Node{
		ID: start[y][x],
		X:  x, Y: y,
		Hash:  algo.HashState(start),
		State: start,
	}
	x, y = algo.GetRootPos(board)
	goal := algo.Node{
		ID: board[y][x],
		X:  x, Y: y,
		Hash:  algo.HashState(board),
		State: board,
	}
	algo.DisplayState(&goal)
	result := a.AStar(&root, &goal, func(a, b *algo.Node) int {
		diff := 0
		size := len(b.State)
		for i := 0; i < size; i++ {
			for j := 0; j < size; j++ {
				if a.State[i][j] != b.State[i][j] {
					diff++
				}
			}
		}
		return diff
	})
	println(len(result))
	for _, node := range result {
		algo.DisplayState(node)
	}
}
