package main

import "fmt"

func RetainFirstHalf(s string) string {
	if len(s) == 0 {
		return " "
	}
	if len(s) == 1 {
		return s
	}
	half := len(s) / 2
	return s[:half]
}

func main() {
	fmt.Println(RetainFirstHalf("grace"))
	fmt.Println(RetainFirstHalf("Clay"))
	fmt.Println(RetainFirstHalf(" "))
}
