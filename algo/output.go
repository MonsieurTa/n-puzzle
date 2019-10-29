package algo

import (
	"encoding/json"
	"fmt"
	"os"
)

type step struct {
	nbr int
	dir string
}

func OutputToJson(nodes []*Node, goal [][]int) {
	var data string = "const current = "
	nodesLen := len(nodes)
	start, err := json.Marshal(nodes[0].State)
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

	steps := make([]step, nodesLen)
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

func getStep(node1 *Node, node2 *Node) step {
	var len int = len(node1.State)
	var nbr int = node2.State[node1.Y][node1.X]
	var str string = "error"

	if node1.X+1 < len && node1.State[node1.Y][node1.X+1] == nbr {
		str = "left"
	} else if node1.X-1 >= 0 && node1.State[node1.Y][node1.X-1] == nbr {
		str = "right"
	} else if node1.Y+1 < len && node1.State[node1.Y+1][node1.X] == nbr {
		str = "up"
	} else if node1.Y-1 >= 0 && node1.State[node1.Y-1][node1.X] == nbr {
		str = "down"
	}

	return step{nbr, str}
}