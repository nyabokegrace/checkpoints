package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Println(0)
	}
	n := atoi(os.Args[1])
	if n < 0 {
		fmt.Println(0)
	}

	sum := 0
	for i := 2; i <= n; i++ {
		if isPrime(i) {
			sum += i
		}
	}
	fmt.Println(sum)
}

func isPrime(n int) bool {
	if n < 2 {
		return false
	}
	for i := 2; i*i < n; i++ {
		if n%i == 0 {
			return false
		}
	}
	return true
}

func atoi(s string) int {
	num := 0
	for i := 0; i < len(s); i++ {
		c := s[i]

		if c < '0' || c > '9' {
			return -1
		}
		num = num*10 + int(c-'0')
	}
	return num
}
