package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	if len(os.Args) == 1 {
		return
	}

	for _, arg := range os.Args[1:] { //hii ni ya argumentrs zote
		for i := 0; i < len(arg); i++ { // hii ni ya one argument
			c := rune(arg[i])         // hii sasa inaconvert each string to runes ndio tuconvert to small kama ni captal
			if c >= 'A' && c <= 'Z' { //hii inacheck capital letters then inaziconvert zote to small
				c += 32
			}
			if i == len(arg)-1 || arg[i+1] == ' ' {
				if c >= 'a' && c <= 'z' {
					c -= 32
				}
			}
			z01.PrintRune(c)
		}
	}
	z01.PrintRune('\n')
}
