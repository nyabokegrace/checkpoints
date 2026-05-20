package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args) != 4 {
		return
	}

	str, old, new := os.Args[1], os.Args[2][0], os.Args[3][0]
	for i := 0; i < len(str); i++ {
		if str[i] == old {
			fmt.Printf("%c", new)
		} else {
			fmt.Printf("%c", str[i])
		}
	}
	fmt.Println()
}
