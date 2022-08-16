package main

import (
	"fmt"
	"runtime"
)

var prompt = "Enter a digit, e.g. 3 " + "or %s to quit."

func demo() {
	fmt.Printf("\ndemo")
}
func main() {
	bool1 := false
	if bool1 {
		fmt.Printf("this value is true\n")
	} else {
		fmt.Printf("the value is false\n")
	}
	if runtime.GOOS == "windows" {
		fmt.Printf("这个是windows")
	}
	demo()
}

func init() {
	if runtime.GOOS == "windows" {
		fmt.Printf("hello\n")
		prompt = fmt.Sprintf(prompt, "Ctrl+Z, Enter")
	} else { //Unix-like
		prompt = fmt.Sprintf(prompt, "Ctrl+D")
	}
}
