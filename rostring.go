package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	if len(os.Args) != 2 {
		z01.PrintRune('\n')
	}

	s := os.Args[1]
	words := []string{}
	word := ""

	for _, v := range s {
		if v != ' ' {
			word += string(v)
		} else if word != "" {
			words = append(words, word)
			word = ""
		}
	}

	if word != "" {
		words = append(words, word)
	}
	if len(words) == 0 {
		z01.PrintRune('\n')
		return
	}
	for i := 1; i < len(words); i++ {
		for _, v := range words[i] {
			z01.PrintRune(v)
		}
		z01.PrintRune(' ')
	}
	for _, v := range words[0] {
		z01.PrintRune(v)
	}
	z01.PrintRune('\n')
}
