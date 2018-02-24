package main

type Element interface {}

type ChainNode struct {
	next *ChainNode
	value Element
}

func NewChainNode(value Element) *ChainNode {
	instance := &ChainNode{
		next:nil,
		value:value,
	}
	return instance
}

type Chain struct {
	head	*ChainNode
	len		uint32  // 链表长度
}

func NewChain() *Chain {
	instance := &Chain{}
	node := &ChainNode{
		next:nil,
	}
	instance.head = node
	instance.len = 1
	return instance
}

func (this *Chain) LenAdd() {
	this.len++
}

func (this *Chain) LenDelete() {
	this.len--
}

func (this *Chain) add(value Element) {
	node := NewChainNode(value)
	tail := this.head
	for {
		if this.head.next == nil {
			break
		}
		tail = tail.next
	}
	tail.next = node
	this.LenAdd()
}

// 逻辑不对。。。。。
func (this *Chain) remove(value Element) {
	cur := this.head
	for cur.next != nil {
		if cur.next.value == value {
			cur.next = cur.next.next
		}
		cur = cur.next
	}
}

