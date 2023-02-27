package main

import (
	"context"
	"fmt"
	//"math"
	"math/rand"
	"sync"
	"time"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)
var throttle *retryThrottling
func init () {
	throttle = NewRetryThrottling(10, 0.1)
}

type Notify struct {
	sync.Mutex
	cond *sync.Cond  // broadcast to cancel other call
	attempt int      // number of call attempt 
	success int      // the first success call index
	done chan int    // make reserve call cancel 
	failed chan int  // make next call fast
	resp string
	// stop chan bool   // all call failed
}

func NewNotify(attempt int) *Notify{
	notify := &Notify{}
	notify.cond = sync.NewCond(notify)
	notify.success = -1
	notify.done = make(chan int, 1)
	notify.failed = make(chan int, attempt)
	// notify.stop = make(chan bool, 1)
	return notify
}

type retryThrottling struct {
	sync.Mutex
	tokens float32
    ratio float32

	thresh float32  // init to tokens/2
	max float32     // init to tokens
}

func NewRetryThrottling(t float32, r float32) *retryThrottling {
	return &retryThrottling{
		tokens: t,
		ratio: r,
		thresh: t/2,
		max: t,
	}
}

func (t *retryThrottling) success() {
	t.Lock()
	defer t.Unlock()
	t.tokens += t.ratio
	if t.tokens > t.max {
		t.tokens = t.max
	}
}

// return true if need throttling
func (t *retryThrottling) failed() bool {
	t.Lock()
	defer t.Unlock()
	t.tokens--
	if t.tokens < 0 {
		t.tokens = 0
	}
	return t.tokens <= t.thresh
}

// return true if need throttling
func (t *retryThrottling) throttling() bool {
	t.Lock()
	defer t.Unlock()
	return t.tokens <= t.thresh
}

func (t *retryThrottling) Tokens() float32 {
	return t.tokens
}

