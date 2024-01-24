package main

import (
	"fmt"
	"reflect"
)

type Writer interface {
	Write(buf []byte) (int, error)
}

type DummyWriter struct{}
func (DummyWriter) Write(buf []byte) (int, error) {
	return len(buf), nil
}

func main() {
	var data [10]*int
	fmt.Printf("reflect.TypeOf(data) = %v\n", reflect.TypeOf(data))
	//var t int = 20
	//var t2 string = string("123") 
	t2 := string("123")
	fmt.Printf("t2 = %s\n", t2)
	//t3 := 20
	var x interface{} = DummyWriter{}
	var y interface{} = "abc"
	fmt.Printf("reflect.TypeOf(x) = %v\n", reflect.TypeOf(x))
	fmt.Printf("reflect.TypeOf(y) = %v\n", reflect.TypeOf(y))
	// Now the dynamic type of y is "string".
	var w Writer
	var ok bool

	// Type DummyWriter implements both
	// Writer and interface{}.
	w, ok = x.(Writer)
	fmt.Println(w, ok) // {} true
	x, ok = w.(interface{})
	fmt.Println(x, ok) // {} true

	// The dynamic type of y is "string",
	// which doesn't implement Writer.
	w, ok = y.(Writer)
	fmt.Println(w, ok) // <nil> false
	w = y.(Writer)     // will panic
}
