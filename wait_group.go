package main

import (
	//"sync"
	"fmt"
	list2 "container/list"
)

func main() {
	//var wg sync.WaitGroup
	//
	//for i := 0; i < 5; i++ {
	//	wg.Add(1)
	//	go func(i int) {
	//		fmt.Println("num: %d", i)
	//		wg.Done()
	//	}(i)
	//}
	//
	//wg.Wait()
	//fmt.Println("main quit")

	list := list2.New()
	list.PushBack(3)
	list.PushBack(4)
	list.PushBack(5)
	list.PushBack(6)

	 n := &list2.Element{}
	for v := list.Front(); v != nil; v = n {
		fmt.Println(v.Value.(int))
		n = v.Next()
		if v.Value.(int) >= 4 {
			list.Remove(v)
		}
	}

	for v := list.Front(); v != nil; v = v.Next() {
		fmt.Println(v.Value.(int))
	}


}
