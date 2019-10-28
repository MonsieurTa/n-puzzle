package utils

func DeepCopy(arr [][]int) [][]int {
	new := make([][]int, len(arr))
	for index, row := range arr {
		tmp := make([]int, len(row))
		copy(tmp, row)
		new[index] = tmp
	}
	return new
}

func ContainsInt(arr []int, nbr int) bool {
	for _, el := range arr {
		if el == nbr {
			return true
		}
	}
	return false
}
