package algo

import (
	"encoding/json"
	"fmt"
	"os"
)

type step struct {
	nbr string
	dir string
}

func OutputToJson(nodes []*Node, goal [][]int) {
	var data string = "const current = "
	nodesLen := len(nodes)
	if nodesLen == 0 {
		return
	}
	start, err := json.Marshal(nodes[0].State.Board)
	if err != nil {
		fmt.Fprint(os.Stderr, "n-puzzle: error converting start to JSON data")
		data += "'error'"
	} else {
		data += string(start)
	}
	data += "\nconst wanted = "
	grid, err := json.Marshal(goal)
	if err != nil {
		fmt.Fprint(os.Stderr, "n-puzzle: error converting goal to JSON data")
		data += "'error'"
	} else {
		data += string(grid)
	}
	data += "\nconst steps = "

	steps := make([]map[string]interface{}, nodesLen)
	for i := 0; i < nodesLen-1; i++ {
		steps[i] = getStep(nodes[i], nodes[i+1])
	}

	stepsJson, err := json.Marshal(steps)
	if err != nil {
		fmt.Fprint(os.Stderr, "n-puzzle: error converting steps to JSON data")
		data += "'error'"
	} else {
		data += string(stepsJson)
	}

	f, err := os.Create("./visu/data.js")
	if err != nil {
		fmt.Fprint(os.Stderr, "n-puzzle: error opening file ./visu/data.js")
	} else {
		_, err := f.WriteString(data)
		if err != nil {
			fmt.Fprintf(os.Stderr, "n-puzzle: error writing in file ./visu/data.js")
		}
	}
}

func getStep(node1 *Node, node2 *Node) map[string]interface{} {
	state1 := node1.State
	state2 := node2.State

	var len int = len(state1.Board)
	var nbr int = state2.Board[state1.BlankY][state1.BlankX]
	var str string = "error"

	if state1.BlankX+1 < len && state1.Board[state1.BlankY][state1.BlankX+1] == nbr {
		str = "left"
	} else if state1.BlankX-1 >= 0 && state1.Board[state1.BlankY][state1.BlankX-1] == nbr {
		str = "right"
	} else if state1.BlankY+1 < len && state1.Board[state1.BlankY+1][state1.BlankX] == nbr {
		str = "up"
	} else if state1.BlankY-1 >= 0 && state1.Board[state1.BlankY-1][state1.BlankX] == nbr {
		str = "down"
	}

	return map[string]interface{}{"nbr": nbr, "dir": str}
}
