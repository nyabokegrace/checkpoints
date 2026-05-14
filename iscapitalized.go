package main

import "fmt"

func IsCapitalized(s string) bool {
	if s == "" {
		return false
	}
	result := true
	for _, c := range s {
		if c == ' ' {
			result = true
		} else if result {
			if c >= 'a' && c <= 'z' {
				return false
			}
			result = false
		}
	}
	return true
}

func main() {
	fmt.Println(IsCapitalized("Hello How Are You"))
	fmt.Println(IsCapitalized("100k #jjhjd"))
}
