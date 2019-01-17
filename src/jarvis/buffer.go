package main

import (
	"fmt"
	"time"
)

func main() {
	//输入的信道数和输出的信道数应该数量一样，
	buf0 := make(chan int,3)
	buf0 <- 10
	buf0 <- 20
	buf0 <- 30

	fmt.Println("capacity is", cap(buf0))
	fmt.Println("length is", len(buf0))
	fmt.Println("read value", <-buf0)
	fmt.Println("length is", len(buf0))
	fmt.Println("new length is", len(buf0))
	fmt.Println(<-buf0)
	fmt.Println(<-buf0)

	fmt.Println()
	//开辟5个值的缓冲区，go后会立即执行write方法内的输出，，之后走for的时候依次输出ch <- i设置的值，单个输出，但是因为要睡3秒，之后才能输出依次输出，此时按照时间线输出
	ch0 := make(chan int,5)
	go write(ch0)
	time.Sleep(1*time.Second)
	for v := range ch0{
		fmt.Println("当前v=",v)
		time.Sleep(2*time.Second)
		
	}
}

func write(ch chan int) {
	for i := 0; i < 10; i++ {
		ch <- i
		fmt.Println("successfully wrote", i, "to ch")
	}
	close(ch)
}
