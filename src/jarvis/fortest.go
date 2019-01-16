package main

import "fmt"

func main() {
	intarr0 := []float64{67.7, 89.8, 21, 78}
	intarr :=append(intarr0, 200)

	for i:=0;i<len(intarr);i++ {
		fmt.Println(intarr[i])
	}
	fmt.Println("\n")

	sum := float64(0);
	for i, v:= range intarr {
		fmt.Printf("%d the element is %.2f\n",i,v)
		sum += v
	}
	fmt.Println(sum)

	a0 := [][]string{
		{"name","age"},
		{"school","live"},
		{"father","monther","sister","brother"},
		{"18"},
	}

	a := append(a0[0],"boy")

	fmt.Println("")
	fmt.Println("a double arr is ",a)
	for i,v:=range a  {
		for j,v1:=range v{
			fmt.Printf("a[%d][%d] is %s \n",i,j,v1)
		}
		fmt.Print("\n")
	}
	fmt.Println("切片")
	newa := a[0:2]
	fmt.Println(newa)
	fmt.Println(len(newa))

	veggies := []string{"potatoes", "tomatoes", "brinjal"}
	fruits := []string{"oranges", "apples"}
	food := append(veggies, fruits...)
	fmt.Println("food:",food)

	sumfloat(1.2,34,5,333223,23,232)

	fmt.Println()
	printCharsAndBytes("eñ%@#@")

	fmt.Println()

	strA:= "hello"
	fmt.Println(mutate([]rune(strA)))

}

func mutate(str []rune) string {
	str[0]= 'p'
	return string(str);
}

func printCharsAndBytes(s string)  {

	for index, rune := range s {
		fmt.Printf("%c starts at byte %d\n", rune, index)
	}

}

func sumfloat(intarr ...float64)  {
	fmt.Println(intarr)
	fmt.Println(intarr[1:len(intarr)-2])
	fmt.Println(cap(intarr[1:len(intarr)-2]))
	sum :=float64(0)
	for _,v:=range intarr {
		sum += v
	}
	fmt.Println(sum)
}
