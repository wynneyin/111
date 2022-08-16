package main

import (
	"fmt"
)

func main() {

	var num1 int = -1
	//一旦成功地匹配到某个分支，在执行完相应代码后就会退出整个 switch 代码块，也就是说您不需要特别使用 break 语句来表示结束。
	switch {
	case num1 < 0:
		fmt.Printf("num<0")
	case num1 <= 0:
		fmt.Printf("num<=0")

	}
}
