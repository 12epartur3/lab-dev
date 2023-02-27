package main

import(
	"fmt"
	"strconv"
	"math"
)


func main() {
	//s := make([]string, 0)
	fmt.Printf("MaxFloat32=%f\n", float32(math.MaxFloat32))
	//var f float32 = float32(math.MaxFloat32)
	var f float32
	for {
		fmt.Printf("Maxfloat32=%f\n", f)
		f = f + 1
	}
	//for i := 0; i < 9999; i++{
	//	f = f - 1
	//}
	fmt.Printf("Maxfloat32=%f\n", f)
	var s = []string {"1", "2", "3"}
	for i := 0; i < 100; i++ {
		s = append(s, strconv.Itoa(i))
		//s[i] = strconv.Itoa(i)
	}
	//fmt.Printf("s = %q\n", s)
	//var graph = make(map[string]map[string]bool)
	//fmt.Printf("graph %t \n", graph["a"]["b"])
}
