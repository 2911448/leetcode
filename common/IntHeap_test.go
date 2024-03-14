package common

import (
	"fmt"
	"testing"
)

func TestIntHeap(T *testing.T) {
	h := NewIntHeap()

	// 插入元素
	h.PushValue(3)
	h.PushValue(1)
	h.PushValue(4)
	h.PushValue(1)
	h.PushValue(5)
	h.PushValue(9)
	h.PushValue(2)

	// 弹出并打印最小值
	for h.Len() > 0 {
		min, ok := h.PopMin()
		if ok {
			fmt.Printf("Popped minimum element: %d\n", min)
		} else {
			break
		}
	}
}
