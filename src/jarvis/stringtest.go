package main

import "fmt"

func showeves(stringA string)  {
	for i,v:=range stringA {
		fmt.Printf("当前第%d个字符串为%c\n",i,v);
	}

}
func main() {
	showeves("abcdef")
}
