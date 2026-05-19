package main

import "fmt"

func Slice(a []string, nbrs ...int) []string {
	n := len(a)
	if n == 0 || len(nbrs) == 0 {
		return nil
	}
	adjust := func(i int) int {
		if i < 0 {
			i += n
		}
		if i < 0 {
			return 0
		}
		if i > n {
			return n
		}
		return i
	}
	start := adjust(nbrs[0])
	end := n
	if len(nbrs) > 1 {
		end = adjust(nbrs[1])
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
