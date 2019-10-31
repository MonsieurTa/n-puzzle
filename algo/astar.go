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

// update modifies the priority and value of an Item in the queue.
func (pq *PriorityQueue) update(item *Node, g, f, h int) {
	item.g = g
	item.f = f
	item.h = h
	heap.Fix(pq, item.index)
}

//Node ...
type Node struct {
	index  int
	parent *Node
	state  state.State
	g      int
	f      int
	h      int
}

func (n Node) copy() *Node {
	return &Node{
		parent: n.parent,
		state:  n.state,
		g:      n.g,
		f:      n.f,
		h:      n.h,
	}
}

func (n Node) getChild() []*Node {
	childState := n.state.GetSurrounding()
	child := []*Node{}
	for _, state := range childState {
		child = append(child, &Node{
			parent: &n,
			state:  *state,
			g:      n.g + 1,
			f:      0,
			h:      0,
		})
	}
}

//Algo ...
type Algo struct {
	start state.State
	goal  state.State
}

func (a Algo) isGoal(state state.State) bool {
	return a.goal.Str == state.Str
}

// Init set start and goal state to algo struct
func (a *Algo) Init(start, goal state.State) {
	a.start, a.goal = start, goal
}

func (a *Algo) AStar(h func(a, b *state.State) int) {
	openSet := PriorityQueue{&Node{
		state: a.start,
		f:     h(&a.start, &a.goal),
		g:     0,
		h:     h(&a.start, &a.goal),
	}}
	heap.Init(&openSet)
	closedSet := make(map[string]*Node, 0)
	for len(openSet) > 0 {
		elem := heap.Pop(&openSet).(*Node)
		if a.isGoal(elem.state) {
			return
		}
		closedSet[elem.state.Str] = elem
		child := elem.getChild()
		for _, children := range child {
			if _, ok := closedSet[children.state.Str]; ok {
				continue
			}

		}
	}
}
