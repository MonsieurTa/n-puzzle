package algo

import (
	"math"
	"reflect"

	"github.com/MonsieurTa/n-puzzle/utils"
)

var Heuristics map[string](func(*Node, *Node) int) = map[string](func(*Node, *Node) int){
	"hamming":   Hamming,
	"gasching":  Gasching,
	"manhattan": Manhattan,
	"euclidian": Euclidian,
	"conflicts": ManhattanXLinear,
}

func Hamming(a, b *Node) int {
	res := 0
	size := len(b.State)
	for i := 0; i < size; i++ {
		for j := 0; j < size; j++ {
			if a.State[i][j] != 0 && a.State[i][j] != b.State[i][j] {
				res++
			}
		}
	}
	return res
}

func Gasching(a, b *Node) int {
	res := 0
	curr := utils.FlattenArray(a.State)
	goal := utils.FlattenArray(b.State)
	size := len(curr)
	for {
		if reflect.DeepEqual(goal, curr) {
			break
		}
		currZeroIndex := utils.FindNbrIndex(curr, 0)
		if goalValue := goal[currZeroIndex]; goalValue != 0 {
			currIndex := utils.FindNbrIndex(curr, goalValue)
			curr[currIndex], curr[currZeroIndex] = curr[currZeroIndex], curr[currIndex]
		} else {
			for i := 0; i < size; i++ {
				if curr[i] != goal[i] {
					curr[i], curr[currZeroIndex] = curr[currZeroIndex], curr[i]
					break
				}
			}
		}

		res++
	}

	return res
}

func Manhattan(a, b *Node) int {
	return distanceHeuristic(a, b, func(x1, x2, y1, y2 int) int {
		return int(math.Abs(float64(x1-x2)) + math.Abs(float64(y1-y2)))
	})
}

func Euclidian(a, b *Node) int {
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
					if a.State[i][j] != 0 && a.State[i][j] == b.State[k][l] {
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

func isInGoalRow(value, row int, state [][]int, goal [][]int) (int, bool) {
	for i := 0; i < len(goal); i++ {
		if goal[row][i] == value {
			return i, true
		}
	}
	return 0, false
}

func isInGoalColumn(value, col int, state [][]int, goal [][]int) (int, bool) {
	for i := 0; i < len(goal); i++ {
		if goal[i][col] == value {
			return i, true
		}
	}
	return 0, false
}

func searchRowConflict(x, xx, y int, state, goal [][]int) int {
	inc := 1
	value := state[y][x]
	ret := 0
	if x > xx {
		inc = -1
	} else if x == xx {
		return 0
	}
	x += inc
	for x != xx {
		currValue := state[y][x]
		if _, ok := isInGoalRow(currValue, y, state, goal); ok {
			if currValue > value {
				ret++
			}
		}
		x += inc
	}
	return ret
}

func searchColumnConflict(y, yy, x int, state, goal [][]int) int {
	inc := 1
	value := state[y][x]
	ret := 0
	if y > yy {
		inc = -1
	}
	y += inc
	for y != yy {
		currValue := state[y][x]
		if _, ok := isInGoalColumn(currValue, x, state, goal); ok {
			if currValue > value {
				ret++
			}
		}
		y += inc
	}
	return ret
}

func LinearConflict(a, b *Node) int {
	ret := 0
	for y := range a.State {
		for x := range a.State[y] {
			value := a.State[y][x]
			if value != 0 {
				if xx, ok := isInGoalRow(value, y, a.State, b.State); ok {
					ret += searchRowConflict(x, xx, y, a.State, b.State)
				} else if yy, ok := isInGoalColumn(value, x, a.State, b.State); ok {
					ret += searchColumnConflict(y, yy, x, a.State, b.State)
				}
			}
		}
	}
	return ret * 2
}

func ManhattanXLinear(a, b *Node) int {
	return Manhattan(a, b) + LinearConflict(a, b)
}
