// takes a slice and an integer as arguments then divides the slice into groups equal to the integer provided.
package main

import (
	"github.com/01-edu/z01"
)

func Chunk(slice []int, size int) {
	if size == 0 {
		z01.PrintRune('\n')
	}
	for i := 0; i < len(slice); i += size { //loops through the slice jumping forward based on the size
		end := i + size       //set the end of the currect chunk
		if end > len(slice) { //if current end goes past the slice end return it forcefully at len(slice)
			end = len(slice)
		}
		z01.PrintRune('[')
		for j := i; j < end; j++ { //loops through the elements of the chunk only from i to the end-1
			putnbr(slice[j])
			if j < end-1 {
				z01.PrintRune(' ')
			}
		}
		z01.PrintRune(']')
	}
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
