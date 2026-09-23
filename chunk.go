// takes a slice and an integer as arguments then divides the slice into groups equal to the intrger provided.
package main

import (
	"github.com/01-edu/z01"
)

func Chunk(slice []int, size int) {
	if size == 0 {
		z01.PrintRune('\n')
	}
	for i := 0; i < len(slice); i += size { //loops through the slice jumping forward based on the size

		z01.PrintRune('[')
		first := true
		for j := i; j < i+size && j < len(slice); j++ { //loops through the elements of the chunk only from i to the end-1
			if !first {
				z01.PrintRune(' ')
			}
			putnbr(slice[j])
			first = false
		}
		z01.PrintRune(']')
		z01.PrintRune(' ')
	}
	z01.PrintRune('\n')
}

func putnbr(n int) {
	if n < 0 {
		z01.PrintRune('-')
		n = -n
	}
	if n >= 10 {
		putnbr(n / 10)
	}
	z01.PrintRune(rune('0' + n%10))
}

func main() {
	Chunk([]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}, 3)
}
