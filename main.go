package main

import (
	"math"

	"github.com/MonsieurTa/n-puzzle/algo"
	"github.com/MonsieurTa/n-puzzle/utils"
)

var start = [][]int{
	// {5, 4, 8},
	// {6, 0, 7},
	// {1, 3, 2},
	// {0, 3, 4},
	// {1, 6, 7},
	// {8, 5, 2},
	{1, 5, 2},
	{6, 0, 3},
	{7, 8, 4},
}

func main() {
	var a algo.Algo
	board := utils.SnailArray(len(start))
	x, y := algo.GetRootPos(start)
	root := algo.Node{
		Parent: nil,
		X:      x, Y: y,
		Hash:  algo.HashState(start),
		State: start,
	}
	x, y = algo.GetRootPos(board)
	goal := algo.Node{
		Parent: nil,
		X:      x, Y: y,
		Hash:  algo.HashState(board),
		State: board,
	}
	a.Init(len(board), &goal)
	algo.DisplayState(&goal)
	result := a.AStar(&root, &goal, func(a, b *algo.Node) int {
		res := 0
		size := len(b.State)
		for i := 0; i < size; i++ {
			for j := 0; j < size; j++ {
				var found bool = false
				var k int
				var l int
				for k = 0; k < size; k++ {
					for l = 0; l < size; l++ {
						if a.State[i][j] == b.State[k][l] {
							found = true
							break
						}
					}
					if found {
						break
					}
				}
				if found {
					res += int(math.Abs(float64(i-k)) + math.Abs(float64(j-l)))
				}
			}
		}
		// for i := 0; i < size; i++ {
		// 	for j := 0; j < size; j++ {
		// 		if a.State[i][j] != b.State[i][j] {
		// 			res++
		// 		}
		// 	}
		// }
		return res
	})
	println(len(result))
	for _, node := range result {
		algo.DisplayState(node)
	}
}
