package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"

	"go.uber.org/atomic"
)

const (
	kFirst = "_first"
	kRetry = "_retry"
)
type Bucket interface {
	Add(label string, n float64)
	//Sub(label string, n float64)

	SubBucket(Bucket)
	AddBucket(Bucket)

	Get(label string) float64
	//Set(label string, n float64)

	Reset()
}
type SlidingWindow struct {
	sync.RWMutex
	windowSize int32
	windowTime int32
	windowIdx  int32
	windowBuckets []Bucket
	sumBucket  Bucket
	done       chan(int)
}

func NewSlidingWindow(windowSize int32, windowTime int32, f func() Bucket) *SlidingWindow{
	var s = &SlidingWindow{
		windowSize: windowSize,
		windowTime: windowTime,
		windowIdx: 0,
		windowBuckets: make([]Bucket, windowSize),
		sumBucket: f(),
		done: make(chan int, 1),
	}
	for i := 0; i < int(s.windowSize); i++ {
		s.windowBuckets[i] = f()
	}
	go func() {
		for{
			//fmt.Printf("in for")
			select {
				case <-time.After(time.Millisecond * time.Duration(s.windowTime)):
				{
					s.Update()
				}
				case <-s.done:
				{
					fmt.Printf("SlidingWindow stopping")
					return
				}
			}
        }
	}()
	return s
}

func (s *SlidingWindow) Stop() {
	fmt.Printf("SlidingWindow start stop")
	s.done <-1
	fmt.Printf("SlidingWindow stop ok")
}

func (s *SlidingWindow) Add(label string, n float64) {
	s.Lock()
	defer s.Unlock()
	s.windowBuckets[s.windowIdx].Add(label, n)
}

func (s *SlidingWindow) Get(label string) float64 {
	s.RLock()
	defer s.RUnlock()
	return s.sumBucket.Get(label) + s.windowBuckets[s.windowIdx].Get(label)
}

func (s *SlidingWindow) GetSumBucket(label string) float64 {
	s.RLock()
	defer s.RUnlock()
	return s.sumBucket.Get(label)
}

func (s *SlidingWindow) Reset() {
	s.Lock()
	defer s.Unlock()
	s.windowBuckets[s.windowIdx].Reset()
}

func (s *SlidingWindow) Update() {
	s.Lock()
	defer s.Unlock()
	// fmt.Printf("windowIdx = %d, l1_first, sumBucket = %f, bucket = %f\n", s.windowIdx, s.sumBucket.Get("l1_first"), s.windowBuckets[s.windowIdx].Get("l1_first"))
	// fmt.Printf("windowIdx = %d, l1_retry, sumBucket = %f, bucket = %f\n", s.windowIdx, s.sumBucket.Get("l1_retry"), s.windowBuckets[s.windowIdx].Get("l1_retry"))
	s.sumBucket.AddBucket(s.windowBuckets[s.windowIdx])
	s.windowIdx = (s.windowIdx + 1) % s.windowSize
	s.sumBucket.SubBucket(s.windowBuckets[s.windowIdx])
	// fmt.Printf("windowIdx = %d, l1_first, sumBucket = %f, bucket = %f\n", s.windowIdx, s.sumBucket.Get("l1_first"), s.windowBuckets[s.windowIdx].Get("l1_first"))
	// fmt.Printf("windowIdx = %d, l1_retry, sumBucket = %f, bucket = %f\n", s.windowIdx, s.sumBucket.Get("l1_retry"), s.windowBuckets[s.windowIdx].Get("l1_retry"))
	s.windowBuckets[s.windowIdx].Reset()
}

type retryThrottling struct {
	tokens atomic.Float64
	failureRatio float64 // max rpc failed percentage of stopping retry, 0~1
	retryRatio   float64 // max label failed percentage of stopping retry, 0~1
	sw     *SlidingWindow
	thresh float64 // init to tokens/2
	max    float64 // init to tokens
}

