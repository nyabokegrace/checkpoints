package main

import "fmt"

<<<<<<< HEAD
func CamelToSnakeCase(s string) string {

	n := len(s)
	result := ""

	for i, v := range s {
		last := s[n-1]

		if len(s) == 0 {
			return s
		}
		if !(v >= 'a' && v <= 'z') && !(v >= 'A' && v <= 'Z') {
			return s
		}
		if last >= 'A' && last <= 'Z' {
			return s
		}
		if i > 0 && (v >= 'A' && v <= 'Z') && (s[i-1] >= 'A' && s[i-1] <= 'Z') {
			return s
		}
		if i > 0 && (v >= 'A' && v <= 'Z') {
			result += "_"
		}
		result += string(v)
=======
func CamelToSnakeCase(s string) string{
	result := ""
	n:=len(s)

	for i , v := range s {
	last := s[n-1]

	//checking for invalidity of characters in camel case//

	if s==""{
		return ""
	}
	if !(v>='a'&& v <= 'z')&& !( v >= 'A'&& v <= 'Z') {
		return s
	}
	if i>0 && ( v >='A' && v <= 'Z') && (s[i-1] >='A' && s[i-1]<= 'Z'){
		return s
	}
	if last >= 'A' && last <='Z' {
		return s
	}
	if  i>0 && v >= 'A'&& v <= 'Z' {
		result += "_"
	}
	result+=string(v)
>>>>>>> 5098bce (changes)
	}
	return result
}

func main() {
	fmt.Println(CamelToSnakeCase("CamelCase"))
<<<<<<< HEAD
}
=======
	fmt.Println(CamelToSnakeCase("CamelCasE"))
	fmt.Println(CamelToSnakeCase("CamelCa#e"))
	fmt.Println(CamelToSnakeCase("CAmelCase"))
	fmt.Println(CamelToSnakeCase("CamelCas2"))
	fmt.Println(CamelToSnakeCase("camelCase"))
}
>>>>>>> 5098bce (changes)
