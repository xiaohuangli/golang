package algorithm

import (
	//"sync"
	"fmt"
)

func main() {
	arr := make([]uint32, 0, 0)
	arr = append(arr, 0, 7, 5, 19, 8, 4, 1, 20, 13, 16)
	//buildHeap(arr, 9)
	sort(arr, 9)

	fmt.Println(arr)

}

// 建立堆。从第一个非叶子节点开始
func buildHeap(sli []uint32, n uint32) {
	for i := n/2; i > 0;i-- {
		heapify(sli, n, i)
	}
}

// 向下构建堆
func heapify(sli []uint32, n, i uint32)  {
	tmp := i
	for {
		maxIndex := i
		// 是否比左孩子大
		if i * 2 <= n && sli[i] < sli[i*2] {
			maxIndex = i*2
		}
		// 是否比右孩子大
		if 2*i+1 <= n && sli[maxIndex] < sli[2*i + 1] {
			maxIndex = 2 * i + 1
		}
		// 不需要交换 退出
		if maxIndex == i {
			break
		}
		// 交换
		sli[i], sli[maxIndex] = sli[maxIndex], sli[i]
		//向下遍历
		i = maxIndex
	}

	fmt.Printf("i= %d, arr=%+v \n", tmp, sli)
}
// 堆排序
func sort(sli []uint32, n uint32)  {
	//先建立堆
	buildHeap(sli, n)
	k := n
	for {
		//堆顶和最后一个交换
		sli[1], sli[k] = sli[k], sli[1]
		k--
		// 重新构建堆
		heapify(sli, k, 1)
		if k < 1 {
			break
		}
	}

}
