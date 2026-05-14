package main

import "fmt"

func Itoa(a int) string{
	if len(a)==0 {
		return "0"
	}
	sign := ""
	if a<0 {
		sign = "-"
		a = -a
	}
	s := ""
for a> 0 {
	s=string(rune('0'+a%10)) + s
	a /= 10
}
return sign + s
}