package main

import (
	"fmt"
	"os"

	"github.com/MonsieurTa/n-puzzle/algo"
	"github.com/MonsieurTa/n-puzzle/gen"
	"github.com/MonsieurTa/n-puzzle/parser"
	"github.com/MonsieurTa/n-puzzle/state"
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
	goal := npuzzle.Goal(len(start))

	startState := state.NewState(start)
	goalState := state.NewState(goal)
	a.Init(startState, goalState)

	fmt.Printf("Initial heuristic scores:\n")
	for key, value := range algo.Heuristics {
		fmt.Printf("- %s: %d\n", key, value(start, goal))
	}
	if gen.IsSolvable(start, goal) {
		a.AStar(npuzzle.Heuristic[0])
		if a.Path != nil {
			if npuzzle.JsonOutput {
				algo.OutputToJson(a.Path, goal)
			}
			fmt.Fprintf(npuzzle.Output, "Time complexity: %d nodes in memory\n", a.Time)
			fmt.Fprintf(npuzzle.Output, "Space complexity: %d nodes evaluated\n", a.Space)
			fmt.Fprintf(npuzzle.Output, "Moves required: %d\n", len(a.Path))
			return
		}
	}
	fmt.Fprint(npuzzle.Output, "# This puzzle is unsolvable\n")

	npuzzle.File.Close()
	npuzzle.Output.Close()
}
