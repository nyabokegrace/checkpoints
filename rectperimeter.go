package main

import "fmt"

func RectPerimeter(w, h int) int {
	if w <= 0 || h <= 0 {
		return -1
	}
	return 2 * (w + h)
}
func main() {
	fmt.Println(RectPerimeter(4, 5))
	fmt.Println(RectPerimeter(0, 2))
}
