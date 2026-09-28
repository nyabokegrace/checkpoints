package main

import "fmt"

func IsCapitalized(s string) bool {
	if s == "" {
		return false
	}
	if s[0] >= 'a' && s[0] <= 'z' {
		return false
	}
	for i := 1; i < len(s); i++ {
		if s[i-1] == ' ' && s[i] >= 'a' && s[i] <= 'z' {
			return false
		}
	}
	return true
}
func main() {
	fmt.Println(IsCapitalized("Hello How Are You"))
	fmt.Println(IsCapitalized("100k #jjhjd"))
}
