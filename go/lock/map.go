package main

import(
	"fmt"
	"time"
	"sync"
)
func main() {
	var conc_map map[int] *sync.Mutex 
	conc_map = make(map[int] *sync.Mutex)
	for i := 0; i < 10; i++ {
		conc_map[i] =  new(sync.Mutex)
	}
	for i := 0; i < 10; i++ {
		go func(conc_map map[int] *sync.Mutex) {
			for k := 0; k < 10; k++ {
				fmt.Printf("lock k = %d\n", k)
				conc_map[k].Lock()
				fmt.Printf("unlock k = %d\n", k)
				conc_map[k].Unlock()
				if k == 9 {
					k = 0
				}
			}
		}(conc_map)
	}
	fmt.Printf("sleep \n")
	time.Sleep(time.Duration(20)*time.Second)
	fmt.Printf("main end \n")


}

