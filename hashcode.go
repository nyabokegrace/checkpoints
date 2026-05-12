package main

import "fmt"

func hashCode(s string) string {
	result := ""
	n := len(s)

	for i:=0 ; i < n; i++ {
		code := (int(s[i]) + n) %127
		if code < 32 {
			code += 33
		}
		result += string(rune(code))
	}
	return result
}

func main() {
	fmt.Println(hashCode("Hello"))
}