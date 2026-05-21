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
	if s == "" {
		fmt.Print("\n")
		return
	}
	w := []string{""}
	for i := 0; i < len(s); i++ {
		if s[i] == ' ' {
			w = append(w, "")
		} else {
			w[len(w)-1] += string(s[i])
		}
	}
	for i := len(w) - 1; i >= 0; i-- {
		for _, r := range w[i] {
			fmt.Print(string(r))
		}
		if i > 0 {
			fmt.Print(" ")
		}
	}
	fmt.Print("\n")
}

//similar to wordflip but takes the string as an argument instead