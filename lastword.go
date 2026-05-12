package main

import "fmt"

func lastWord(s string) string {
	word := ""
	last := ""

	for i := 0; i < len(s); i++ {
		if s[i] != ' ' {
			word += string(s[i])
		} else {
			if word != "" {
				last = word
			}
			word = ""
		}
	}
	if word != "" {
		last = word
	}
	return last + "\n"
}

func main() {
	fmt.Println(lastWord("Hello world"))
}
