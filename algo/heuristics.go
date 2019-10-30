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
	return distanceHeuristic(a, b, func(x1, x2, y1, y2 int) int {
		return int(math.Abs(float64(x1-x2)) + math.Abs(float64(y1-y2)))
	})
}

func EuclidianHeuristic(a, b *Node) int {
	return distanceHeuristic(a, b, func(x1, x2, y1, y2 int) int {
		return int(math.Pow(float64(x1-x2), 2) + math.Pow(float64(y1-y2), 2))
	})
}

func distanceHeuristic(a, b *Node, get func(int, int, int, int) int) int {
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
				res += get(i, k, j, l)
			}
		}
	}
	return res
}
