package main

import "fmt"

func PrintIf(arg string) string {
	if len(arg) >= 3 {
		return "G\n"
	}
	return "Invalid input \n"
}

func main() {
	fmt.Println(PrintIf(" "))
}
