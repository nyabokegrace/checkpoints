package main

import (
	"os"
	"fmt"
)

func main(){
	if len(os.Args)!=3 {
	fmt.Println()	
		return
	}
s:= os.Args[1]+ os.Args[2]
seen:= make(map[rune]bool)

for _, c:= range s {
	if !seen [c] {
		fmt.Print(string(rune(c)))
		seen[c]=true
	}
}
fmt.Println()
}