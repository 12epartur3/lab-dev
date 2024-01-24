package main

import(
	"fmt"
)

type AnyByteSlice = interface {
	int
}

func mString(s string) {
	s = "789"
}

func mMap(m map[string]string) {
	m["123"] = "456"
}

func mSlice(s []string) {
	fmt.Printf("cap = %v, len = %v\n", cap(s), len(s))
	s = append(s, "789")
	fmt.Printf("cap = %v, len = %v\n", cap(s), len(s))
}
func main() {
	var s string
	var m map[string]string = map[string]string{}
	var sl []string = make([]string,0 ,10)
	mString(s)
	mMap(m)
	mSlice(sl)
	fmt.Printf("s = %v\n", s)
	fmt.Printf("m = %v\n", m)
	fmt.Printf("sl = %v\n", sl)
}
