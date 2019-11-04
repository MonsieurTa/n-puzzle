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
func (pq PriorityQueue) Less(i, j int) bool {
	return pq[i].f < pq[j].f
}

func (pq PriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
	pq[i].index = i
	pq[j].index = j
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
	*pq = old[:n-1]
	return item
}

// update modifies the priority and value of an Node in the queue.
func (pq *PriorityQueue) update(item *Node, f int) {
	item.f = f
	heap.Fix(pq, item.index)
}

//Node ...
type Node struct {
	index int
	State *state.State
	h     int
	f     int
}

func (n Node) reconstructPath(cameFrom map[string]*Node) []*Node {
	path := make([]*Node, 0)
	elem := &n
	for cameFrom[elem.State.Key] != nil {
		path = append([]*Node{elem}, path...)
		elem = cameFrom[elem.State.Key]
	}
	return path
}

func (n Node) getChild(fn func([][]int, [][]int) int, goal [][]int, gScore map[string]int) []*Node {
	childState := n.State.GetSurrounding()
	child := []*Node{}
	for _, item := range childState {
		if _, ok := gScore[item.Key]; !ok {
			gScore[item.Key] = MaxInt
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
	return a.goal.Key == state.Key
}

// Init set start and goal state to algo struct
func (a *Algo) Init(start, goal *state.State) {
	a.start, a.goal = start, goal
}

//AStar A* algorithm
func (a *Algo) AStar(fn func([][]int, [][]int) int, cost int) {
	start := &Node{
		State: a.start,
		h:     fn(a.start.Board, a.goal.Board),
	}
	openSet := PriorityQueue{start}
	openSetMembers := map[string]*Node{start.State.Key: start}
	gScore := map[string]int{start.State.Key: 0}
	closedSet := make(map[string]*Node)
	cameFrom := make(map[string]*Node)
	for len(openSet) > 0 {
		elem := heap.Pop(&openSet).(*Node)
		delete(openSetMembers, elem.State.Key)
		if a.isGoal(elem.State) {
			a.Time = len(closedSet)
			a.Space = len(closedSet) + len(openSet)
			a.Path = elem.reconstructPath(cameFrom)
			return
		}
		closedSet[elem.State.Key] = elem
		child := elem.getChild(fn, a.goal.Board, gScore)
		tentativeGScore := gScore[elem.State.Key] + cost
		for _, children := range child {
			if _, ok := closedSet[children.State.Key]; !ok {
				childrenGScore := gScore[children.State.Key]
				if tentativeGScore < childrenGScore {
					cameFrom[children.State.Key] = elem
					gScore[children.State.Key] = tentativeGScore
					f := children.h + tentativeGScore
					if item, ok := openSetMembers[children.State.Key]; !ok {
						children.f = f
						heap.Push(&openSet, children)
						openSetMembers[children.State.Key] = children
					} else {
						openSet.update(item, f)
					}
				}
			}
		}
	}
}
