package main

import "fmt"

func CheckNumber(arg string) bool {
	for _, v := range arg {
		if v >= '0' && v <= '9' {
			return true
		}
	}
	return false
}

func main() {
	fmt.Println(CheckNumber("9"))
	fmt.Println(CheckNumber("a"))
}
