package main

import (
	"os"

	"fmt"
)

func main() {
	if len(os.Args) != 2 {
		return
	}
	s := os.Args[1]
	word := false
	printed := false

	for _, c := range s {
		if c != ' ' && c != '\t' {
			if word {
				fmt.Print(rune(' '))
				fmt.Print(rune(' '))
				fmt.Print(rune(' '))
			}
			fmt.Print(rune(c))
			word = true
			printed = true
		} else if word {
			word = false
		}
	}
	if printed {
		fmt.Print("\n")
	}
}
