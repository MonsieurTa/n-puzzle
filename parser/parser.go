package parser

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func Parse() ([][]int, error) {
	reader := bufio.NewReader(os.Stdin)
	sizeStr, err := reader.ReadString('\n')
	if err != nil {
		return nil, err
	}
	sizeStr = strings.TrimSpace(sizeStr)
	size, err := strconv.Atoi(sizeStr)
	if err != nil {
		return nil, err
	}
	if size <= 1 {
		return nil, errors.New("n-puzzle: size must be greater than 1")
	}
	tab := make([][]int, size)
	for i := 0; i < size; i++ {
		str, err := reader.ReadString('\n')
		if err != nil {
			return nil, err
		}
		numbers, err := checkValidity(str, size, i)
		if err != nil {
			return nil, err
		}
		tab[i] = numbers
	}
	return tab, err
}

func checkValidity(str string, size int, line int) ([]int, error) {
	numbers := make([]int, size)
	var hashtagIndex int = strings.IndexByte(str, '#')
	if hashtagIndex > 0 {
		str = str[:hashtagIndex]
	}
	words := strings.Fields(str)
	for i := 0; i < len(words); i++ {
		nbr, err := strconv.Atoi(words[i])
		if err != nil {
			return nil, err
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
