package main

import (
	"fmt"
	"time"
)

func sendData(sendChan chan <- int )  {
	fmt.Println(sendChan)
	sendChan <- 10
	//<-sendChan
}

func producer(chnl chan <- int) {
	for i := 0; i < 10; i++ {
		chnl <- i
	}
	close(chnl)
}

func testchan(test chan int)  {
	test <- 10
}

func main() {
	cha0 := make(chan int)

	go testchan(cha0)

	fmt.Println(<-cha0)


	cha1 := make(chan int)
	go sendData(cha1)
	fmt.Println(cha1)
	fmt.Println(<-cha1)

	ch := make(chan  int)
	go producer(ch )
	for {
		v, ok := <-ch
		if ok == false {
			break
		}
		fmt.Println("Received ", v, ok)
	}

	//求乘法  和  加法
	num := 12345678
	chan2 := make(chan int)
	chan3 := make(chan int)
	go chanc(num,chan2)
	go chana(num,chan3)

	data2,data3 := <-chan2,<-chan3
	fmt.Println("chan2为",data2)
	fmt.Println("chan3为",data3)
	fmt.Println(data2+data3)

}


func chanc(number int ,cint chan int)  {
	fmt.Println("c您输入的数字是",number)
	sum :=0

	for number !=0{
		yu := number % 10
		sum += yu*yu
		number /= 10
		fmt.Printf("做乘法-您输入的number是%d,sum是%d \n",number,sum)
		time.Sleep(1*100)  //发现channel是并发的，如果某个信道阻塞其他channel不会停止，程序逻辑会按照时间线线输出
	}
	cint <- sum
}

func chana(number int ,cint chan int)  {
	fmt.Println("a您输入的数字是",number)
	sum :=0

	for number !=0{
		yu := number % 10
		sum += yu+yu
		number /= 10
		fmt.Printf("做加法-您输入的number是%d,sum是%d \n",number,sum)

	}
	cint <- sum

}
