package main

import(
	"fmt"
	"os"
	"reflect"
	"io"
)


func main() {
	var w io.Writer = os.Stdout
	fmt.Println(reflect.TypeOf(w)) 
}
