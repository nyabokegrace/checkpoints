package main

import "fmt"

func RevConcatAlternate(slice1, slice2 []int) []int {
	i, j := len(slice1)-1, len(slice2)-1
	res := make([]int, 0, len(slice1)+len(slice2))

	for i >= 0 && j >= 0 {
		if len(slice1) >= len(slice2) {
			res = append(res, slice1[i], slice2[j])
		} else {
			res = append(res, slice2[j], slice1[i])
		}
		i--
		j--
	}
	for i >= 0 {
		res = append(res, slice1[i])
		i--
	}

	for j >= 0 {
		res = append(res, slice2[j])
		j--
	}
	return res
}

func main() {
	fmt.Println(RevConcatAlternate([]int{1, 2, 3}, []int{4, 5, 6}))
	fmt.Println(RevConcatAlternate([]int{2, 3, 4, 5, 6}, []int{7, 8, 9}))
	fmt.Println(RevConcatAlternate([]int{1, 2, 3}, []int{}))
}
