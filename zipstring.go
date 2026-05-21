package main

import (
    "fmt"
)

func ZipString(s string) string {
    if len(s) == 0 {
        return ""
    }

    result := ""
    count := 1

    for i := 1; i <= len(s); i++ {
        // trigger flush when reaching end OR character changes
        if i == len(s) || s[i] != s[i-1] {
            tmp := count
            digits := ""

            if tmp == 0 {
                digits = "0"
            } else {
                for tmp > 0 {
                    digits = string(rune('0'+(tmp%10))) + digits
                    tmp /= 10
                }
            }

            result += digits
            result += string(s[i-1])

            count = 1
        } else {
            count++
        }
    }

    return result
}

func main()  {
	fmt.Println(ZipString("aaaaaaaaaaaaaaaqqwedeeetrtrrgghghhyy"))
}