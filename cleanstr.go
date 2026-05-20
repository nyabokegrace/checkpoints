package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Println("\n")
		return
	}
	s := os.Args[1]
	word := false
	firstword := true

	for _, c := range s {
		if c != ' ' && c != '\t' {
			if !word && !firstword {
				fmt.Println(" ")
			}
			fmt.Println(c)
			word = true
			firstword = false
		} else {
			word = false
		}
	}
	fmt.Println("\n")
}
