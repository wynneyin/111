package main

import (
	"fmt"
)

func main() {
	var i1 = 5
	fmt.Printf("An integer: %d, its location in memory: %p\n", i1, &i1)
	var intP *int
	intP = &i1
	*intP = 6
	fmt.Printf("The value at memory location %p is %d\n", intP, *intP)
	fmt.Printf("An integer: %d, its location in memory: %p\n", i1, &i1)
}

//你不能获取字面量或常量的地址，
// const i = 5
// ptr := &i //error: cannot take the address of i
// ptr2 := &10 //error: cannot take the address of 10

//空指针反向引用不可取
//    var p *int = nil
//*p = 0
