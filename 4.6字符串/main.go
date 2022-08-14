package main

import (
	"fmt"
	"strings"
)

func main() {
	var str1 string
	str1 = "这a世界有那么多人"
	var str2 string = str1 + "多幸运有个我们"
	str2 += str1
	fmt.Print(str2, "长度为", len(str2), "\n")
	var str3 string = "acd"
	fmt.Println("第2个字符", str3[1])
	fmt.Println("hello, world")

	//4.7 strings 和 strconv 包
	var str string = "This is an example of a string"
	fmt.Printf("T/F? Does the string \"%s\" have prefix %s? ", str, "Th")
	fmt.Printf("%t\n", strings.HasPrefix(str, "Th"))

}
