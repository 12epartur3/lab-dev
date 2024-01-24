package main
import (
	"unsafe"
	"fmt"
	"reflect"
)

type I32Type int32

func main() {
	var a int = 1
	//var b I32Type = I32Type(a)
	fmt.Printf("a = %d, &a = %p, unsafe.Pointer(&a) = %p\n", a, &a, unsafe.Pointer(&a))
	fmt.Printf("reflect.TypeOf(a) = %v, reflect.TypeOf(&a) = %v, reflect.TypeOf(unsafe.Pointer(&a)) = %v\n", reflect.TypeOf(a), reflect.TypeOf(&a), reflect.TypeOf(unsafe.Pointer(&a)))
}