func clientCall(ctx context.Context, request_id int, idx int) (err error, rsp string) {
	rand.Seed(time.Now().UnixNano())
	diceRoll := rand.Intn(10) + 1 // 1 ~ 10
	switch idx {
	case 0:
		diceRoll = 300
	case 1:
		diceRoll = 300
	default:
		diceRoll = 300
	}
	fmt.Printf("[%v][%08x][invoke] [%d] set cost %dms\n", time.Now().Format("2006-01-02 15:04:05"), request_id, idx, diceRoll)
	var ok bool
	select {
		case <- ctx.Done():
			//fmt.Printf("[%v][%08x][invoke] [%d] done, err = %v \n", time.Now().Format("2006-01-02 15:04:05"), request_id, idx, ctx.Err())
		case <- time.After(time.Duration(diceRoll) * time.Millisecond):
			//fmt.Printf("[%v][%08x][invoke] [%d] done, success \n", time.Now().Format("2006-01-02 15:04:05"), request_id, idx)
			ok = true
	}
	if rand.Intn(100) + 1 <= 11 {
		ok = false
	}
	// if idx == 0 {
	// 	ok = false
	// }
	//ok = false
	if !ok {
		return fmt.Errorf("yy error = %v", ctx.Err()), "i am error response"
	}
	return nil, "i am ok response"
}
func invoke(request_id int, idx int, time_out int, wg *sync.WaitGroup, ntf *Notify) {
	log.Lock()
	log.invoke_total++
	log.Unlock()
	startT := time.Now()
	
	//fmt.Printf("[%v][%08x][invoke] [%d] start\n", time.Now().Format("2006-01-02 15:04:05"), request_id, idx)
	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(time_out) * time.Second)
	//time.Sleep(time.Duration(1) * time.Second)

	defer wg.Done()
	defer cancel()

	waitNotify(request_id, idx, ntf, cancel) // sync go waitNotify
	go waitContext(ctx,request_id, idx)
	ntf.Lock() // sync wait waitNotify goroutine create successful until ntf.cond.Wait()
	if ntf.success != -1 { 
		ntf.cond.Broadcast() // broadcast to all other invoke cancel if alread successful
		tc := time.Since(startT)
		fmt.Printf("[%v][%08x][invoke] [%d] cancel dup call cost = %v\n", time.Now().Format("2006-01-02 15:04:05"), request_id, idx, tc)
		ntf.Unlock()
		return 
	}
	ntf.Unlock()

	err, resp := clientCall(ctx, request_id, idx)

	ntf.Lock()
	defer ntf.Unlock()
	if err == nil {
		log.Lock()
		log.invoke_success++
		log.Unlock()
		if ntf.success == -1 { // first success call
			fmt.Printf("[%v][%08x][invoke] [%d] success\n", time.Now().Format("2006-01-02 15:04:05"), request_id, idx)
			ntf.resp = resp
			ntf.done <- idx
			ntf.success = idx
		} else {
			fmt.Printf("[%v][%08x][invoke] [%d] dup success\n", time.Now().Format("2006-01-02 15:04:05"), request_id, idx)
		}
		throttle.success() // update throttle
		ntf.cond.Broadcast() // broadcast to all other invoke to cancel
	} else {
		log.Lock()
		log.invoke_failed++
		log.Unlock()
		if ntf.success == -1 { // get resp if not yet successful
			ntf.resp = resp
		}
		
		ntf.failed <- idx // fast retry next index
		throttle.failed() // update throttle
		st, ok := status.FromError(err)
		if ok {
			fmt.Printf("[%v][%08x][invoke] [%d] failed, err = %v, code = %d\n", time.Now().Format("2006-01-02 15:04:05"), request_id, idx, err, st.Code())
			if st.Code() == codes.DeadlineExceeded {
				//return sscommon.StatusErrorRpcTimeout
			}
		} else {
			fmt.Printf("[%v][%08x][invoke] [%d] failed, err = %v\n", time.Now().Format("2006-01-02 15:04:05"), request_id, idx, err)
		}
	}
	tc := time.Since(startT)
	log.Lock()
	log.invoke_cost += tc.Milliseconds()
	log.Unlock()
	fmt.Printf("[%v][%08x][invoke] [%d] end, cost = %v\n", time.Now().Format("2006-01-02 15:04:05"), request_id, idx, tc)
}

func waitNotify(request_id int, idx int, ntf *Notify, cancel context.CancelFunc) {
	ntf.Lock() // sync wait goroutine create successful until ntf.cond.Wait()
	go func() {
		//if idx == 1 {time.Sleep(time.Duration(3) * time.Second)}
		fmt.Printf("[%v][%08x][waitNotify] [%d] start...\n", time.Now().Format("2006-01-02 15:04:05"), request_id, idx)

		ntf.cond.Wait()
		ntf.Unlock()
		if idx != ntf.success {
			fmt.Printf("[%v][%08x][waitNotify] call [%d] cancel\n", time.Now().Format("2006-01-02 15:04:05"), request_id, idx)
			cancel()
		}
		fmt.Printf("[%v][%08x][waitNotify] [%d] end...\n", time.Now().Format("2006-01-02 15:04:05"), request_id, idx)
	}()
	
}

func waitContext(ctx context.Context, request_id int, idx int) {
	select {
	case <- ctx.Done():
		fmt.Printf("[%v][%08x][waitContext] [%d] done, err = %v\n", time.Now().Format("2006-01-02 15:04:05"), request_id, idx, ctx.Err())
	}
}

