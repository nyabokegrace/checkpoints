package main

import "fmt"

func isPrime(b int) bool {
	if b < 2 {
		return false
	}
	for i := 2; i*i < b; i++ {
		if b%i == 0 {
			return false
		}
	}
	return true
}

func findPrevPrime(a int) int {
	if a < 2 {
		return 0
	}
	for a >= 2 {
		if isPrime(a) {
			return a
		}
		a--
	}
	return 0
}

func main() {
	fmt.Println(findPrevPrime(20))
}
