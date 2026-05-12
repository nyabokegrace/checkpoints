package main

import "fmt"

func thirdTimeIsACharm(s string) string{
	result :=""
	
	for i:=2; i<len(s);i+=3 {
		result= result + string(s[i])
	}
	return result
}

func main(){
	fmt.Println(thirdTimeIsACharm("123456789"))
	fmt.Println(thirdTimeIsACharm("grace"))
	fmt.Println(thirdTimeIsACharm(""))
	fmt.Println(thirdTimeIsACharm("clay"))
	fmt.Println(thirdTimeIsACharm("Beckham"))

}