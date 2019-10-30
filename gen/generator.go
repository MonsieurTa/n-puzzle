package gen

import (
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

func IsSolvable(data [][]int) bool {
	size := len(data)
	// Copying the 2d array to 1d
	arr := make([]int, 0)
	for _, row := range data {
		arr = append(arr, row...)
	}

	nbr := 0
	blankRowOdd := isBlankRowOdd(data)
	for i := 0; i < len(arr); i++ {
		for j := i + 1; j < len(arr); j++ {
			if arr[j] != 0 && arr[j] < arr[i] {
				nbr++
			}
		}
	}

	if size%2 != 0 {
		return nbr%2 != 0
	} else {
		return (nbr%2 == 0) == blankRowOdd
	}
}

func isBlankRowOdd(data [][]int) bool {
	size := len(data)
	for i := 0; i < size; i++ {
		for j := 0; j < size; j++ {
			if data[i][j] == 0 {
				return ((size - i) % 2) != 0
			}
		}
	}
	return false
}
