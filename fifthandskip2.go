package main

import "fmt"

func FifthAndSkip(str string) string {
	clean := ""
	for _, c := range str {
		if c != ' ' {
			clean += string(c)
		}
	}
	
	if len(clean) < 5 {
		if str == "" { return "\n" }
		return "Invalid Input\n"
	}
	
	filtered := ""
	for i := 0; i < len(clean); i++ {
		if i%6 < 5 {
			filtered += string(clean[i])
		}
	}

	result := ""
	for i := 0; i < len(filtered); i++ {
		if i > 0 && i%5 == 0 {
			result += " "
		}
		result += string(filtered[i])
	}
	return result + "\n"
}

func main()  {
	fmt.Println(FifthAndSkip("abcdefghijklmnopqrstuwxyz"))
}
