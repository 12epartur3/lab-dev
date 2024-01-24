package main
import(
        "fmt"
	"reflect"
)

type get interface {
	getBase() base
}

type base struct {
        b1 int
        b2 float64
}

type derive1 struct {
        b base
        d1 int
        d2 float64
}
func (d1 *derive1) getBase() base {
	return d1.b
}

type derive2 struct {
        b base
        d3 int
        d4 float64
}
func (d2 *derive2) getBase() base {
	return d2.b
}

func testGetBase(d any) {
	value, ok := d.(get)
	fmt.Printf("value = %v, ok = %t\n", value, ok)
	fmt.Printf("reflect.TypeOf(d) = %v, reflect.TypeOf(value) = %v\n", reflect.TypeOf(d), reflect.TypeOf(value))
	fmt.Printf("base = %+v\n", value.getBase())
}
func switchI(d any) {
	//switch reflect.TypeOf(d) {
	switch d.(type) {
		case *derive1:
			fmt.Printf("type d1\n")
		case *derive2:
			fmt.Printf("type d2\n")
		//default:
		//	fmt.Printf("default\n")
	}
}
func main() {
        var x int = 10;
        var i any = &x;
        var i2 any = i
        fmt.Printf("x = %d, i = %v, i2 = %v\n", x, *i.(*int), *i2.(*int))
        x++
        //val := i.(*int)
        //*val = 0
        fmt.Printf("x = %d, i = %v, i2 = %v\n", x, *i.(*int), *i2.(*int))
        //fmt.Printf("x = %d, i = %v, i2 = %v, val = %d\n", x, i, i2, *val)
        d1 := &derive1{
		b:  base{
			b1: 12,
			b2: 12.34,
		},
                d1: 123,
                d2: 123.456,
        }
        d2 := &derive2{
		b:  base{
			b1: 1112,
			b2: 1112.34,
		},
                d3: 11123,
                d4: 11123.456,
        }
	//var s any = d1
        fmt.Printf("d1 = %+v\n", d1)
        fmt.Printf("d2 = %+v\n", d2)
        //fmt.Printf("s = %+v, reflect.TypeOf(s) = %v\n", s, reflect.TypeOf(s))
	//value, ok := s.(derive1)
	//fmt.Printf("value = %v, ok = %t\n", value, ok)
	//testGetBase(d1)
	//testGetBase(d2)
	fmt.Printf("reflect.TypeOf(d1) = %v, reflect.ValueOf(d1) = %v\n", reflect.TypeOf(d1), reflect.ValueOf(d1))
	switchI(d1)
	switchI(d2)
	switchI(10)
}
