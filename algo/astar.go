package algo

import (
	"container/heap"

	"github.com/MonsieurTa/n-puzzle/state"
)

const MaxUint = ^uint(0)
const MaxInt = int(MaxUint >> 1)

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

func (n Node) getChild(fn func([][]int, *state.State) int, goal *state.State, gScore map[string]int) []*Node {
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
func (a *Algo) AStar(fn func([][]int, *state.State) int, cost int) {
	start := &Node{
		State: a.start,
		h:     fn(a.start.Board, a.goal),
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
		child := elem.getChild(fn, a.goal, gScore)
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
