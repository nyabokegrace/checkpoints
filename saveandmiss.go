package main

import "fmt"

func SaveAndMiss(arg string, num int) string {
	if num <= 0 {
		return arg
	}
	result := ""
	for i := 0; i < len(arg); i++ {
		if (i/num)%2 == 0 {
			result += string(arg[i])
		}
	}
	return result
}

func main() {
	fmt.Println(SaveAndMiss("helloandhowareyou", 2))
	fmt.Println(SaveAndMiss("1234567890", 3))
}
