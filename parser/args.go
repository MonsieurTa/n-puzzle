package parser

import (
	"flag"
	"fmt"
	"os"

	"github.com/MonsieurTa/n-puzzle/state"
	"github.com/MonsieurTa/n-puzzle/utils"

	"github.com/MonsieurTa/n-puzzle/algo"
)

// Struct containing all n-puzzle data
type Data struct {
	Goal       goalFunction
	Heuristic  heuristic
	File       *os.File
	Output     *os.File
	JsonOutput bool
	Greedy     bool
}

// We store the data here, so that we don't need to pass it to functions anymore
var globalData *Data

// The function to generate the goal
type goalFunction func(int) [][]int

// String() method needed for flag#var
func (i *goalFunction) String() string {
	return "The goal function"
}

// Set() method needed for flag#var, used to set the goal function
func (i *goalFunction) Set(value string) error {
	if *i != nil {
		return fmt.Errorf("goal function already defined")
	}
	goal, ok := utils.Goals[value]
	if !ok {
		return fmt.Errorf("unknown goal")
	}
	*i = goal
	return nil
}

// Array of heuristic functions, typed to be used in flag#var
type heuristic func([][]int, *state.State) int

// String() method needed for flag#var
func (i *heuristic) String() string {
	return ""
}

// Set() method needed for flag#var, used to add a new heuristic
func (i *heuristic) Set(value string) error {
	if *i != nil {
		return fmt.Errorf("heuristic already selected")
	}
	heuristic, ok := algo.Heuristics[value]
	if !ok {
		return fmt.Errorf("unknown heuristic")
	}
	*i = heuristic
	return nil
}

// ParseArgs function will parse cli args, and return an error
func ParseArgs(data *Data) error {
	globalData = data
	initData()

	var inputFile string
	var outputFile string

	flag.Var(&data.Heuristic, "heuristic", "an heuristic algorithm between "+getHeuristicNames())
	flag.Var(&data.Goal, "goal", "a goal between "+getGoalNames())
	flag.StringVar(&inputFile, "f", "", "a file to read in, stdin by default")
	flag.StringVar(&outputFile, "o", "", "a file to output in, stdout by default")
	flag.BoolVar(&data.JsonOutput, "json", false, "output or not to json file")
	flag.BoolVar(&data.Greedy, "g", false, "greedy mode")
	flag.Parse()

	if len(inputFile) > 0 {
		file, err := os.Open(inputFile)
		data.File = file
		if err != nil {
			return err
		}
	}
	if len(outputFile) > 0 {
		file, err := os.Create(outputFile)
		data.Output = file
		if err != nil {
			return err
		}
	}
	if data.Heuristic == nil {
		data.Heuristic = algo.Manhattan
	}
	if data.Goal == nil {
		data.Goal = utils.SnailArray
	}
	return nil
}

// This function will concatenate the map of heuristic names
// into a large string, separed by commas
func getHeuristicNames() string {
	ret := ""
	for name := range algo.Heuristics {
		if len(ret) > 0 {
			ret += ", "
		}
		ret += name
	}
	return ret
}

// This function will concatenate the map of heuristic names
// into a large string, separed by commas
func getGoalNames() string {
	ret := ""
	for name := range utils.Goals {
		if len(ret) > 0 {
			ret += ", "
		}
		ret += name
	}
	return ret
}

// Inits the globalData struct
// Default File/Output to stdin/stdout
func initData() {
	globalData.File = os.Stdin
	globalData.Output = os.Stdout
}
