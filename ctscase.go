package main

import "fmt"

func CamelToSnakeCase(s string) string {

	n := len(s)
	result := ""

	for i, v := range s {
		last := s[n-1]

		if len(s) == 0 {
			return s
		}
		if !(v >= 'a' && v <= 'z') && !(v >= 'A' && v <= 'Z') {
			return s
		}
		if last >= 'A' && last <= 'Z' {
			return s
		}
		if i > 0 && (v >= 'A' && v <= 'Z') && (s[i-1] >= 'A' && s[i-1] <= 'Z') {
			return s
		}
		if i > 0 && (v >= 'A' && v <= 'Z') {
			result += "_"
		}
		result += string(v)
	}
	return result
}

func main() {
	fmt.Println(CamelToSnakeCase("CamelCase"))
}
