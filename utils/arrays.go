package utils

// DeepCopy copies a 2d int slice to a fresh new one
func DeepCopy(arr [][]int) [][]int {
	new := make([][]int, len(arr))
	for index, row := range arr {
		tmp := make([]int, len(row))
		copy(tmp, row)
		new[index] = tmp
	}
	return new
}

// ContainsInt basically returns an int whether the slice
// contains or not the int given
func ContainsInt(arr []int, nbr int) bool {
	return FindNbrIndex(arr, nbr) != -1
}

// Takes a 2d slice of ints and flattens it to 1d
func FlattenArray(data [][]int) []int {
	arr := make([]int, 0)
	for _, row := range data {
		arr = append(arr, row...)
	}
	return arr
}

// Returns the index of an int in an int slice
func FindNbrIndex(data []int, nbr int) int {
	size := len(data)
	for i := 0; i < size; i++ {
		if data[i] == nbr {
			return i
		}
	}
	return -1
}
