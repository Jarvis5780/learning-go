package main

import (
	"time"
	"fmt"
)

func main() {
	//ch := make(chan string)
	var ch chan string
	go process(ch)
	for{
		time.Sleep(1000*time.Microsecond);
		select {
		case v:=<-ch:
			fmt.Println("程序结束，输出",v)
			return
		default:
			fmt.Println("no value")

		}
	}
}

func process(ch1 chan string)  {
	time.Sleep(10500*time.Microsecond)
	ch1 <- "from process"
}
