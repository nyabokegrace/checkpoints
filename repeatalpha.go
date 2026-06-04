package main

import "fmt"

func repeatAlpha(s string) string {
	result := ""
	count := 1
	for _, c := range s {
		if c >= 'a' && c <= 'z' {
			count = (int(c - 'a' + 1))
		}
		if c >= 'A' && c <= 'Z' {
			count = (int(c - 'A' + 1))
		}
		for i := 0; i < count; i++ {
			result += string(c)
		}
	}
	return result
}

func main() {
	fmt.Println(repeatAlpha("abcdefghijklmnopqrstuvwxyz"))
}
