package main

import(
	"time"
	"fmt"
)

func deferfunc(beginTime time.Time, i int) {
	cost := time.Since(beginTime).Microseconds()
	fmt.Printf("cost = %d us, i = %d\n", cost, i)
}

func call() {
	x := 10
	defer func() {deferfunc(time.Now(), x)
	}()
	x = x*x
	time.Sleep(time.Second * 2)
	fmt.Println("call end")
}

func main() {
	call()
}
