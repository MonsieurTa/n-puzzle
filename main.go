package main

import (
	"fmt"
	"os"

	"github.com/MonsieurTa/n-puzzle/parser"
)

func main() {
	tab, err := parser.Parse()
	if err != nil {
		fmt.Fprintf(os.Stderr, "%v\n", err)
	} else {
		for i := 0; i < len(tab); i++ {
			for j := 0; j < len(tab[i]); j++ {
				fmt.Printf("x:%d - y:%d = %d\n", i, j, tab[i][j])
			}
		}
	}
}
