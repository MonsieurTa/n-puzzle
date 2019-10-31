package parser

import (
	"flag"
	"fmt"
	"os"

	"github.com/MonsieurTa/n-puzzle/algo"
)

// Struct containing all n-puzzle data
type Data struct {
	heuristicNames map[string]bool
	Heuristic      heuristicArray
	File           *os.File
	Output         *os.File
	JsonOutput     bool
}

// Array of heuristic functions, typed for easier usage
type heuristicArray []func(*algo.Node, *algo.Node) int

// We store the data here, so that we don't need to pass it to functions anymore
var globalData *Data

// String() method needed for flag#var
func (i *heuristicArray) String() string {
	return "Some heuristics"
}

// Set() method needed for flag#var, used to add a new heuristic
func (i *heuristicArray) Set(value string) error {
	heuristic, ok := algo.Heuristics[value]
	if !ok {
		return fmt.Errorf("unknown heuristic")
	}
	ok, present := globalData.heuristicNames[value]
	if ok || present {
		return fmt.Errorf("heuristic already present")
	}
	*i = append(*i, heuristic)
	globalData.heuristicNames[value] = true
	return nil
}

// ParseArgs function will parse cli args, and return an error
func ParseArgs(data *Data) error {
	globalData = data
	initData()

	var inputFile string
	var outputFile string

	flag.Var(&data.Heuristic, "heuristic", "an heuristic algorithm between "+getHeuristicNames())
	flag.StringVar(&inputFile, "f", "", "a file to read in, stdin by default")
	flag.StringVar(&outputFile, "o", "", "a file to output in, stdout by default")
	flag.BoolVar(&data.JsonOutput, "json", false, "output or not to json file")
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
	println(len(data.Heuristic))
	if len(data.Heuristic) == 0 {
		data.Heuristic = append(data.Heuristic, algo.Manhattan)
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

// Inits the globalData struct
// Default File/Output to stdin/stdout
func initData() {
	globalData.heuristicNames = map[string]bool{}
	globalData.Heuristic = []func(*algo.Node, *algo.Node) int{}
	globalData.File = os.Stdin
	globalData.Output = os.Stdout
}
