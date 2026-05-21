package main

import (
	"fmt"
	"os"
) //for piscine import os only//

func main() {
	if len(os.Args) != 4 {
		return
	}
//for piscine initialize result:=""//
	str, old, new := os.Args[1], os.Args[2][0], os.Args[3][0]
	for i := 0; i < len(str); i++ {
		if str[i] == old {
			fmt.Printf("%c", new) //return result+=string(new)//
		} else {
			fmt.Printf("%c", str[i])//return result+=string(s[i])//
		}
	}
	fmt.Println()
}
