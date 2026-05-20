package main

import "fmt"

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
	}
	return result
}

func main() {
	fmt.Println(CamelToSnakeCase("CamelCase"))
	fmt.Println(CamelToSnakeCase("CamelCasE"))
	fmt.Println(CamelToSnakeCase("CamelCa#e"))
	fmt.Println(CamelToSnakeCase("CAmelCase"))
	fmt.Println(CamelToSnakeCase("CamelCas2"))
	fmt.Println(CamelToSnakeCase("camelCase"))
}
