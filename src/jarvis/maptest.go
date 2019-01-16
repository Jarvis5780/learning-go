package main

import "fmt"

func main() {
	userinfo := map[string]int{
		"jarvis":1000,
		"stark":203,
		"tony":10023,
	}
	userinfo["sum"] = 2321
	delete(userinfo,"tony")
	for i,v:=range userinfo {
		fmt.Printf("key to value %s=>%d\n",i,v)
	}
}
