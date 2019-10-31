package utils

var Goals map[string](func(int) [][]int) = map[string](func(int) [][]int){
	"snail":     SnailArray,
	"zeroend":   ZeroEnd,
	"zerostart": ZeroStart,
}

// This functions generates the goal asked by the subject, as a snail
func SnailArray(size int) [][]int {
	snail := makeIntDoubleSlice(size)
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

// This functions generates a possible goal, with the zero at the end
func ZeroEnd(size int) [][]int {
	arr := makeIntDoubleSlice(size)
	i := 1
	for x := 0; x < size; x++ {
		for y := 0; y < size; y++ {
			arr[x][y] = i
			i++
		}
	}
	arr[size-1][size-1] = 0
	return arr
}

// This functions generates a possible goal, with the zero at the start
func ZeroStart(size int) [][]int {
	arr := makeIntDoubleSlice(size)
	i := 0
	for x := 0; x < size; x++ {
		for y := 0; y < size; y++ {
			arr[x][y] = i
			i++
		}
	}
	return arr
}

// Allocates a 2d int slice which size is in params
func makeIntDoubleSlice(size int) [][]int {
	arr := make([][]int, size)
	for i := 0; i < size; i++ {
		arr[i] = make([]int, size)
	}
	return arr
}
