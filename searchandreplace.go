//working piscine version with z01.PrintRune

package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	if len(os.Args) != 4 {
		return
	}
	str, old, new := os.Args[1], os.Args[2][0], os.Args[3][0]
	for i := 0; i < len(str); i++ {
		if str[i] == old {
			z01.PrintRune(rune(new))
		} else {
			z01.PrintRune(rune(str[i]))
		}
	}
	z01.PrintRune('\n')
}

//output
//PS C:\Users\Clay\checkpoints> go run testing.go "hello" "e" "a"
//hallo
//PS C:\Users\Clay\checkpoints>
