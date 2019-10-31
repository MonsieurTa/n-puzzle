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
	var a algo.Algo

	parser.ParseArgs(&npuzzle)
	start, err := parser.Parse(&npuzzle)
	if err != nil {
		fmt.Fprint(os.Stderr, err, "\n")
		return
	}
	board := npuzzle.Goal(len(start))
	a.Init(start, board)

	fmt.Printf("Initial heuristic scores:\n")
	for key, value := range algo.Heuristics {
		fmt.Printf("- %s: %d\n", key, value(&root, &goal))
	}
	if gen.IsSolvable(start, board) {
		a.AStar(&root, &goal, npuzzle.Greedy, npuzzle.Heuristic)
		if a.Path != nil {
			if npuzzle.JsonOutput {
				algo.OutputToJson(a.Path, goal.State)
			}
			fmt.Fprintf(npuzzle.Output, "Time complexity: %d nodes in memory\n", a.TimeComplex)
			fmt.Fprintf(npuzzle.Output, "Size complexity: %d nodes evaluated\n", a.SizeComplex)
			fmt.Fprintf(npuzzle.Output, "Moves required: %d\n", len(a.Path))
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
