package algo

import (
	"bytes"
	"crypto/md5"
	"encoding/gob"
	"encoding/hex"
	"fmt"
)

const MaxUint = ^uint(0)
const MaxInt = int(MaxUint >> 1)

type Node struct {
	ID    int
	X, Y  int
	Hash  string
	State [][]int
}

type Algo struct {
	CameFrom []*Node
	GScore   map[string]int
	FScore   map[string]int
}

func HashState(State [][]int) string {
	base := ""
	for _, row := range State {
		base += fmt.Sprint(row)
	}
	hash := md5.Sum([]byte(base))
	res := hex.EncodeToString(hash[:])
	return res
}

func getLowestCost(list []*Node, goal *Node, h func(*Node, *Node) int) (*Node, []*Node) {
	index := 0
	score := h(list[index], goal)
	for i, item := range list[1:] {
		ret := h(item, goal)
		if ret < score {
			score = ret
			index = i
		}
	}
	res := list[index]
	list = append(list[:index], list[index+1:]...)
	return res, list
}

func GetRootPos(board [][]int) (int, int) {
	for Y, row := range board {
		for X, col := range row {
			if col == 0 {
				return X, Y
			}
		}
	}
	return -1, -1
}

func clone(a, b interface{}) {
	buff := new(bytes.Buffer)
	enc := gob.NewEncoder(buff)
	dec := gob.NewDecoder(buff)
	enc.Encode(a)
	dec.Decode(b)
}

func shiftTile(newX int, newY int, elem *Node, size int) *Node {
	var cpy *Node

	if newX < 0 || newY < 0 || newX >= size || newY >= size {
		return nil
	}
	clone(elem, &cpy)
	cpy.ID = elem.State[newY][newX]
	tmp := cpy.State[cpy.Y][cpy.X]
	cpy.State[cpy.Y][cpy.X] = cpy.State[newY][newX]
	cpy.State[newY][newX] = tmp
	cpy.X = newX
	cpy.Y = newY
	cpy.Hash = HashState(cpy.State)
	return cpy
}

func processNode(elem *Node, a *Algo) []*Node {
	child := make([]*Node, 0)
	size := len(elem.State)
	shifts := []*Node{
		shiftTile(elem.X, elem.Y-1, elem, size),
		shiftTile(elem.X, elem.Y+1, elem, size),
		shiftTile(elem.X-1, elem.Y, elem, size),
		shiftTile(elem.X+1, elem.Y, elem, size),
	}
	for _, item := range shifts {
		if item != nil {
			child = append(child, item)
			if _, ok := a.GScore[item.Hash]; !ok {
				a.GScore[item.Hash] = MaxInt
			}
			if _, ok := a.FScore[item.Hash]; !ok {
				a.GScore[item.Hash] = MaxInt
			}
		}
	}
	return child
}

func reconstructPath(cameFrom []*Node, current *Node) []*Node {
	node := current
	totalPath := []*Node{node}
	for {
		node = cameFrom[node.ID]
		if node == nil {
			break
		}
		totalPath = append([]*Node{node}, totalPath...)
	}
	return totalPath
}

func (a *Algo) Init(size int) {
	totalSize := size * size
	a.CameFrom = make([]*Node, totalSize)
	a.GScore = make(map[string]int)
	a.FScore = make(map[string]int)
}

func DisplayState(a *Node) {
	for _, row := range a.State {
		fmt.Printf("%v\n", row)
	}
	println()
}

func ContainsHash(array []*Node, ref *Node) bool {
	for _, elem := range array {
		if elem.Hash == ref.Hash {
			return true
		}
	}
	return false
}

func (a *Algo) AStar(start *Node, goal *Node, h func(*Node, *Node) int) []*Node {
	openSet := []*Node{start}
	closedSet := make([]*Node, 0)
	a.GScore[start.Hash] = 0
	a.FScore[start.Hash] = h(start, goal)
	for len(openSet) > 0 {
		elem, newOpenList := getLowestCost(openSet, goal, h)
		openSet = newOpenList
		if elem.Hash == goal.Hash {
			return reconstructPath(a.CameFrom, elem)
		}
		closedSet = append(closedSet, elem)
		child := processNode(elem, a)
		for _, children := range child {
			if ContainsHash(closedSet, children) {
				continue
			}
			currGScore := a.GScore[elem.Hash] + 1
			if currGScore < a.GScore[children.Hash] {
				a.CameFrom[children.ID] = elem
				a.GScore[children.Hash] = currGScore
				a.FScore[children.Hash] = a.GScore[children.Hash] + h(children, goal)
				if !ContainsHash(openSet, children) {
					openSet = append(openSet, children)
				}
			}
		}
	}
	return nil
}
