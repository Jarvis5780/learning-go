package main

import (
	"fmt"
	"time"
)

func printInt()  {
	for i:=1;i<=5 ;i++  {
		time.Sleep(40*time.Millisecond)
		fmt.Printf("当前输出printInt函数%d \n",i);
	}
}

func printstr()  {
	for j:='a';j<='e';j++ {
		time.Sleep(50*time.Millisecond)
		fmt.Printf("当前输出printInt函数%c \n",j);

	}
}

func printsIntD()  {
	for i:=100;i>=95 ;i--  {
		time.Sleep(100*time.Millisecond)
		fmt.Printf("当前输出printIntD函数%d \n",i);
	}
}

func main()  {
	go printInt()
	go printstr()
	go printsIntD()
	time.Sleep(3000*time.Millisecond)
	fmt.Println("sleep 3000 over",)
}