func HedgingCall(request_id int, pwg *sync.WaitGroup,) {
	defer pwg.Done()
	var wg sync.WaitGroup
	var attempt int = 5
	var time_out = 3
	var interval = 500
	ntf := NewNotify(attempt)
	fmt.Printf("[%v][%08x][HedgingCall] display notify = %+v\n", time.Now().Format("2006-01-02 15:04:05"), request_id, ntf)
	fmt.Printf("[%v][%08x][HedgingCall] time_out = %+vs, interval = %dms\n\n", time.Now().Format("2006-01-02 15:04:05"), request_id, time_out, interval)
	startT := time.Now()
FOR:
	for idx := 0; idx < attempt; idx++ {
		wg.Add(1)
		//fmt.Printf("[%v][%08x][HedgingCall] go idx = [%d]\n", time.Now().Format("2006-01-02 15:04:05"), request_id, idx)

		go invoke(request_id, idx, time_out, &wg, ntf) // timeout 4s

		//fmt.Printf("[%v][%08x][HedgingCall] select idx = [%d]\n", time.Now().Format("2006-01-02 15:04:05"), request_id, idx)

		select {
		case <-time.After(time.Duration(interval) * time.Millisecond):
			/*if idx + 1 < attempt*/ {fmt.Printf("[%v][%08x][HedgingCall] select wait idx = [%d]\n", time.Now().Format("2006-01-02 15:04:05"), request_id, idx + 1)}
		case failed_idx := <-ntf.failed:
			fmt.Printf("[%v][%08x][HedgingCall] select failed idx = [%d]\n", time.Now().Format("2006-01-02 15:04:05"), request_id, failed_idx)
			fmt.Printf("[%v][%08x][HedgingCall] throttling = %f idx = [%d]\n", time.Now().Format("2006-01-02 15:04:05"), request_id, throttle.Tokens(), idx)
			if throttle.throttling() {
				fmt.Printf("[%v][%08x][HedgingCall] throttling break idx = [%d]\n", time.Now().Format("2006-01-02 15:04:05"), request_id, idx + 1)
				break FOR
			}
		case success_idx := <-ntf.done:
			fmt.Printf("[%v][%08x][HedgingCall] select success idx = [%d]\n", time.Now().Format("2006-01-02 15:04:05"), request_id, success_idx)
			break FOR
		}
	}
	fmt.Printf("[%v][%08x][HedgingCall] wg.Wait\n", time.Now().Format("2006-01-02 15:04:05"), request_id)
	wg.Wait()
	fmt.Printf("[%v][%08x][HedgingCall] wg.Done\n", time.Now().Format("2006-01-02 15:04:05"), request_id)
	ntf.cond.Broadcast() // broadcast to release failed`s waitNotify

	tc := time.Since(startT)

	fmt.Printf("[%v][%08x][HedgingCall] ********** [%d] done, token = %f, cost %v, resp = %s **********\n", time.Now().Format("2006-01-02 15:04:05"), request_id, ntf.success, throttle.Tokens(),tc, ntf.resp)
	log.Lock()
	if  ntf.success != -1 {
		log.req_success++
	} else {log.req_failed++}
	log.req_cost += tc.Milliseconds()
	log.Unlock()
}

type LOG struct {
	sync.Mutex
	req_total      int
	req_success    int
	req_failed     int
	req_cost       int64

	invoke_total   int
	invoke_success int
	invoke_failed  int
	invoke_cost    int64
}

var log LOG
func(l *LOG) display() {
	fmt.Printf("req total = %d, success = %d, failed = %d, failed_rate = %f%%, avg_cost = %dms\n", l.req_total, l.req_success, l.req_failed, float32(l.req_failed)/float32(l.req_total) * 100, l.req_cost/int64(l.req_total))
	fmt.Printf("invoke total = %d, success = %d, failed = %d, failed_rate = %f%%, avg_cost = %dms\n", l.invoke_total, l.invoke_success, l.invoke_failed, float32(l.invoke_failed)/float32(l.invoke_total) * 100, l.invoke_cost/int64(l.invoke_total))
}
func main () {

	var pwg sync.WaitGroup
	for i := 0; i < 100; i++ {
		log.req_total++
		pwg.Add(1)
		go HedgingCall(i, &pwg)
		//HedgingCall(i, &pwg)
	}
	pwg.Wait()
	
	// time.Sleep(time.Duration(2) * time.Second)
	log.display()
	fmt.Printf("main done\n")
}
