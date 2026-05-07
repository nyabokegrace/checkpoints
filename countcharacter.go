package main

import "fmt"

func CountCharacter(s string) int {
	count := 0
	for _, ch := range s {
		if (ch >= 'a' && ch <= 'z') || (ch >= 'A' && ch <= 'Z' || ch >= '0' && ch <= '9') {
			count++
		}
	}
	return count
}

func main() {
	fmt.Println(CountCharacter("AEiOu78"))
	fmt.Println(CountCharacter("grace"))
}
