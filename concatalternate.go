package main

import "fmt"

func ConcatAlternate(slice1, slice2 []int) []int {
	i, j := 0, 0
	res := make([]int, 0, len(slice1)+len(slice2))

	for i < len(slice1) && j < len(slice2) {
		if len(slice1) >= len(slice2) {
			res = append(res, slice1[i], slice2[j])
		} else {
			res = append(res, slice2[j], slice1[i])
		}
		i++
		j++
	}
	for i < len(slice1) {
		res = append(res, slice1[i])
		i++
	}

	for j < len(slice2) {
		res = append(res, slice2[j])
		j++
	}
	return res
}

func main() {
	fmt.Println(ConcatAlternate([]int{1, 2, 3}, []int{4, 5, 6}))
	fmt.Println(ConcatAlternate([]int{2, 3, 4, 5, 6}, []int{7, 8, 9}))
	fmt.Println(ConcatAlternate([]int{1, 2, 3}, []int{}))
}
