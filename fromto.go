package main

import "fmt"

func FromTo(from int, to int) string {
	if from < 0 || from > 99 || to < 0 || to > 99 {
		return "invalid\n"
	}
	result := ""
	if from <= to {
		for i := from; i <= to; i++ {
			result += string(rune('0' + i/10))
			result += string(rune('0' + i%10))

			if i != to {
				result += ", "
			}
		}
	} else {
		for i := from; i >= to; i-- {
			result += string(rune('0' + i/10))
			result += string(rune('0' + i%10))

			if i != to {
				result += ", "
			}
		}
	}
	return result + "\n"
}

func main() {
	fmt.Println(FromTo(10, 20))
}
