package main
import "fmt"
func ConcatSlice(slice1, slice2 []int) []int {
	res:= make([]int, len(slice1)+ len(slice2))
	for i:=0; i<len(slice1); i++{
		res[i]=slice1[i]
	}
	for i:=0; i<len(slice2); i++{
		res[len(slice1)+i]= slice2[i]
	}
	return res
}


//shorter version

//func ConcatSlice(slice1, slice2 []int) []int {
//	return append(slice1,slice2...)
//}


func main()  {
	fmt.Println(ConcatSlice([]int{1, 2, 3}, []int{4, 5, 6}))
}