func NewRetryThrottling(t float64, fr float64, rr float64) *retryThrottling {
	var windowSize int32 = 5 
	var windowTime int32 = 1000
	var throttle =  &retryThrottling{
		failureRatio:   fr,
		retryRatio: rr,
		sw: NewSlidingWindow(windowSize, windowTime, NewRpcStatusBucket),
		thresh:  t / 2,
		max:     t,
	}
	throttle.tokens.Store(t)
	return throttle
}

func (t *retryThrottling) success() {
	t.tokens.Add(t.failureRatio)
	if t.tokens.Load() > t.max {
		t.tokens.Store(t.max)
	}
}

func (t *retryThrottling) failed() {
	t.tokens.Sub(1)
	if t.tokens.Load() < 0 {
		t.tokens.Store(0)
	}
}

// return true if need throttling
func (t *retryThrottling) throttling(label string) bool {
	if t.tokens.Load() <= t.thresh {
		return true
	}
	return t.labelRetryRatio(label) > t.retryRatio
}

func (t *retryThrottling) attemptIdx(idx int, label string) {
	if idx == 0 {
		t.sw.Add(label + kFirst, 1)
	} else {
		t.sw.Add(label + kRetry, 1)
	}
}

func (t *retryThrottling) Tokens() float64 {
	return t.tokens.Load()
}

func (t *retryThrottling) labelRetryRatio(label string) float64{
	first := t.sw.Get(label + kFirst)
	retry := t.sw.Get(label + kRetry)
	if first == 0 && retry != 0{
		return 1
	}
	if first == 0 && retry == 0{
		return 0
	}
	return retry / first
}

func (t *retryThrottling) labelFirst(label string) float64{
	return t.sw.Get(label + kFirst)
}

func (t *retryThrottling) labelRetry(label string) float64{
	return t.sw.Get(label + kRetry)
}



type rpcStatus struct {
	labelMap map[string]float64
}

func NewRpcStatusBucket () Bucket {
	r := &rpcStatus{
		labelMap: make(map[string]float64),
	}
	return r
}

func (s *rpcStatus) Add(label string, n float64) {
	s.Set(label, s.Get(label) + n)
}

func (s *rpcStatus) AddBucket(b Bucket) {
	for k := range (b.(*rpcStatus)).labelMap {
		fmt.Printf("AddBucket1 %s sumBucket = %f, bucket = %f\n", k, s.Get(k),  b.Get(k))
		s.Set(k, s.Get(k) + b.Get(k))
		fmt.Printf("AddBucket2 %s sumBucket = %f, bucket = %f\n", k, s.Get(k),  b.Get(k))
	}
}

func (s *rpcStatus) SubBucket(b Bucket) {
	for k := range s.labelMap {
		fmt.Printf("SubBucket1 %s sumBucket = %f, bucket = %f\n", k,  s.Get(k),  b.Get(k))
		s.Set(k, s.Get(k) - b.Get(k))
		fmt.Printf("SubBucket2 %s sumBucket = %f, bucket = %f\n", k,  s.Get(k),  b.Get(k))
	}
}

func (s *rpcStatus) Get(label string) float64 {
	return s.labelMap[label]
}

func (s *rpcStatus) Set(label string, n float64) {
	s.labelMap[label] = n
}

func (s *rpcStatus)Reset() {
	s.labelMap = make(map[string]float64)
}

func Tick(ms int32) (func() (bool ,int32), func()) {
    var timeIsUp atomic.Bool
	var tick atomic.Int32
	fmt.Printf("init time is up=%t, tick=%d\n", timeIsUp.Load(), tick.Load())
	done := make(chan int, 1)
	go func(){
                for{
                        fmt.Printf("in for\n")
                        select {
                                case <-time.After(time.Millisecond * time.Duration(ms)):
                                {
										timeIsUp.Store(true)
										t := tick.Load()
										tick.Store((t + 1) % 2)
                                        fmt.Printf("time is up, time = %dms, tick=%d\n", ms, tick.Load())
                                }
                                case <-done:
                                {
                                        fmt.Printf("tick stop\n")
                                        return
                                }
                        }
                }
        }()

	Trigger := func() (bool, int32) {
                if !timeIsUp.Load() {
                        //fmt.Printf("timeIsUp = false, tick=%d\n", tick.Load())
                        return false, tick.Load()
                }

		preTimeIsUp := timeIsUp.Swap(false)
                fmt.Printf("timeIsUp=%t, tick=%d\n", preTimeIsUp, tick.Load())
                return preTimeIsUp, tick.Load()
        }
	Stop := func() {
		fmt.Printf("stop\n")
		done <-1
		fmt.Printf("stop ok\n")
	}
	return Trigger, Stop
}

