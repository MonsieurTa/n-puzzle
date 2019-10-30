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
