package main

import "fmt"

func FirstWord(s string) string {
	if len(s) == 0 {
		return "invalid"
	}
	result := ""
	i := 0

	for i < len(s) && s[i] == ' ' {
		i++ //loop through s even after encountering a space
	}
	for i < len(s) && s[i] != ' ' {
		result = result + string(s[i])
		i++ // loop through s  and if s[i] is not a space add result (that is the end of the first word)
	}
	return result + "\n"
}
//return the result
func main() {
	fmt.Println(FirstWord("Hello world"))
	fmt.Println(FirstWord(""))

}
