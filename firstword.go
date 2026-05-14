package main

import "fmt"

func FirstWord(s string) string {
	if len(s) == 0 {
		return "invalid"
	}
	result := ""
	i := 0

	for i < len(s) && s[i] == ' ' {
		i++
	}
	for i < len(s) && s[i] != ' ' {
		result = result + string(s[i])
		i++
	}
	return result + "\n"
}

func main() {
	fmt.Println(FirstWord("Hello world"))
	fmt.Println(FirstWord(""))

}
