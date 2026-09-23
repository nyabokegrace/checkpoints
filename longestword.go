package main

import "fmt"

func LongestWord(arr string) string {
	longest := ""
	current := ""
	hasDigit := false // Tracks if the current word has any numbers

	for i := 0; i < len(arr); i++ {
		if arr[i] != ' ' && arr[i] != '\t' && arr[i] != '\n' {
			// If the character is a digit, flag it
			if arr[i] >= '0' && arr[i] <= '9' {
				hasDigit = true
			}
			current += string(arr[i])
		} else {
			// Only check the length if the word has NO digits
			if !hasDigit && len(current) > len(longest) {
				longest = current
			}
			current = "" 
			hasDigit = false // Reset the flag for the next word
		}
	}

	// Check the final word remaining after the loop ends
	if !hasDigit && len(current) > len(longest) {
		longest = current
	}

	return longest + "\n"
}

func main() {
	fmt.Println(LongestWord("Hello i am a girl"))               // Expected: "Hello"
	fmt.Println(LongestWord("Hel2lo i am a girl"))             // Expected: "girl" ("Hel2lo" is ignored)
	fmt.Println(LongestWord("Hello i am a girrrrrrrrrl"))       // Expected: "girrrrrrrrrl"
	fmt.Println(LongestWord("Hello i am a umemen2yanawekoce"))   // Expected: "umemenyanawekoce"
	fmt.Println(LongestWord("123 123uijjj42"))                  // Expected: "" (Both have digits)
}
