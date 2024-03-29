package parser

import (
	"bufio"
	"errors"
	"fmt"
	"strconv"
	"strings"

	"github.com/MonsieurTa/n-puzzle/utils"
)

// Parses the data it reads in data.File, returning
// a 2d int slice and an error.
func Parse(data *Data) ([][]int, error) {
	reader := bufio.NewReader(data.File)
	sizeStr, err := readLine(reader)
	if err != nil {
		return nil, err
	}
	sizeStr = strings.TrimSpace(sizeStr)
	size, err := strconv.Atoi(sizeStr)
	if err != nil {
		return nil, fmt.Errorf("n-puzzle: got %s while expected an integer", sizeStr)
	}
	if size <= 1 {
		return nil, errors.New("n-puzzle: size must be greater than 1")
	}
	tab := make([][]int, size)
	for i := 0; i < size; i++ {
		str, err := readLine(reader)
		if err != nil {
			return nil, err
		}
		numbers, err := checkValidity(str, size, i)
		if err != nil {
			return nil, err
		}
		tab[i] = numbers
	}
	err = validateTab(tab, size)
	if err != nil {
		return nil, err
	}
	return tab, nil
}

// Reads a line from the reader given in params
// Handles commentaries and returns string data or the error
func readLine(reader *bufio.Reader) (string, error) {
	str, err := reader.ReadString('\n')
	if err != nil {
		return "", fmt.Errorf("n-puzzle: error reading data from input")
	}
	hashtagIndex := strings.IndexByte(str, '#')
	if hashtagIndex >= 0 {
		str = str[:hashtagIndex]
	}
	str = strings.TrimSpace(str)
	if len(str) > 0 {
		return str, nil
	}
	return readLine(reader)
}

// Function used to parse lines of ints, and check their validity
// Returns an int slice or an error
func checkValidity(str string, size int, line int) ([]int, error) {
	numbers := make([]int, size)
	words := strings.Fields(str)
	for i := 0; i < len(words); i++ {
		nbr, err := strconv.Atoi(words[i])
		if err != nil || nbr < 0 {
			return nil, fmt.Errorf("n-puzzle: got %s while expected an unsigned integer on line %d", words[i], line+1)
		}
		if i >= size {
			return nil, fmt.Errorf("n-puzzle: line %d has %d numbers instead of %d", line+1, len(words), size)
		}
		numbers[i] = nbr
	}
	if len(words) < size {
		return nil, fmt.Errorf("n-puzzle: line %d has %d numbers instead of %d", line+1, len(words), size)
	}
	return numbers, nil
}

// Returns an error not, whether the 2d int slice is valid or not
func validateTab(tab [][]int, size int) error {
	maxNbr := (size * size) - 1
	foundNbrs := make([]int, size*size)
	for i := 0; i < size*size; i++ {
		foundNbrs[i] = -1
	}
	for _, row := range tab {
		for _, nbr := range row {
			if nbr > maxNbr {
				return fmt.Errorf("n-puzzle: got %d while the maximum number is %d", nbr, maxNbr)
			}
			if utils.ContainsInt(foundNbrs, nbr) {
				return fmt.Errorf("n-puzzle: found the number %d more than once", nbr)
			}
			foundNbrs = append(foundNbrs, nbr)
		}
	}
	return nil
}
