package main

import(
	"fmt"
)
type benchKey struct {
	K1 string
	K2 int
}
type benchmark struct {
	K1 string
	K2 string
}

func main() {
	bm := make(map[string]benchmark)
	benchmarkMap := make(map[benchKey]*benchmark)
	k1 := benchKey{
		K1: "key1",
		K2: 1,
	}
	k2 := benchKey{
		K1: "key2",
		K2: 2,
	}
	benchmarkMap[k1] = &benchmark{}
	v := benchmarkMap[k1]
	v.K1 = "1"
	v.K2 = "2"
	fmt.Printf("benchmarkMap = %+v\n", benchmarkMap)
	fmt.Printf("benchmarkMap[k1] = %+v\n", benchmarkMap[k1])
	fmt.Printf("benchmarkMap[k2] = %+v\n", benchmarkMap[k2])
	bm["1"] = benchmark{
		K1: "1",
		K2: "1",
	}
	v2 := bm["1"]
	v2.K1 = "2"
	v2.K2 = "2"
	fmt.Printf("bm = %+v\n", bm)
	fmt.Printf("v2 = %+v\n", v2)
}
