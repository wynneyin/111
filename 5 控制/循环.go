package main

import (
	"fmt"
)

func main() {
	var str string
	str = "G"
	for i := 0; i < 5; i++ {

		fmt.Print("这是", str, "\n")
		str += "G"
	}
	var i int = 5
	for i > 0 {
		fmt.Print(i)
		i = i - 1
	}

LABEL1:
	for i := 0; i <= 5; i++ {
		for j := 0; j <= 5; j++ {
			if j == 4 {
				continue LABEL1
			}
			fmt.Printf("i is: %d, and j is: %d\n", i, j)
		}
	}
}
