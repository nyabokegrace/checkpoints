package main

import "fmt"

func contains(s string, c rune) bool {
	for _, v := range s {
		if v == c {
			return true
		}
	}
	return false
}

func weAreUnique(str1, str2 string) int {
	if str1 == "" && str2 == "" {
		return -1
	}
	unique := ""

	for _, c := range str1 {
		if !contains(str2, c) && !contains(unique, c) {
			unique = unique + string(c)
		}
	}
	for _, c := range str2 {
		if !contains(str1, c) && !contains(unique, c) {
			unique = unique + string(c)
		}
	}
	return len(unique)
}

func main() {
	fmt.Println(weAreUnique("clay", "grace"))
	fmt.Println(weAreUnique("ring", "caakes"))

}
