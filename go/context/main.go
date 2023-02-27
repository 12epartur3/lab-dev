package main

import(
	"fmt"
	"context"
	"time"
)


func main()  {
    ctx,cancel := context.WithCancel(context.Background())
    //ctx,cancel := context.WithTimeout(context.Background(), 5 * time.Second)
    //defer cancel()
    //ctx_,cancel_ := context.WithTimeout(context.Background(), 3 * time.Second)
    //defer cancel_()
    //select {
    //    case <- ctx_.Done():
    //    	fmt.Printf("ctx_ done, err = %v\n", ctx_.Err())
    //}
    //select {
    //    case <- ctx.Done():
    //    	fmt.Printf("ctx done, err = %v\n", ctx.Err())
    //}
    //go Speak(ctx)
    //time.Sleep(5*time.Second)
    go TestCancel(ctx)
    time.Sleep(2*time.Second)
    fmt.Printf("call cancel\n")
    cancel()
    time.Sleep(5*time.Second)
}

func TestCancel(ctx context.Context) {
	fmt.Printf("i am sleep\n")
	time.Sleep(4*time.Second)
	fmt.Printf("i am weak up\n")
}

func Speak(ctx context.Context)  {
    for {
	time.Sleep(1*time.Second)
        select {
        case <- ctx.Done():
            fmt.Printf("i am stop, err = %v\n", ctx.Err())
            return
        default:
            fmt.Println("balabalabalabala")
        }
    }
}

