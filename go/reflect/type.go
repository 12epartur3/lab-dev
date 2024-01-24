package main

import (
	"fmt"
	"reflect"
	"io"
	"os"
)

type myErr struct{
	i int
	s string
}
func (m *myErr) Error() string {
	return ""
}
func main() {
	var err error = &myErr{
		i: 99,
		s: "9",
	}
	var r io.Reader = os.Stdin // os.Stdin is of type *os.File which implements io.Reader
	//var e error
	var a any
	fmt.Printf("Elem = %v\n", reflect.TypeOf((*error)(nil)).Elem())
	fmt.Printf("Type = %v\n", reflect.TypeOf(a))
	v := reflect.ValueOf(r) // r is interface wrapping *os.File value
	fmt.Printf("v.Elem() = %v\n", v.Elem())

	//var val int = 999;
	var i any = err
	fmt.Printf("type = %v\n", reflect.TypeOf(i))
	fmt.Printf("value = %v\n", reflect.ValueOf(i).String())

	errType := reflect.TypeOf(i)
	errValue := reflect.ValueOf(i)
	fmt.Printf("errType.String() = %v\n", errType.String())
	fmt.Printf("errType.Kind() = %v\n", errType.Kind())
	fmt.Printf("errType.Elem() = %v\n", errType.Elem())
	fmt.Println()
	fmt.Printf("errValue.Type() = %v\n", errValue.Type())
	fmt.Printf("errValue.IsNil() = %v\n", errValue.IsNil())
	fmt.Printf("errValue.String() = %v\n", errValue.String())
	fmt.Printf("errValue.Elem() = %v\n", errValue.Elem())
}
