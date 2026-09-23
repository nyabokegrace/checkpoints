package main

import "fmt"

func ZipString(s string) string {
	res := ""
	for i := 0; i < len(s); i++ {
		count := 1
		for i+1 < len(s) && s[i] == s[i+1] {
			count++
			i++
		}
		res += string(rune(count+'0')) + string(s[i])
	}
	return res
}

func main() {
	fmt.Println(ZipString("aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaabbc"))
	fmt.Println(ZipString("gggggggggggrraaaceee"))
	fmt.Println(ZipString(""))
	fmt.Println(ZipString("clay"))
	fmt.Println(ZipString("Beckham"))
}
