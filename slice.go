package main

import "fmt"

func Slice(a []string, nbrs ...int) []string {
	if len(a) == 0 || len(nbrs) == 0 {
		return nil
	}
	length := len(a)
	start := nbrs[0]

	if start < 0 {
		start += length //if you said a negative number it means you want me start from the end eg when you say start is -1 it will be (-1 +=length )meaning its the last one so we will be counting backwards.
	}
	end := length      //by default we will pick all the numbers/characters all tjhe way to the end
	if len(nbrs) > 1 { //if you gave me twoi numbers
		end = nbrs[1] // the second number is the end
		if end < 0 {  //if its negative we do it backwards
			end += length
		}
	}
	if start < 0 { //make sure i start in the first car
		start = 0
	}
	if end > length {
		end = length
	}
	if start >= end {
		return nil
	}
	return a[start:end]
}

func main() {
	a := []string{"coding", "algorithm", "ascii", "package", "golang"}
	fmt.Printf("%#v\n", Slice(a, 1))
	fmt.Printf("%#v\n", Slice(a, 2, 10))
	fmt.Printf("%#v\n", Slice(a, -3))
	fmt.Printf("%#v\n", Slice(a, -2, -1))
	fmt.Printf("%#v\n", Slice(a, 2, 0))
}