var N int = 1000

func mockReq(wg *sync.WaitGroup, t *retryThrottling, label string, id string) {
	percentage := 30
	var retryCount int
	var failureCount int
	for i := 0; i < N; i++ {
		t.attemptIdx(0, label)
		cost := rand.Intn(100) % 20 + 10
		time.Sleep(time.Millisecond * time.Duration(cost))
		rand.Seed(time.Now().UnixNano())
		diceRoll := rand.Intn(100) + 1 // 1 ~ 100
		fmt.Printf("id = %s, lable=%s, first, tokens=%f, first=%f, retry=%f, retryratio=%f, throttling=%t\n", id, label, t.Tokens(), t.labelFirst(label), t.labelRetry(label), t.labelRetryRatio(label), t.throttling(label))
		if diceRoll <= percentage {
			t.failed()
			failureCount ++
		} else {
			t.success()
		}
		if diceRoll <= percentage  && !t.throttling(label){
		//if diceRoll <= percentage {
			t.attemptIdx(1, label)
			fmt.Printf("id = %s, label=%s, retry, tokens=%f, first=%f, retry=%f, retryratio=%f, throttling=%t\n", id, label, t.Tokens(),t.labelFirst(label), t.labelRetry(label), t.labelRetryRatio(label), t.throttling(label))
			retryCount ++
		}
	}
	fmt.Printf("id=%s, label=%s, total=%d, failureCount=%d, retryCount=%d, sum faliureRatio=%f, retryRatio=%f, \n", id, label, N, failureCount, retryCount, float32(failureCount) / float32(N), float32(retryCount) / float32(N))
	wg.Done()
}

func mockSlide(sw *SlidingWindow, label string) {
	percentage := 30
	retry_label := label+"_retry"
	for i := 0; i < N; i++ {
		cost := rand.Intn(100) % 10 + 10
		time.Sleep(time.Millisecond * time.Duration(cost))
		rand.Seed(time.Now().UnixNano())
		diceRoll := rand.Intn(100) + 1 // 1 ~ 100
		sw.Add(label, 1)
		retryRatio := sw.Get(retry_label) / sw.Get(label)
		// if i >= 1000 {
		// 	percentage = 10
		// }
		fmt.Printf("label=%s, Get=%f, retry=%f, GetSumBucket=%f, retry=%f, ratio=%f\n", label, sw.Get(label), sw.Get(retry_label), sw.GetSumBucket(label), sw.GetSumBucket(retry_label), retryRatio)
		if diceRoll <= percentage && retryRatio < 0.05{
			sw.Add(retry_label, 1)
		}

	}

}
func main() {
	var failureRatio float64
	var retryRatio float64
	failureRatio = 0.8
	retryRatio = 0.1
	t := NewRetryThrottling(100, failureRatio, retryRatio)
	var wg sync.WaitGroup
	wg.Add(1)
	go mockReq(&wg, t, "l1", "1")
	// wg.Add(1)
	// go mockReq(&wg, t, "l1", "2")
	// wg.Add(1)
	// go mockReq(&wg, t, "l2", "1")
	// wg.Add(1)
	// go mockReq(&wg, t, "l2", "2")
	// wg.Add(1)
	// go mockReq(&wg, t, "l3", "2")
	wg.Wait()
	//t.success()
	//var count int
	//for{
	//	fmt.Printf("trig=%t\n", t.Trigger())
	//	time.Sleep(time.Second)
	//	count ++
	//	if count == 8 {
	//		t.Stop()
	//	}
	//}
	// sw := NewSlidingWindow(5, 1000, NewRpcStatusBucket)
	// mockSlide(sw, "l1")
}
