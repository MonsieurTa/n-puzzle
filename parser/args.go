package parser

import (
	"flag"
	"fmt"
	"os"

	"github.com/MonsieurTa/n-puzzle/algo"
)

type Data struct {
	Heuristic heuristicArray
	File      *os.File
	Output    *os.File
}

type heuristicArray []func(*algo.Node, *algo.Node) int

func (i *heuristicArray) String() string {
	return "Some heuristics"
}

func (i *heuristicArray) Set(value string) error {
	heuristic, ok := algo.Heuristics[value]
	if !ok {
		return fmt.Errorf("unknown heuristic")
	}
	*i = append(*i, heuristic)
	return nil
}

func ParseArgs(data *Data) error {
	initData(data)

	var inputFile string
	var outputFile string

	flag.Var(&data.Heuristic, "heuristic", "an heuristic algorithm between "+getHeuristicNames())
	flag.StringVar(&inputFile, "f", "", "a file to read in, stdin by default")
	flag.StringVar(&outputFile, "o", "", "a file to output in, stdout by default")
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
	if len(data.Heuristic) == 0 {
		data.Heuristic = append(data.Heuristic, algo.Manhattan)
	}
	return nil
}

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

func initData(data *Data) {
	data.Heuristic = []func(*algo.Node, *algo.Node) int{}
	data.File = os.Stdin
	data.Output = os.Stdout
}
