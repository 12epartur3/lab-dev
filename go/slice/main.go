package main

import(
	"fmt"
)
type addr struct {
	Addr string
}
type zkS struct {
	s []string
}


func main() {
	var array []*int
	fmt.Printf("len(nil) = %d\n", len(array))
	return
	zk := &zkS{}
	//s = s[:0]
	zk.s = append(zk.s, "123")
	zk.s = append(zk.s, "456")
	zk.s = zk.s[:0]
	fmt.Printf("s = %v\n", zk.s)
	s :=[...] int {1, 2, 3, 4}
        s1 := s[1:3]
        fmt.Printf("s = %v\n", s)
        fmt.Printf("s1= %v\n", s1)
	s1[0] = 999
        fmt.Printf("s = %v\n", s)
        fmt.Printf("s1= %v\n", s1)
}
