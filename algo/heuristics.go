package algo

import (
	"math"
)

var Heuristics map[string](func(*Node, *Node) int) = map[string](func(*Node, *Node) int){
	"hamming":   DefaultHeuristic,
	"manhattan": ManhattanHeuristic,
	"euclidian": EuclidianHeuristic,
}

func DefaultHeuristic(a, b *Node) int {
	res := 0
	size := len(b.State)
	for i := 0; i < size; i++ {
		for j := 0; j < size; j++ {
			if a.State[i][j] != b.State[i][j] {
				res++
			}
		}
	}
	return res
}

func ManhattanHeuristic(a, b *Node) int {
	return distanceHeuristic(a, b, func(x1, x2, y1, y2 int) int {
		return int(math.Abs(float64(x1-x2)) + math.Abs(float64(y1-y2)))
	})
}

func EuclidianHeuristic(a, b *Node) int {
	return distanceHeuristic(a, b, func(x1, x2, y1, y2 int) int {
		return (x1-x2)*(x1-x2) + (y1-y2)*(y1-y2)
	})
}

func distanceHeuristic(a, b *Node, get func(int, int, int, int) int) int {
	res := 0
	size := len(b.State)
	for i := 0; i < size; i++ {
		for j := 0; j < size; j++ {
			found := false
			var k int
			var l int
			for k = 0; k < size; k++ {
				for l = 0; l < size; l++ {
					if a.State[i][j] == b.State[k][l] {
						found = true
						break
					}
				}
				if found {
					break
				}
			}
			if found {
				res += get(i, k, j, l)
			}
		}
	}
	return res
}

func conflicDir(value, x, y int, goal *Node) (int, string) {
	for i := 0; i < len(goal.State); i++ {
		if goal.State[y][i] == value && i != x {
			return i, "row"
		} else if goal.State[i][x] == value && i != y {
			return i, "col"
		}
	}
	return -1, ""
}

func countColumnConflict(state [][]int, y, i int) int {
	var idx int
	offset := 1
	ret := 0

	if i > y {
		offset = -1
	}
	idx = i + offset
	for idx != y {
		if state[y][idx] != 0 && state[y][i] < state[y][i] {
			ret++
		}
		idx += offset
	}
	if state[y][idx] != 0 && state[y][i] < state[y][i] {
		ret++
	}
	return ret
}

func countRowConflict(state [][]int, x, i int) int {
	var idx int
	offset := 1
	ret := 0

	if i > x {
		offset = -1
	}
	idx = i + offset
	for idx != x {
		if state[idx][x] != 0 && state[i][x] < state[idx][x] {
			ret++
		}
		idx += offset
	}
	if state[idx][x] != 0 && state[i][x] < state[idx][x] {
		ret++
	}
	return ret
}

func LinearConflict(a, b *Node) int {
	res := 0
	for y := range a.State {
		for x := range a.State[y] {
			i, dir := conflicDir(a.State[y][x], x, y, b)
			if dir == "row" {
				res += countRowConflict(a.State, x, i)
			} else if dir == "col" {
				res += countColumnConflict(a.State, y, i)
			}
		}
	}
	return res * 2
}

func ManhattanXLinear(a, b *Node) int {
	return LinearConflict(a, b) + ManhattanHeuristic(a, b)
}
