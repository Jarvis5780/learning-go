package main

import (
	"fmt"
	"geometry/rectangle"
	"log"
	"math"
)

var rectLen, rectWidth float64 = 6, 17
var _=math.Max(1,2)

func init() {
	println("main package initialized")
	if rectLen < 0 {
		log.Fatal("length is less than zero")
	}
	if rectWidth < 0 {
		log.Fatal("width is less than zero")
	}
}

func main() {
	fmt.Println("Geometrical shape properties")
	/*Area function of rectangle package used*/
	fmt.Println("area of rectangle %.2f\n", rectangle.Area(rectLen, rectWidth))
	/*Diagonal function of rectangle package used*/
	fmt.Println("diagonal of the rectangle %.2f ", rectangle.Diagonal(rectLen, rectWidth))

	rectangle.Show()
	if rectLen>10 {
		fmt.Println("rectLen > 10")
	}else if rectLen>5 && rectLen<=10 {
		fmt.Println("rectLen less 10")
	}else{
		fmt.Println("less 5")
	}

	for i:=1; i<=10 ;i++  {
		if i==5 {
			continue
		}
		if i>8 {
			break
		}
		fmt.Printf("%d",i)
	}

	fmt.Println("--------------------")


	//5050计算
	sum :=0
	for i:=1;i<=100 ;i++  {
		sum = sum + i
		fmt.Println("当前第",i,"次相加",sum-i,"+",i,"",sum)
	}

	//乘法表
	for i:=1;i<=9 ;i++  {
		for j:=1;j<=i ;j++  {
			fmt.Print(j,"*",i,"=",j*i," ")
		}
		fmt.Println("")

	}
}
