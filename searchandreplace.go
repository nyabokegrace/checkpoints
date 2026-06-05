working piscine version with z01.PrintRune 

package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	if len(os.Args) != 4 {
		return
	}
	str1, str2, str3 := os.Args[1], os.Args[2][0], os.Args[3][0]
	for i := 0; i < len(str1); i++ {
		if str1[i] == str2 {
			z01.PrintRune(rune(str3))
		} else {
			z01.PrintRune(rune(str1[i]))
		}
	}
	z01.PrintRune('\n')
}


//output
//PS C:\Users\Clay\checkpoints> go run testing.go "hello" "e" "a"
//hallo
//PS C:\Users\Clay\checkpoints> 
