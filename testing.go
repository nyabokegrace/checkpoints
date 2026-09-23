package main

import "fmt"

func CountRepeats(s string) string {
	res := ""
	for i := 0; i < len(s); i++ {
		count := 1
		for i+1 < len(s) && s[i] == s[i+1] {
			count++
			i++
		}
		if count < 2 {
			res += string(s[i])
		} else {
			res += string(s[i]) + string(rune(count+'0'))
		}

	}
	return res
}

func main() {
	fmt.Println(ZipString("aaaaaaabbc"))
	fmt.Println(ZipString("gggggggrraaaceee"))
	fmt.Println(ZipString(""))
	fmt.Println(ZipString("clay"))
	fmt.Println(ZipString("Beckham"))
}
