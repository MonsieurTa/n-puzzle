package algo

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"

	"github.com/MonsieurTa/n-puzzle/utils"
)

const MaxUint = ^uint(0)
const MaxInt = int(MaxUint >> 1)

type Node struct {
	Parent      *Node
	X, Y        int
	Hash        string
	State       [][]int
	Heuristique int
}

type Algo struct {
	Goal     *Node
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

func getLowestCost(mapper map[string]*Node, goal *Node, h func(*Node, *Node) int) *Node {
	var res *Node
	score := MaxInt
	for _, item := range mapper {
		if item.Heuristique < score {
			score = item.Heuristique
			res = item
		}
	}
	delete(mapper, res.Hash)
	return res
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

func clone(a *Node, b *Node) {
	b.Heuristique = a.Heuristique
	b.Hash = a.Hash
	b.Parent = a.Parent
	b.State = utils.DeepCopy(a.State)
	b.X = a.X
	b.Y = a.Y
}

func shiftTile(newX int, newY int, elem *Node, size int) *Node {
	var cpy Node = Node{}

	if newX < 0 || newY < 0 || newX >= size || newY >= size {
		return nil
	}
	clone(elem, &cpy)
	cpy.Parent = nil
	tmp := cpy.State[cpy.Y][cpy.X]
	cpy.State[cpy.Y][cpy.X] = cpy.State[newY][newX]
	cpy.State[newY][newX] = tmp
	cpy.X = newX
	cpy.Y = newY
	cpy.Hash = HashState(cpy.State)
	return &cpy
}

func processNode(elem *Node, a *Algo, h func(*Node, *Node) int) []*Node {
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
			item.Heuristique = h(item, a.Goal)
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
	for node.Parent != nil {
		node = node.Parent
		totalPath = append([]*Node{node}, totalPath...)
	}
	return totalPath
}

func (a *Algo) Init(size int, goal *Node) {
	totalSize := size * size
	a.CameFrom = make([]*Node, totalSize)
	a.GScore = make(map[string]int)
	a.FScore = make(map[string]int)
	a.Goal = goal
}

func DisplayState(a *Node) {
	for _, row := range a.State {
		fmt.Printf("%v\n\n", row)
	}
}

func (a *Algo) AStar(start *Node, goal *Node, h func(*Node, *Node) int) []*Node {
	openSet := map[string]*Node{start.Hash: start}
	closedSet := map[string]*Node{}
	a.GScore[start.Hash] = 0
	a.FScore[start.Hash] = h(start, goal)
	for len(openSet) > 0 {
		elem := getLowestCost(openSet, goal, h)
		if elem.Hash == goal.Hash {
			return reconstructPath(a.CameFrom, elem)
		}
		closedSet[elem.Hash] = elem
		child := processNode(elem, a, h)
		for _, children := range child {
			if _, ok := closedSet[children.Hash]; ok {
				continue
			}
			currGScore := a.GScore[elem.Hash] + 1
			if currGScore < a.GScore[children.Hash] {
				children.Parent = elem
				a.GScore[children.Hash] = currGScore
				// fmt.Printf("Heuristique: %d\n", h(children, goal))
				a.FScore[children.Hash] = a.GScore[children.Hash] + children.Heuristique
				if _, ok := openSet[children.Hash]; !ok {
					openSet[children.Hash] = children
				}
			}
		}
	}
	return nil
}
