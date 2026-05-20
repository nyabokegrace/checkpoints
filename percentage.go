package main

import "fmt"

func PercentOf(part , total float32) float32 {
	if total == 0 {
		return 0 
	}
	return (part / total) * 100
}

func main() {
	fmt.Println(PercentOf(50, 200))	
}
