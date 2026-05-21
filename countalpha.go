package main

import "fmt"

func CountAlpha(s string) int {
	count := 0
	for _, ch := range s {
		if (ch >= 'a' && ch <= 'z') || (ch >= 'A' && ch <= 'Z') 
			count++
		}
		return count
	}

func main() {
	fmt.Println(CountAlpha("grac*#e"))
}

//counts special characters too//
