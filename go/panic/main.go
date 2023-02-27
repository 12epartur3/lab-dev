package main
import(
	"fmt"
	"time"
)

func waitCi(ci <-chan int) {
	fmt.Printf("wait ci\n")
	<-ci
	fmt.Printf("end ci\n")
}
func main() {
	defer func() {
		if p := recover(); p != nil {
			fmt.Printf("recover panic: %v\n", p)
		}
	}()
	var ci chan int
	ci = make(chan int, 1)
	go waitCi(ci)
	fmt.Printf("step 1\n")
	go panic("i am panic")
	select {}
	fmt.Printf("step 2\n")
	fmt.Printf("step 3\n")
	fmt.Printf("step 4\n")
	ci <- 1
	time.Sleep(time.Duration(1) * time.Second)
}
