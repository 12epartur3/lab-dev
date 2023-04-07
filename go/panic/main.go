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
func testPanic() {
	fmt.Printf("step 0\n")
	panic("i am panic")
	fmt.Printf("step 1\n")
}
func main() {
	defer func() {
		if p := recover(); p != nil {
			fmt.Printf("recover panic: %v\n", p)
		}
	}()
	testPanic()
	fmt.Printf("step 2\n")
	fmt.Printf("step 3\n")
	fmt.Printf("step 4\n")
	time.Sleep(time.Duration(1) * time.Second)
}
