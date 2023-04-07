package main
import(
	"fmt"
	"time"
)

func sendCh(ch chan int) {
	for i := 0; i < 9999; i++ {
		fmt.Printf("send e %d\n", i)
		ch <- i
		time.Sleep(time.Second * 1)
		if i == 5 {
			close(ch)
			break
		}
	}
}
func main() {
	ch := make(chan int, 0)
	go sendCh(ch)
	for e := range ch {
		fmt.Printf("get e = %v\n", e)
	}
	fmt.Printf("end\n")
}
