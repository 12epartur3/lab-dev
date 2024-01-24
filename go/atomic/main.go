package main

import(
	"fmt"
	"sync/atomic"
	"unsafe"
)

type V struct {
	a int
	b string
	c float64
}

func main() {
	v1 := &V{
		a: 10,
		b: "yy",
		c: 100,
	}

	//var usfP unsafe.Pointer
	//atomic.StorePointer(&usfP, unsafe.Pointer(v1))
	//v := (*V)(atomic.LoadPointer(&usfP))

	var P *V
	fmt.Printf("P = %v\n", P)
	atomic.StorePointer((*unsafe.Pointer)(unsafe.Pointer(&P)), unsafe.Pointer(v1))
	v := (*V)(atomic.LoadPointer((*unsafe.Pointer)(unsafe.Pointer(&P))))
	fmt.Printf("v = %v\n", *v)
	fmt.Printf("v.a = %d, v.b = %s, v.c = %f\n", v.a, v.b, v.c)

	var P32 *uint32 = new(uint32)
	atomic.StoreUint32(P32, uint32(32))
	v32 := atomic.LoadUint32(P32)
	fmt.Printf("v32 = %v\n", v32)
}
