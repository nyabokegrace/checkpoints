package main

import "fmt"

func Itoa(n int) string {
	if n == 0 {
		return "0"
	}

	sign := ""
	if n < 0 {
		sign = "-"
		n = -n
	}

	s := ""
	for n > 0 {
		s = string(rune('0'+n%10)) + s
		n /= 10
	}
	return sign + s
}

func main() {
	fmt.Println(Itoa(12345))
	fmt.Println(Itoa(0))
}
