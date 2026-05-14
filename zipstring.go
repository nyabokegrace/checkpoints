package main

import "fmt"

func ZipString(s string) string {
	if s == "" {
		return ""
	}
	result := ""
	count := 1
	for i := 1; i < len(s); i++ {
		if s[i] == s[i-1] {
			count++
		} else {
			result += string(rune(count+'0')) + string(s[i-1])
			count = 1
		}
	}
	result += string(rune(count+'0')) + string(s[len(s)-1])
	return result
}

func main() {
	fmt.Println(ZipString("aaabbc"))
	fmt.Println(ZipString("grraaaceee"))
	fmt.Println(ZipString(""))
	fmt.Println(ZipString("clay"))
	fmt.Println(ZipString("Beckham"))
}
