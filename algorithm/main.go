package main

import (
	"fmt"
)

func main()  {
	testBFS()
}

func testHeapSort()  {
	arr := make([]uint32, 0, 0)
	arr = append(arr, 0, 7, 5, 19, 8, 4, 1, 20, 13, 16)
	ptrHeapSort := NewHeapSort(arr, 5)
	ptrHeapSort.sort()

	fmt.Println(arr)
}

func testBFS()  {
	ptrGraph := NewGraph(8)
	ptrGraph.AddEdge(0, 1)
	ptrGraph.AddEdge(0, 3)
	ptrGraph.AddEdge(1, 2)
	ptrGraph.AddEdge(1, 4)
	ptrGraph.AddEdge(2, 5)
	ptrGraph.AddEdge(3, 4)
	ptrGraph.AddEdge(4, 5)
	ptrGraph.AddEdge(4, 6)
	ptrGraph.AddEdge(5, 7)
	ptrGraph.AddEdge(6, 7)

	for w, l := range ptrGraph.adjacency {
		fmt.Printf("第 %d 顶点, 相邻顶点 ： ", w)
		for v := l.Front(); v != nil; v = v.Next() {
			fmt.Print(v.Value.(int), " ")
		}
		fmt.Print("\n")

	}
	ptrGraph.BFS(0, 6)
}


