package main

import (
	"fmt"
)

func PrintMemory(arr [10]byte) {
	// Print hexadecimal representation
	for i, b := range arr {
		fmt.Printf("%02x", b)

		if (i+1)%4 == 0 {
			fmt.Print("\n")
		} else {
			fmt.Print(" ")
		}
	}

	// If last line is not complete
	if len(arr)%4 != 0 {
		fmt.Println()
	}

	// Print ASCII representation
	for _, b := range arr {
		if b >= 32 && b <= 126 {
			fmt.Printf("%c", b)
		} else {
			fmt.Print(".")
		}
	}
	fmt.Println()
}

func main() {
	PrintMemory([10]byte{'h', 'e', 'l', 'l', 'o', 16, 21, '*'})
}
