package main

import(
	"fmt"
)
type Itfc0 interface {
	isOK() bool
}

type Itfc interface {
	isOK() bool
	getNumber() int
}

type S0 struct {
	ok bool
}

func (s *S0) isOK() bool {
	return s.ok;
}

type S struct {
	ok bool
	number int
}

func (s *S) isOK() bool {
	return s.ok;
}

func (s *S) getNumber() int {
	return s.number
}

var _ Itfc = (*S)(nil)

func main() {
	s := &S{}
	fmt.Printf("ok = %t\n", s.isOK())

	var x interface{}
	x = "hello"
	val, ok := x.(*S)
	fmt.Printf("ok = %t, val = %v\n", ok, val)
}
