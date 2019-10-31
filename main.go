package main

import (
	"fmt"
	"os"

	"github.com/MonsieurTa/n-puzzle/algo"
	"github.com/MonsieurTa/n-puzzle/gen"
	"github.com/MonsieurTa/n-puzzle/parser"
)

func main() {
	var npuzzle parser.Data
	parser.ParseArgs(&npuzzle)

	start, err := parser.Parse(&npuzzle)
	if err != nil {
		fmt.Fprint(os.Stderr, err, "\n")
		return
	}

	var a algo.Algo
	board := npuzzle.Goal(len(start))
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
	fmt.Printf("Initial heuristic scores:\n")
	for key, value := range algo.Heuristics {
		fmt.Printf("- %s: %d\n", key, value(&root, &goal))
	}
	a.Init(len(board), &goal)
	if gen.IsSolvable(start, board) {
		result := a.AStar(&root, &goal, npuzzle.Greedy, npuzzle.Heuristic)
		if result.Nodes != nil {
			// for _, node := range result.Nodes {
			// 	displayState(&npuzzle, node)
			// }
			if npuzzle.JsonOutput {
				algo.OutputToJson(result.Nodes, goal.State)
			}
			fmt.Fprintf(npuzzle.Output, "Time complexity: %d nodes in memory\n", result.TimeComplex)
			fmt.Fprintf(npuzzle.Output, "Size complexity: %d nodes evaluated\n", result.SizeComplex)
			fmt.Fprintf(npuzzle.Output, "Moves required: %d\n", len(result.Nodes))
			return
		}
	}
	displayState(&npuzzle, &root)
	fmt.Fprint(npuzzle.Output, "# This puzzle is unsolvable\n")

	npuzzle.File.Close()
	npuzzle.Output.Close()
}

func displayState(data *parser.Data, a *algo.Node) {
	for _, row := range a.State {
		fmt.Fprintf(data.Output, "%v\n", row)
	}
	fmt.Fprint(data.Output, "\n")
}
