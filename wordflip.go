package main

import "fmt"

func wordflip(s string) string {
	w := []string{}
	t := ""

	for _, c := range s {
		if c != ' ' {
			t += string(c)
		} else if t != "" {
			w = append(w, t)
			t = ""
		}
	}
	if t != "" {
		w = append(w, t)
	}
	if len(w) == 0 {
		return "invalid input" + "\n"
	}
	r := ""
	for i := len(w) - 1; i >= 0; i-- {
		r += w[i]

		if i > 0 {
			r += " "
		}
	}
	return r + "\n"
}

func main() {
	fmt.Println(wordflip("My Name is non of your business"))
	fmt.Println(wordflip("  My Name is non    of your business "))
	fmt.Println(wordflip(""))
	fmt.Println(wordflip("	"))
} 


//flips a string to print in reverse but the words are in still order

// initialize w(stores a slice of string) and t(build each word caharacter by character)
//loop  through the string s character by character and if not a space, add it to t to create a string
//if a space is found, append t to w(the word to the slice)
//also if t is an empty string(spaces) they are not added (if t!= "" append(w, t))  
//if is an empty slice, return invalid
// initialize result that reverses the word in a for loop (r += w[i])
//if i > 0 (adds spaces except after the last word) , r= r+ " "
// return the whole slice r + "\n"