package main

import (
	"fmt"
	"math"
	"os"

	"github.com/MonsieurTa/n-puzzle/algo"
	"github.com/MonsieurTa/n-puzzle/parser"
	"github.com/MonsieurTa/n-puzzle/utils"
)

func main() {
	start, err := parser.Parse()
	if err != nil {
		fmt.Fprint(os.Stderr, err, "\n")
		return
	}

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
