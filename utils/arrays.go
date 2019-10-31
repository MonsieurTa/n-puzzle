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
	for _, el := range arr {
		if el == nbr {
			return true
		}
	}
	return false
}
