package main

import (
	"fmt"
	"os/exec"
)

func main() {

	//for index, val := range os.Environ() {
	//	fmt.Print(index, val, "\n")
	//}
	//version := os.Getenv("VERSION")
	//fmt.Print(version, "\n")
	//os.Setenv("FOO", "1")
	//fmt.Println("FOO:", os.Getenv("FOO"))
	//fmt.Println("BAR:", os.Getenv("BAR"))
	//
	//fmt.Println()
	//for _, e := range os.Environ() {
	//	pair := strings.SplitN(e, "=", 2)
	//	fmt.Println(pair[0])
	//}
	out, _ := exec.Command("uname", "-r").Output()
	//output := string(out[:])
	fmt.Printf("%T\n", out)
	fmt.Print(string(out))
}
