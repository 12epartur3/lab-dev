package main

import(
	"math"
	"fmt"
)

func Round(f float64, n int) float64 {
    pow10_n := math.Pow10(n)
    return math.Trunc((f+0.5/pow10_n)*pow10_n) / pow10_n
}

func main() {
	var f float64
	f = 1.23455678
	fmt.Printf("%.8f\n", f)
	fmt.Printf("%.8f\n", Round(f, 4))
}
