package main

import "fmt"

func FifthAndSkip(str string) string {
	if str == "" {
		return "\n"
	}

	count := 0

	// Count non-space chars
	for _, ch := range str {
		if ch != ' ' {
			count++
		}
	}

	if count < 5 {
		return "Invalid Input\n"
	}

	result := ""
	kept := 0
	total := 0

	for _, ch := range str {
		if ch == ' ' {
			continue
		}

		total++

		// Skip every 6th character
		if total%6 == 0 {
			continue
		}

		// Add space after every 5 kept chars
		if kept > 0 && kept%5 == 0 {
			result += " "
		}

		result += string(ch)
		kept++
	}

	return result + "\n"
}

func main() {
	fmt.Print(FifthAndSkip("abcdefghijklmnopqrstuwxyz"))
	fmt.Print(FifthAndSkip("This is a short sentence"))
	fmt.Print(FifthAndSkip("1234"))
}