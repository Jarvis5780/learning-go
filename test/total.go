package main

import (
	"fmt"
)

func total(price, num int) (mul, sum, total int) {
	mul = price * num
	sum = num
	total = price + num
	return
}

func main() {
	book := 10
	num := 3
	bookmul, sum, _ := total(book, num)
	fmt.Println("bookmul price total is", bookmul, ",book sum", sum)

}
