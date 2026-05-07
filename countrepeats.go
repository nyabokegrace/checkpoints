package main 

import(
	"strconv"
	 
	"fmt"
)

func CountRepeats(s string) string {
	 if s == "" {
		return ""
	 }
	 var result string
	 var counter = 1

	 for i := 0; i < len(s); i++ {
		if i+1 < len(s) && string(s[i]) == string(s[i+1]) {
			counter++
		} else {
			result = result + string (s[i])
			if counter > 1 {
				result = result + strconv.Itoa(counter)
			}
			counter = 1
		}
	 }
return result
}

func main() {
	fmt.Println(CountRepeats("aaabbbccc"))
	fmt.Println(CountRepeats("abc"))
	fmt.Println(CountRepeats("abbccc"))
}