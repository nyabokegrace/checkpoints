package main

import "fmt"

func NotDecimal(dec string) string {
	if dec == "" {
		return "\n"
	}
	dot := -1
	for i := 0; i < len(dec); i++ {
		c := dec[i]
		if c == '.' {
			if dot != -1 {
				return dec + "\n"
			}
			dot = i
		} else if c < '0' || c > '9' {
			if i != 0 || c != '-' {
				return dec + "\n"
			}
		}
	}

	if dot == -1 {
		return dec + "\n"
	}

	res := ""
	started:= false
	for i := 0; i < len(dec); i++ {
		if i != dot {
			res += string(dec[i])
		}
	}
	end := len(res) - 1
	for end > dot && res[end] == '0' {
		end--
	}
	if end == dot {
		end--
	}

	if end < 0 {
		return dec + "\n"
	}
	return res[:end+1] + "\n"
}

func main() {
	fmt.Print(NotDecimal("0.1"))
	fmt.Print(NotDecimal("174.2"))
	fmt.Print(NotDecimal("0.1255"))
	fmt.Print(NotDecimal("1.20525856"))
	fmt.Print(NotDecimal("-0.0f00d00"))
	fmt.Print(NotDecimal(""))
	fmt.Print(NotDecimal("-19.525856"))
	fmt.Print(NotDecimal("1952"))
}
