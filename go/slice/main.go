package main

import(
	"fmt"
	"strings"
)
type addr struct {
	Addr string
}
type zkS struct {
	s []string
}

func dedupFunc(dup_key []string) ([]string) {
	appearanceMap := map[string]int{}
	deDupBlacklistKws := []string{}

	for _, v := range dup_key {
		appearKey := strings.ToUpper(v)
		appearanceMap[appearKey] += 1
		if appearanceMap[appearKey] == 1 {
			deDupBlacklistKws = append(deDupBlacklistKws, v)
		}
	}

	return deDupBlacklistKws
}


func main() {
	var sls []int64
	sls = append(sls, 1)
	sls = append(sls, 2)
	sls = append(sls, 3)
	fmt.Printf("sls = %v\n", sls)
	sls = nil
	fmt.Printf("sls = %v\n", sls)
	sls = append(sls, 1)
	sls = append(sls, 2)
	sls = append(sls, 3)
	fmt.Printf("sls = %v\n", sls)
}
