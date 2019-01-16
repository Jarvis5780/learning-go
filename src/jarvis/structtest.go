package main

import (
	"fmt"
)

type one struct {
	firstname string
	lastnamr string
	age int
}

type rule struct{
	school string
	home  string
	word string
}

type user struct{
	username string
	age int
	sex string
	ruleMSg rule
}

type boy struct{
	string
	int
}

type name struct {
	firstName string
	lastName string
}

func main() {
	jarvis := one{
		firstname:"张",
		lastnamr:"三",
		age:18,
	}
	strack :=one{"小","罗伯特",20}


	jarvis.lastnamr = "三宝"

	paper := &one{"small","paper",19}

	fmt.Println("strack",strack)
	fmt.Println("jarvis",jarvis)
	fmt.Println((*paper).firstname)

	var jarvisA user
	jarvisA.username = "jarvisA"
	jarvisA.age = 18
	jarvisA.sex = "boy"
	jarvisA.ruleMSg = rule{"good student","father","good man"}
	fmt.Println(jarvisA)

	boyA := boy{"sex",10}
	fmt.Println(boyA)

	name1 := name{"Steve", "Jobs"}
	name2 := name{"Steve", "Jobs"}
	if name1 == name2 {
		fmt.Println("name1 and name2 are equal")
	} else {
		fmt.Println("name1 and name2 are not equal")
	}

	name3 := name{firstName:"Steve", lastName:"Jobs"}
	name4 := name{}
	name4.firstName = "Steve"
	if name3 == name4 {
		fmt.Println("name3 and name4 are equal")
	} else {
		fmt.Println("name3 and name4 are not equal")
	}

}
