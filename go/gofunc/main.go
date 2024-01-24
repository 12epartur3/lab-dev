package main

import(
	"fmt"
	"time"
)


func test () {
	var sls = []int {1, 2, 3, 4, 5}
	fmt.Printf("sls0 = %v\n", sls)
	go func(){
		fmt.Printf("sls1 = %v\n", sls)
			time.Sleep(2 * time.Second)
		fmt.Printf("sls2 = %v\n", sls)
	}()
	return
}
func main() {

	test()
	fmt.Printf("out\n")
	time.Sleep(10 * time.Second)
}
