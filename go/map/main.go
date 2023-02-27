package main

import(
	"fmt"
)


func main() {
	var m map[string] string
	m = make(map[string] string) 
	m["k1"] = "v1"
	var k string
	k = "k1"
	fmt.Printf("k = %v, v = %v \n", k, m[k])
	k = "k2"
	if m[k] == "" {
		fmt.Printf("k = %v, v = %v \n", k, m[k])
	}
}
