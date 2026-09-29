package main

import (
	"fmt"
)

func ZipString(s string) string {
	if s == "" {
		return ""
	}
	result := ""
	counter := 1
	for i := 0; i < len(s)-1; i++ {
		if s[i] == s[i+1] {
			counter++
		} else {
			if counter >= 10 {
				result += string(rune(counter/10+'0')) + string(rune(counter%10+'0'))
			} else {
				result += string(rune(counter + '0'))
			}
			result += string(s[i])
			counter = 1
		}
	}
	if counter >= 10 {
		result += string(rune(counter/10+'0')) + string(rune(counter%10+'0'))
	} else {
		result += string(rune(counter + '0'))
	}
	result += string(s[len(s)-1])
	return result
}

func main() {
	fmt.Println(ZipString("YouuungFellllas"))
	fmt.Println(ZipString("Thee quuick browwn fox juumps over the laaazy dog"))
	fmt.Println(ZipString("Helloo Therre!"))
}
