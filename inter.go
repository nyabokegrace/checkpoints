package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args) != 3 {
		return
	}
	s1,s2 := os.Args[1],os.Args[2]
	seen:= make(map[rune]bool)

	for _, ch := range s1 {
		if !seen[ch]{
			for _,ch1:= range s2 {
				if ch==ch1 {
					fmt.Print(string(rune(ch)))
					seen[ch]=true
					break
				}
			}
		}
	}
	fmt.Println()
}