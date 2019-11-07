package algo

import "container/heap"

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
