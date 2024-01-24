package main

import(
	"fmt"
	"yytest/dir1"
	"yytest/dir1/dir2"
)

func main() {
	algo := dir1pkg.EvaluationAlog{
		Control: "123",
		Treatment: "456",
	}
	fmt.Printf("algo = %v\n", algo)
	s := dir2pkg.NewServiceImpl("test server")
	fmt.Printf("server = %v\n", s)
}


