package main

import (
	"sync"
	"fmt"
	"container/list"
	"time"
)

func printList(mylist *list.List) {
	for e := mylist.Front(); e != nil; e = e.Next() {
		fmt.Printf("%s ",e.Value.(string))
	}
	fmt.Printf("\n");
}

func main() {
	var nameListMap map[string] *list.List
	var lockMap map[string] *sync.Mutex

	nameListMap = make(map[string] *list.List)
	lockMap = make(map[string] *sync.Mutex)

	if _, found := nameListMap["name"]; !found {
		nameListMap["name"] = list.New()
	}
	if _, found := lockMap["name"]; !found {
		lockMap["name"] = new(sync.Mutex)
	}
	go func(lockMap map[string] *sync.Mutex) {
		lockMap["name"].Lock()
		defer lockMap["name"].Unlock()
		fmt.Printf("start lock 5 s\n")
		time.Sleep(time.Duration(5)*time.Second)
		fmt.Printf("end lock\n")
	}(lockMap)
	time.Sleep(time.Duration(1)*time.Second)
	fmt.Printf("start list\n");
	lockMap["name"].Lock()
	fmt.Printf("real start list\n");
	nameListMap["name"].PushBack("yuanye")
	nameListMap["name"].PushBack("naruto")
	for k, v := range nameListMap {
		fmt.Printf("k = %s\n",k);
		printList(v);
	}
	var nameList *list.List = nameListMap["name"]
	printList(nameList)
	var name string = nameList.Front().Value.(string)
	fmt.Printf("name = %v\n", name)
	nameList.MoveToBack(nameList.Front())
	lockMap["name"].Unlock()
	for k, v := range nameListMap {
		fmt.Printf("k = %s\n",k);
		printList(v);
	}
	go func(nameListMap map[string] *list.List) {
		fmt.Printf("go routine start\n");
		var nameList *list.List = nameListMap["name"]
			printList(nameList)
		fmt.Printf("go routine end\n");
	}(nameListMap)
	time.Sleep(time.Duration(1)*time.Second)
}
