package main

import "github.com/01-edu/z01"

func PrintMemory(arr [10]byte) {
	hex := "0123456789abcdef"

	for i, b := range arr {
		z01.PrintRune(rune(hex[b>>4]))
		z01.PrintRune(rune(hex[b&0x0f]))

		if i == 3 || i == 7 || i == 9 {
			z01.PrintRune('\n')
		} else {
			z01.PrintRune(' ')
		}
	}

	for _, b := range arr {
		if b >= 32 && b <= 126 {
			z01.PrintRune((rune(b)))
		} else {
			z01.PrintRune('.')
		}
	}
	z01.PrintRune('\n')
}

func main() {
	PrintMemory([10]byte{'h', 'e', 'l', 'l', 'o', 16, 21, '*'})
}
