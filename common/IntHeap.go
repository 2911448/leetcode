package common

import (
	"container/heap"
)

// IntHeap 是一个实现了 heap.Interface 的结构体
type IntHeap []int

func (h IntHeap) Len() int { return len(h) }

func (h IntHeap) Less(i, j int) bool { return h[i] < h[j] }

func (h IntHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h *IntHeap) Push(x interface{}) {
	*h = append(*h, x.(int))
}

func (h *IntHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

// 初始化一个空的小顶堆
func NewIntHeap() *IntHeap {
	h := &IntHeap{}
	heap.Init(h)
	return h
}

// 向堆中插入一个元素
func (h *IntHeap) PushValue(value int) {
	heap.Push(h, value)
}

// 从堆中弹出并返回最小值
func (h *IntHeap) PopMin() (min int, ok bool) {
	if h.Len() == 0 {
		return 0, false
	}
	min, _ = h.Pop().(int)
	return min, true
}
