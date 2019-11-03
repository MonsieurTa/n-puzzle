package algo

import (
	"math"
	"reflect"

	"github.com/MonsieurTa/n-puzzle/utils"
)

var Heuristics map[string](func([][]int, [][]int) int) = map[string](func([][]int, [][]int) int){
	"hamming":   Hamming,
	"gasching":  Gasching,
	"manhattan": Manhattan,
	"Euclidean": Euclidean,
	"conflicts": ManhattanXLinear,
}

func Hamming(a, b [][]int) int {
	res := 0
	size := len(b)
	for i := 0; i < size; i++ {
		for j := 0; j < size; j++ {
			if a[i][j] != 0 && a[i][j] != b[i][j] {
				res++
			}
		}
	}
	return res
}

func Gasching(a, b [][]int) int {
	res := 0
	curr := utils.FlattenArray(a)
	goal := utils.FlattenArray(b)
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

func Manhattan(a, b [][]int) int {
	return distance(a, b, func(x1, x2, y1, y2 int) int {
		return int(math.Abs(float64(x2-x1)) + math.Abs(float64(y2-y1)))
	})
}

func Euclidean(a, b [][]int) int {
	return distance(a, b, func(x1, x2, y1, y2 int) int {
		return (x1-x2)*(x1-x2) + (y1-y2)*(y1-y2)
	})
}

func distance(a, b [][]int, get func(int, int, int, int) int) int {
	res := 0
	size := len(b)
	for i := 0; i < size; i++ {
		for j := 0; j < size; j++ {
			if a[i][j] != 0 && a[i][j] != b[i][j] {
				found := false
				for k := 0; k < size && !found; k++ {
					for l := 0; l < size && !found; l++ {
						if b[k][l] == a[i][j] {
							res += get(j, l, i, k)
							found = true
						}
					}
				}
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
	m := len(state) / 2
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
			if x > m && (inc == 1 && currValue > value || inc == -1 && currValue < value) {
				ret++
			} else if x < m && (inc == 1 && currValue < value || inc == -1 && currValue > value) {
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
	m := len(state) / 2
	if y > yy {
		inc = -1
	}
	y += inc
	for y != yy {
		currValue := state[y][x]
		if _, ok := isInGoalColumn(currValue, x, state, goal); ok {
			if y > m && (inc == 1 && currValue > value || inc == -1 && currValue < value) {
				ret++
			} else if y < m && (inc == 1 && currValue < value || inc == -1 && currValue > value) {
				ret++
			}
		}
		y += inc
	}
	return ret
}

func LinearConflict(a, b [][]int) int {
	ret := 0
	for y := range a {
		for x := range a[y] {
			value := a[y][x]
			if value != 0 {
				if xx, ok := isInGoalRow(value, y, a, b); ok {
					ret += searchRowConflict(x, xx, y, a, b)
				} else if yy, ok := isInGoalColumn(value, x, a, b); ok {
					ret += searchColumnConflict(y, yy, x, a, b)
				}
			}
		}
	}
	return ret * 2
}

func ManhattanXLinear(a, b [][]int) int {
	return Manhattan(a, b) + LinearConflict(a, b)
}
