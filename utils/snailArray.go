package utils

func SnailArray(size int) [][]int {
	snail := make([][]int, size)
	for i := 0; i < size; i++ {
		snail[i] = make([]int, size)
	}
	top, left, bottom, right := 0, 0, size-1, size-1
	i := 1
	maxSize := size * size
	for left < right {
		for c := left; c < right; c++ {
			snail[top][c] = i
			i++
		}
		for c := top; c < bottom; c++ {
			snail[c][right] = i
			i++
		}
		for c := right; c > left; c-- {
			snail[bottom][c] = i
			i++
		}
		for c := bottom; c > top; c-- {
			if i == maxSize {
				snail[c][left] = 0
			} else {
				snail[c][left] = i
			}
			i++
		}
		left++
		right--
		top++
		bottom--
	}
	return snail
}
