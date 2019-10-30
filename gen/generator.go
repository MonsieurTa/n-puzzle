package gen

import (
	"math"
	"math/rand"
	"time"
)

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

func IsSolvable(data [][]int, finalGrid [][]int) bool {
	// Converting 2d arrays to 1d
	arr := flattenArray(data)
	goalArr := flattenArray(finalGrid)

	nbr := 0
	for i := 0; i < len(arr)-1; i++ {
		for j := i + 1; j < len(arr); j++ {
			vi := arr[i]
			vj := arr[j]
			if findNbrIndex(goalArr, vj) < findNbrIndex(goalArr, vi) {
				nbr++
			}
		}
	}

	x, y := findZeroPosition(data)
	goalX, goalY := findZeroPosition(finalGrid)
	wantedModulo := int(math.Abs(float64(x-goalX)) + math.Abs(float64(y-goalY)))

	return (wantedModulo%2 == 0) == (nbr%2 == 0)
}

func flattenArray(data [][]int) []int {
	arr := make([]int, 0)
	for _, row := range data {
		arr = append(arr, row...)
	}
	return arr
}

func findNbrIndex(data []int, nbr int) int {
	size := len(data)
	for i := 0; i < size; i++ {
		if data[i] == nbr {
			return i
		}
	}
	return -1
}

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
