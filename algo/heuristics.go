package algo

import (
	"math"
	"reflect"

	"github.com/MonsieurTa/n-puzzle/state"
	"github.com/MonsieurTa/n-puzzle/utils"
)

var Heuristics map[string](func([][]int, *state.State) int) = map[string](func([][]int, *state.State) int){
	"hamming":   Hamming,
	"gasching":  Gasching,
	"manhattan": Manhattan,
	"euclidean": Euclidean,
	"conflicts": ManhattanXLinear,
}

func Hamming(a [][]int, b *state.State) int {
	res := 0
	size := len(b.Board)
	for i := 0; i < size; i++ {
		for j := 0; j < size; j++ {
			if a[i][j] != 0 && a[i][j] != b.Board[i][j] {
				res++
			}
		}
	}
	return res
}

func Gasching(a [][]int, b *state.State) int {
	res := 0
	curr := utils.FlattenArray(a)
	goal := utils.FlattenArray(b.Board)
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

func Manhattan(a [][]int, b *state.State) int {
	return distance(a, b, func(x1, x2, y1, y2 int) int {
		return int(math.Abs(float64(x2-x1)) + math.Abs(float64(y2-y1)))
	})
}

func Euclidean(a [][]int, b *state.State) int {
	return distance(a, b, func(x1, x2, y1, y2 int) int {
		return (x1-x2)*(x1-x2) + (y1-y2)*(y1-y2)
	})
}

func distance(a [][]int, b *state.State, get func(int, int, int, int) int) int {
	res := 0
	size := len(b.Board)
	for i := 0; i < size; i++ {
		for j := 0; j < size; j++ {
			if a[i][j] != 0 && a[i][j] != b.Board[i][j] {
				x2, y2 := b.CachedPos(a[i][j])
				res += get(j, x2, i, y2)
			}
		}
	}
	return res
}

func rowConflicts(x, y int, row []int, goal *state.State, rowMap map[int]bool) int {
	res := 0
	for i := range row {
		if i != x {
			if xGoal, yGoal := goal.CachedPos(row[i]); yGoal == y {

				res++
			}
		}
	}
	return res
}

func columnConflicts(x, y int, goal *state.State) int {
	res := 0

	return res
}

func LinearConflict(a [][]int, b *state.State) int {
	res := 0

	for y, row := range a {
		for x, tile := range row {
			if tile != 0 && tile != b.Board[y][x] {
				goalX, goalY := b.CachedPos(tile)
				if x == goalX {
					// COLUMN CONFLICT
				} else if y == goalY {
					// ROW CONFLICT
				}
			}
		}
	}
	return res
}

func ManhattanXLinear(a [][]int, b *state.State) int {
	return Manhattan(a, b) + LinearConflict(a, b)
}
