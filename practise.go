package main

import "os"
import "math"
import "fmt"

func main(){
	argu := os.Args[1]
	l := len(argu)
	println(l)
	R := 0

	if islower(argu){
		R+=26
	}
	if isUpper(argu){
		R+=26
	}
	if isDigit(argu){
		R+=10
	}
	if isSpecial(argu){
		R+=32
	}
	println(R)
	e := float64(l) * math.Log2(float64(R))
	fmt.Printf("%.2f",(e))
}


func islower(str string)bool{
	for _, char := range str{
		if char >= 'a' && char <= 'z'{
			return true
		}

	}
	return false
}

func isUpper(str string)bool{
	for _, char := range str{
		if char >= 'A' && char <= 'Z'{
			return true
		}
	}
	return false
}

func isDigit(str string)bool{
	for _, char := range str{
		if char >= '0' && char <= '9'{
			return true
		}
	}
	return false
}

func isSpecial(str string)bool{
	for _,char := range str{
		if !(( char >= 'a' && char <= 'z') || (char >= 'A' && char <= 'Z') || (char >= '0' && char <= '9' )){
			return true
		}
	}
	return false
}