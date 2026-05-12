package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args) != 4 {
		return
	}
	str, old, new := os.Args[1], os.Args[2], os.Args[3]
	for i := 0; i < len(str); i++ {
		if i< (len(old) && str[i] == old[0]) && len(old) == 1 {
			fmt.Println((new[0]))
		} else {
			fmt.Println((str[i]))
		}
	}
	fmt.Println("\n")
}
