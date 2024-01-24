package main

import(
	"time"
	"fmt"
)

func deferfunc(beginTime time.Time, e error) {
	cost := time.Since(beginTime).Microseconds()
	fmt.Printf("cost = %d us\n", cost)
	if e != nil {
		fmt.Printf("deferfunc e = %v\n", e)
	}
}

func test() (err error){
	defer func() {
		if err != nil {
			fmt.Printf("test error = %s\n", err.Error())
		}
	}()
	//err = fmt.Errorf("dod")
	return fmt.Errorf("dod123")
}
func call() {
	//x := 10
	var err error
	defer func() {
		testdefer(err)
	}()
	//err = test()
	//fmt.Printf("test error = %s\n", err.Error())
	//defer func() {
	//	if err != nil {
	//		fmt.Printf("defer e = %v\n", err)
	//	}
	//	deferfunc(time.Now(), err)
	//}()
	if err == nil {
		fmt.Printf("err = nil\n")
	}
	err = fmt.Errorf("test error")
	//if err != nil {
	//	fmt.Printf("err = %v\n", err)
	//}
	//time.Sleep(time.Second * 2)
	//fmt.Println("call end")
}

func testdefer(err error) {
	if err != nil {
		fmt.Printf("defer e = %v\n", err)
	} else {
		fmt.Printf("defer e = nil\n")
	}
}
func main() {
	call()
}
