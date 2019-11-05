package main

import (
	"fmt"
	"os"
	"strconv"

	"github.com/MonsieurTa/n-puzzle/algo"
	"github.com/MonsieurTa/n-puzzle/gen"
	"github.com/MonsieurTa/n-puzzle/parser"
	"github.com/MonsieurTa/n-puzzle/state"
)

func getCost(b bool) int {
	if b {
		return 0
	}
	return 1
}

func handleGenerator() (bool, error) {
	args := os.Args[1:]
	if len(args) > 0 && args[0] == "gen" {
		if len(args) < 2 {
			return true, fmt.Errorf("n-puzzle: generator: ./n-puzzle gen <nbr>\n")
		}
		nbr, err := strconv.Atoi(args[1])
		if err != nil {
			return true, fmt.Errorf("n-puzzle: got '%s' while expected an integer\n", args[1])
		}
		if nbr > 100 || nbr < 2 {
			return true, fmt.Errorf("n-puzzle: number must be between 2 and 100\n")
		}
		data := gen.Generate(nbr)
		fmt.Print(nbr, "\n")
		for _, line := range data {
			for _, el := range line {
				fmt.Print(el, " ")
			}
			fmt.Print("\n")
		}
		return true, nil
	}
	return false, nil
}

func main() {
	var npuzzle parser.Data
	var a algo.Algo

	generated, err := handleGenerator()
	if err != nil {
		fmt.Fprint(os.Stderr, err)
	}
	if generated {
		return
	}

	parser.ParseArgs(&npuzzle)
	start, err := parser.Parse(&npuzzle)
	if err != nil {
		fmt.Fprint(os.Stderr, err, "\n")
		return
	}
	goal := npuzzle.Goal(len(start))

	startState := state.NewState(start)
	goalState := state.NewState(goal)
	goalState.CacheBoard()
	a.Init(startState, goalState)

	fmt.Printf("Initial heuristic scores:\n")
	for key, value := range algo.Heuristics {
		fmt.Printf("- %s: %d\n", key, value(start, goalState))
	}
	if gen.IsSolvable(start, goal) {
		a.AStar(npuzzle.Heuristic, getCost(npuzzle.Greedy))
		if a.Path != nil {
			if npuzzle.JsonOutput {
				algo.OutputToJson(a.Path, goal)
			}
			fmt.Fprintf(npuzzle.Output, "Time complexity: %d nodes evaluated\n", a.Time)
			fmt.Fprintf(npuzzle.Output, "Space complexity: %d nodes in memory\n", a.Space)
			fmt.Fprintf(npuzzle.Output, "Moves required: %d\n", len(a.Path))
			return
		}
	}
	fmt.Fprint(npuzzle.Output, "This puzzle is unsolvable!\n")

	npuzzle.File.Close()
	npuzzle.Output.Close()
}
