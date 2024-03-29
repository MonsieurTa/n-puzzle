// Package gen provides functions to generate a n-puzzle grid
// It also provides functions to check their solvability
package gen

import (
	"math"
	"math/rand"
	"time"

	"github.com/MonsieurTa/n-puzzle/utils"
)

// Generate function will return a 2d slice of int, which size is in params
// Random seed is based on time
func Generate(size int) [][]int {
	ret := make([][]int, size)
	rand.Seed(time.Now().UTC().UnixNano())
	for i := 0; i < size; i++ {
		ret[i] = make([]int, size)
	}

	for i := 0; i < size*size; i++ {
		for {
			x, y := rand.Intn(size), rand.Intn(size)
			if ret[x][y] == 0 {
				ret[x][y] = i
				break
			}
		}
	}

	return ret
}

// IsSolvable function takes in params the start grid and the final one
// Returns a boolean whether the start grid is solvable or not
// It'll basically compare the manhattan distance between the zero in
// both grids % 2, to the number of inversions in the final grid
func IsSolvable(data [][]int, finalGrid [][]int) bool {
	// Converting 2d arrays to 1d
	arr := utils.FlattenArray(data)
	goalArr := utils.FlattenArray(finalGrid)

	nbr := 0
	for i := 0; i < len(arr)-1; i++ {
		for j := i + 1; j < len(arr); j++ {
			if utils.FindNbrIndex(goalArr, arr[j]) < utils.FindNbrIndex(goalArr, arr[i]) {
				nbr++
			}
		}
	}

	x, y := findZeroPosition(data)
	goalX, goalY := findZeroPosition(finalGrid)
	wantedModulo := int(math.Abs(float64(x-goalX)) + math.Abs(float64(y-goalY)))

	return (wantedModulo%2 == 0) == (nbr%2 == 0)
}

// Returns x and y of the 0 in the grid taken in params
func findZeroPosition(data [][]int) (int, int) {
	size := len(data)
	for i := 0; i < size; i++ {
		for j := 0; j < size; j++ {
			if data[i][j] == 0 {
				return i, j
			}
		}
	}
	return -1, -1
}
