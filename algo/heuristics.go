package algo

import "math"

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
	res := 0
	size := len(b.State)
	for i := 0; i < size; i++ {
		for j := 0; j < size; j++ {
			var found bool = false
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
				res += int(math.Abs(float64(i-k)) + math.Abs(float64(j-l)))
			}
		}
	}
	return res
}
