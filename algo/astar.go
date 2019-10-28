package algo

import (
	"crypto/md5"
	"fmt"
)

type node struct {
	x, y  int
	hash  []byte
	state [][]int
}

func hashState(state [][]int) []byte {
	h := md5.New()
	base := ""
	for _, row := range state {
		base += fmt.Sprint(row)
	}
	return h.Sum([]byte(base))
}

func getLowestCost(list []node, h func([][]int) int) node {
	score := h(list[0].state)
	var index int
	for i, item := range list[1:] {
		ret := h(item.state)
		if ret < score {
			score = ret
			index = i
		}
	}
	res := list[index]
	list = append(list[:index], list[index+1:]...)
	return res
}

func getRootPos(board [][]int) (int, int) {
	for y, row := range board {
		for x, col := range row {
			if col == 0 {
				return x, y
			}
		}
	}
	return -1, -1
}

func processNode(elem node) []node {
	child := make([]node, 0)
	size := len(elem.state)

}

func AStar(start [][]int, goal [][]int, h func([][]int) int) [][][]int {
	x, y := getRootPos(start)
	startNode := node{
		x: x, y: y,
		hash:  hashState(start),
		state: start,
	}
	openSet := []node{startNode}
	closedSet := make([]node, 0)
	for len(openSet) > 0 {
		elem := getLowestCost(openSet, h)
		closedSet = append(closedSet, elem)
		child := processNode(elem)
	}
	return nil
}
