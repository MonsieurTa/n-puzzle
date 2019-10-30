package main

import (
	"fmt"
	"os"

	"github.com/MonsieurTa/n-puzzle/gen"

	"github.com/MonsieurTa/n-puzzle/algo"
	"github.com/MonsieurTa/n-puzzle/parser"
	"github.com/MonsieurTa/n-puzzle/utils"
)

func main() {
	// TODO: Handle those args
	// args := os.Args[1:]

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
	if gen.IsSolvable(board) {
		result := a.AStar(&root, &goal, algo.ManhattanHeuristic)
		println(len(result))
		for _, node := range result {
			algo.DisplayState(node)
		}
		algo.OutputToJson(result, goal.State)
	} else {
		fmt.Print("Unsolvable !")
	}
}
