package main

import "os/exec"
import "fmt"

func main(){
	exec.Command("cd", "checkpoints/")
	cmd := exec.Command("cat", "printmemory.go")
	out, _ := cmd.Output()

	fmt.Println(string(out))
}