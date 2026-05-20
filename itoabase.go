package main

import "fmt"

func ItoaBase(n, base int) string {
	if base < 2 || base > 16 {
		return " "
	}
	sign := " "
	if n < 0 {
		sign = "-"
		n = -n
	}
	if n == 0 {
		return "0"
	}
	result := " "
	for n > 0 {
		result = string(rune('0'+n%base)) + result
		n /= base
	}
	return sign + result
}

func main() {
	fmt.Println(ItoaBase(10, 2))
}
