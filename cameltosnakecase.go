package main

import "fmt"

func CamelToSnakeCase(s string) string {
	if s == "" {
		return " "
	}
	n := len(s)
	lastChar := s[n-1]
	if (lastChar >= 'A' && lastChar <= 'Z') || (lastChar < 'a' && lastChar > 'z') {
		return s
	}
	result := []byte{}
	for i := 0; i < n; i++ {
		char := s[i]

		if !((char >= 'a' && char <= 'z') || (char >= 'A' && char <= 'Z')) {
			return s
		}
		if i > 0 && char >= 'A' && char <= 'Z' {
			prevChar := s[i-1]
			if prevChar >= 'A' && prevChar <= 'Z' {
				return s
			} else {
				result = append(result, '_')
			}
		}
		if char >= 'A' && char <= 'Z' {
			char = char + 32
		}
		result = append(result, char)
	}
	return string(result)
}

func main() {
	fmt.Println(CamelToSnakeCase("CamelCase"))
	fmt.Println(CamelToSnakeCase("CamelCasE"))
	fmt.Println(CamelToSnakeCase("CamelCa#e"))
	fmt.Println(CamelToSnakeCase("CAmelCase"))
	fmt.Println(CamelToSnakeCase("CamelCas2"))
	fmt.Println(CamelToSnakeCase("camelCase"))
}
