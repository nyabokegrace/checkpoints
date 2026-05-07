package main

import "fmt"

func PrintIfNot(arg string) string {
	if len(arg) <= 3 {
		return "G\n"
	}
	return "Invalid input\n"
}

func main() {
	fmt.Println(PrintIfNot("Grace"))
	fmt.Println(PrintIfNot(" "))
	fmt.Println(PrintIfNot("GRA"))
}
