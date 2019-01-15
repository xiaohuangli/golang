package main

import (
	//"sync"
	"fmt"
)

type HeapSort struct {
	sli 		[]uint32
	sortNum 	uint32
}

func NewHeapSort(sli []uint32, sortN uint32) *HeapSort {
	return &HeapSort{
		sli:sli,
		sortNum:sortN,
	}
}

// 建立堆。从第一个非叶子节点开始
func (this *HeapSort) buildHeap() {
	for i := this.sortNum/2; i > 0;i-- {
		this.heapify(this.sortNum, i)
	}
}

// 向下构建堆
func (this *HeapSort) heapify(n, i uint32)  {
	tmp := i
	for {
		maxIndex := i
		// 是否比左孩子大
		if i * 2 <= n && this.sli[i] < this.sli[i*2] {
			maxIndex = i*2
		}
		// 是否比右孩子大
		if 2*i+1 <= n && this.sli[maxIndex] < this.sli[2*i + 1] {
			maxIndex = 2 * i + 1
		}
		// 不需要交换 退出
		if maxIndex == i {
			break
		}
		// 交换
		this.sli[i], this.sli[maxIndex] = this.sli[maxIndex], this.sli[i]
		//向下遍历
		i = maxIndex
	}

	fmt.Printf("i= %d, arr=%+v \n", tmp, this.sli)
}
// 堆排序
func (this *HeapSort) sort()  {
	//先建立堆
	this.buildHeap()
	k := this.sortNum
	for {
		//堆顶和最后一个交换
		this.sli[1], this.sli[k] = this.sli[k], this.sli[1]
		k--
		// 重新构建堆
		this.heapify(k, 1)
		if k < 1 {
			break
		}
	}

}