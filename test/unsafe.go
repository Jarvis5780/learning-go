package main

import (
	"fmt"
	"unsafe"
)

const name1 string = "Jarvia"
const INTVAL = 21121212

var age int = 1
var c1, c2, c3 chan int
var i1, i2 int = -2, 3

func main() {
	println(unsafe.Sizeof(INTVAL))
	println(len(name1))
	if age >= 18 {
		print("可以干坏事了")
	} else {
		print("还不行")
	}

	select {
	case i1 = <-c1:
		fmt.Printf("received ", i1, " from c1\n")
	case c2 <- i2:
		fmt.Printf("sent ", i2, " to c2\n")
	case i3, ok := (<-c3): // same as: i3, ok := <-c3
		if ok {
			fmt.Printf("received ", i3, " from c3\n")
		} else {
			fmt.Printf("c3 is closed\n")
		}
	default:
		fmt.Printf("no communication\n")
	}
}
