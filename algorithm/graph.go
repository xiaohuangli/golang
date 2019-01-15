package main

import (
	"container/list"
	"fmt"
)

type graph struct {
	vertexNum int	// 顶点数量
	adjacency []*list.List // 邻接表表示，顶点-> 可达的顶点list
}

// 初始化图大小，vertexNum：顶点数量
func NewGraph(vertexNum int) *graph {
	return &graph{
		vertexNum : vertexNum,
		adjacency : make([]*list.List, vertexNum,vertexNum),
	}
}

func (this *graph) AddEdge(vertex1, vertex2 int) {
	l1 := this.adjacency[vertex1]
	if l1 == nil {
		l1 = list.New()
		this.adjacency[vertex1] = l1
	}
	l1.PushBack(vertex2)

	l2 := this.adjacency[vertex2]
	if l2 == nil {
		l2 = list.New()
		this.adjacency[vertex2] = l2
	}
	l2.PushBack(vertex1)
}
//广度优先遍历
func (this *graph) BFS(s, t int) {
	if s == t {
		return
	}
	visit := make([]bool, this.vertexNum, this.vertexNum)
	visit[s] = true
	queue := list.New()
	prev := make([]int, this.vertexNum, this.vertexNum)
	for i := 0; i < this.vertexNum; i++ {
		prev[i] = -1
	}
	queue.PushBack(s)

	for {
		if queue.Len() == 0 {
			break
		}
		node := queue.Front()
		w := node.Value.(int)

		queue.Remove(node)
		adj := this.adjacency[w]
		for v := adj.Front(); v != nil; v = v.Next() {
			q := v.Value.(int)
			if visit[q] == false {
				prev[q] = int(w)
				if q == t {
					print(prev, s, t)
					return
				}
				visit[q] = true
				queue.PushBack(q)
			}
		}
	}
}




func print(prev []int, s, t int) { // 递归打印 s->t 的路径
	if prev[t] != -1 && t != s {
		print(prev, s, prev[t])
	}
	fmt.Print(t,  " ")
}

