package main

import(
	"fmt"
	"strconv"
)



func main () {
	var i int = 1
	fmt.Printf("str(i) = %s\n", strconv.Itoa(i))
	i = -1
	fmt.Printf("str(i) = %s\n", strconv.Itoa(i))
}
