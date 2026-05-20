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
