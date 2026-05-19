<<<<<<< HEAD
package main
=======
package main 
>>>>>>> 5098bce (changes)

import "fmt"

func CamelToSnakeCase(s string) string {
<<<<<<< HEAD
	if s == "" {
		return s
	}
	result := ""
	n := len(s)

	last := s[n-1]

	for i, v := range s {
		if last >= 'A' && last <= 'Z' {
			return s
		}
		if i > 0 && !(v >= 'a' && v <= 'z') && !(v >= 'A' && v <= 'Z') {
			return s
		}
		if i > 0 && (v >= 'A' && v <= 'Z') && (s[i-1] >= 'A' && s[i-1] <= 'Z') {
			return s
		}
		if i>0 && v >= 'A' && v <= 'Z' {
			result += "_"
		}
		result= result + string(v)
	}
	return result
}

func main() {
	fmt.Println(CamelToSnakeCase("CamelCase"))
}
=======
	result:= ""
	n :=len(s)

	if s == "" {
		return s
	}
	
	for i,v := range s
}
>>>>>>> 5098bce (changes)
