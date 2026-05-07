package main

import "fmt"

func RetainLastHalf(s string) string {
	if len(s) == 0 {
		return " "
	}
	if len(s) == 1 {
		return s
	}
	half := len(s) / 2
	return s[half:]
}

func main() {
	fmt.Println(RetainLastHalf("grace"))
	fmt.Println(RetainLastHalf("Clay"))
	fmt.Println(RetainLastHalf(" "))
}
