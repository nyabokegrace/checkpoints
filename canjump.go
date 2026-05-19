package main

import "fmt"

func CanJump(array []uint) bool {
	if len(array) == 0 {
		return false
	}
	currentIndex := 0
	lastIndex := len(array) - 1

	for currentIndex < lastIndex {
		steps := int(array[currentIndex])

		if steps == 0 {
			return false
		}
		currentIndex += steps
	}
	return currentIndex == lastIndex
}

func main() {
	input1 := []uint{2, 3, 1, 1, 4}
	fmt.Println(CanJump(input1))

	input2 := []uint{3, 2, 1, 0, 4}
	fmt.Println(CanJump(input2))

	input3 := []uint{0, 1, 1, 3}
	fmt.Println(CanJump(input3))
}
