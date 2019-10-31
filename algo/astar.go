package algo

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"math"

	"github.com/MonsieurTa/n-puzzle/utils"
)

const MaxUint = ^uint(0)
const MaxInt = int(MaxUint >> 1)

//Node ...
type Node struct {
	Parent    *Node
	X, Y      int
	Hash      string
	State     [][]int
	Heuristic int
	Fscore    int
}

//Algo ...
type Algo struct {
	Goal     *Node
	CameFrom []*Node
	GScore   map[string]int
	FScore   map[string]int
}

//Return ...
type Return struct {
	TimeComplex int
	SizeComplex int
	Nodes       []*Node
}

//HashState ...
func HashState(State [][]int) string {
	base := ""
	for _, row := range State {
		base += fmt.Sprint(row)
	}
	hash := md5.Sum([]byte(base))
	res := hex.EncodeToString(hash[:])
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
	b.Heuristic = a.Heuristic
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

func processNode(elem *Node, a *Algo, h []func(*Node, *Node) int) []*Node {
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
			item.Heuristic = applyHeuristics(item, a.Goal, h)
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

//Init ...
func (a *Algo) Init(size int, goal *Node) {
	totalSize := size * size
	a.CameFrom = make([]*Node, totalSize)
	a.GScore = make(map[string]int)
	a.FScore = make(map[string]int)
	a.Goal = goal
}

func binarySearch(list []*Node, x int) int {
	var ret int
	low := 0
	high := len(list) - 1

	for low <= high {
		m := int(math.Floor(float64(low+high) / 2))
		ret = m
		if list[m].Fscore < x {
			low = m + 1
		} else if list[m].Fscore > x {
			high = m - 1
		} else {
			return m
		}
	}
	if len(list) == 0 {
		return 0
	} else if list[ret].Fscore < x {
		return ret + 1
	}
	return ret
}

func applyHeuristics(node1 *Node, node2 *Node, arr []func(*Node, *Node) int) int {
	ret := 0
	for _, h := range arr {
		ret += h(node1, node2)
	}
	return ret
}

func isSort(list []*Node) bool {
	for i := 0; i < len(list)-1; i++ {
		if list[i].Fscore > list[i+1].Fscore {
			return false
		}
	}
	return true
}

func (a *Algo) AStar(start *Node, goal *Node, greedy bool, h []func(*Node, *Node) int) Return {
	var elem *Node
	// Initializing return
	var ret Return
	ret.Nodes = nil

	sortedNode := []*Node{start}
	openSet := map[string]*Node{start.Hash: start}
	closedSet := map[string]*Node{}
	a.GScore[start.Hash] = 0
	start.Fscore = applyHeuristics(start, goal, h)
	for len(openSet) > 0 {
		elem, sortedNode = sortedNode[0], sortedNode[1:]
		delete(openSet, elem.Hash)
		if elem.Hash == goal.Hash {
			ret.Nodes = reconstructPath(a.CameFrom, elem)
			ret.SizeComplex = len(openSet) + len(closedSet)
			ret.TimeComplex = len(closedSet)
			return ret
		}
		closedSet[elem.Hash] = elem
		child := processNode(elem, a, h)
		for _, children := range child {
			if _, ok := closedSet[children.Hash]; ok {
				continue
			}
			currGScore := a.GScore[elem.Hash] + 1
			if greedy == true || currGScore < a.GScore[children.Hash] {
				children.Parent = elem
				a.GScore[children.Hash] = currGScore
				if greedy {
					children.Fscore = children.Heuristic
				} else {
					children.Fscore = a.GScore[children.Hash] + children.Heuristic
				}
				if _, ok := openSet[children.Hash]; !ok {
					openSet[children.Hash] = children
					i := binarySearch(sortedNode, children.Fscore)
					sortedNode = append(sortedNode[:i], append([]*Node{children}, sortedNode[i:]...)...)
				}
			}
		}
	}
	return ret
}
