package algo

import (
	"container/heap"

	"github.com/MonsieurTa/n-puzzle/state"
)

const MaxUint = ^uint(0)
const MaxInt = int(MaxUint >> 1)

//PriorityQueue ...
type PriorityQueue []*Node

//Len return the len of the prority queue
func (pq PriorityQueue) Len() int { return len(pq) }

//Less return priority condition
func (pq PriorityQueue) Less(i, j int) bool { return pq[i].f < pq[j].f }

func (pq PriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
	pq[i].index = j
	pq[j].index = i
}

func (pq *PriorityQueue) Push(x interface{}) {
	n := len(*pq)
	item := x.(*Node)
	item.index = n
	*pq = append(*pq, item)
}

func (pq *PriorityQueue) Pop() interface{} {
	old := *pq
	n := len(old)
	item := old[n-1]
	old[n-1] = nil  // avoid memory leak
	item.index = -1 // for safety
	*pq = old[0 : n-1]
	return item
}

// update modifies the priority and value of an Node in the queue.
func (pq *PriorityQueue) update(item *Node, g, f, h int) {
	item.h = h
	heap.Fix(pq, item.index)
}

//Node ...
type Node struct {
	index  int
	parent *Node
	State  *state.State
	h      int
	f      int
}

func (n Node) copy() *Node {
	return &Node{
		parent: n.parent,
		State:  n.State,
		h:      n.h,
	}
}

func (n Node) reconstructPath(cameFrom map[string]*Node, size int) []*Node {
	path := make([]*Node, size)
	elem := &n
	for elem.parent != nil {
		path = append([]*Node{elem}, path...)
		elem = cameFrom[elem.State.Str]
	}
	return path
}

func (n Node) getChild(fn func([][]int, [][]int) int, goal [][]int, gScore map[string]int) []*Node {
	childState := n.State.GetSurrounding()
	child := []*Node{}
	for _, item := range childState {
		if _, ok := gScore[item.Str]; !ok {
			gScore[item.Str] = MaxInt
		}
		h := fn(item.Board, goal)
		child = append(child, &Node{
			State: item,
			h:     h,
		})
	}
	return child
}

//Algo ...
type Algo struct {
	start *state.State
	goal  *state.State
	Path  []*Node
	Time  int
	Space int
}

func (a Algo) isGoal(state *state.State) bool {
	return a.goal.Str == state.Str
}

// Init set start and goal state to algo struct
func (a *Algo) Init(start, goal *state.State) {
	a.start, a.goal = start, goal
}

func (a *Algo) AStar(fn func([][]int, [][]int) int) {
	start := &Node{
		State: a.start,
		h:     fn(a.start.Board, a.goal.Board),
	}
	fScore := make(map[string]int)
	gScore := make(map[string]int)
	openSet := PriorityQueue{start}
	openSetMembers := make(map[string]string)
	fScore[start.State.Str] = start.h
	gScore[start.State.Str] = 0
	heap.Init(&openSet)
	closedSet := make(map[string]*Node, 0)
	cameFrom := make(map[string]*Node)
	for len(openSet) > 0 {
		elem := heap.Pop(&openSet).(*Node)
		delete(fScore, elem.State.Str)
		delete(openSetMembers, elem.State.Str)
		if a.isGoal(elem.State) {
			a.Time = len(closedSet)
			a.Space = len(closedSet) + len(openSet)
			a.Path = elem.reconstructPath(cameFrom, gScore[elem.State.Str])
			return
		}
		closedSet[elem.State.Str] = elem
		child := elem.getChild(fn, a.goal.Board, gScore)
		tentativeGScore := gScore[elem.State.Str] + 1
		for _, children := range child {
			if _, ok := closedSet[children.State.Str]; !ok {
				childrenGScore := gScore[children.State.Str]
				if tentativeGScore < childrenGScore {
					children.parent = elem
					cameFrom[children.State.Str] = elem
					gScore[children.State.Str] = tentativeGScore
					children.f = children.h + tentativeGScore
					fScore[children.State.Str] = children.f
					if _, ok := openSetMembers[children.State.Str]; !ok {
						heap.Push(&openSet, children)
						openSetMembers[children.State.Str] = children.State.Str
					}
				}
			}
		}
	}
}
