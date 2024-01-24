// Golang Program to illustrate the usage of
// CompareAndSwapInt64 function

// Including main package
package main

// importing fmt and sync/atomic
import (
	"fmt"
	"sync/atomic"
)

// Main function
func main() {

	// Assigning variable values to the int64
	var (
		i int64 = 686788787
	)

	// Swapping
	var old_value = atomic.SwapInt64(&i, 56677)

	// Printing old value and swapped value
	fmt.Println("Swapped:", i, ", old value:", old_value)

	// Calling CompareAndSwapInt64
	// method with its parameters
	Swap := atomic.CompareAndSwapInt64(&i, 156677, 908998)

	// Displays true if swapped else false
	fmt.Println(Swap)
	fmt.Println("The Value of i is: ",i)
}

