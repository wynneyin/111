package main

import "fmt"

var num int = 10
var num_2, num_3 int

func main() {
	num_2, num_3 = getX2AndX3(5)
	//fmt.Printf(" ", num_2, num_3, "\n")
	PrintValues()
	num_2, num_3 = getX2AndX3_2(6)
	//fmt.Printf(" ", num_2, num_3)
	PrintValues()

}

func getX2AndX3(input int) (int, int) {
	return 2 * input, 3 * input
}

//命名返回值
func getX2AndX3_2(input int) (x2 int, x3 int) {
	x2 = 2 * input
	x3 = 3 * input
	return
}

func PrintValues() {
	fmt.Printf("num = %d, 2 num = %d, 3 num = %d\n", num, num_2, num_3)
}
