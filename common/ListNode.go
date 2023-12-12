package common

import (
	"strconv"
	"strings"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

func MakeListNode(num []int) *ListNode {
	if len(num) == 0 {
		return nil
	}

	head := &ListNode{}
	current := head

	for _, v := range num {
		current.Next = &ListNode{
			Val: v,
		}
		current = current.Next
	}

	return head.Next
}

func (ln *ListNode) PrintListNode() string {
	var res strings.Builder
	if ln == nil {
		return res.String()
	}
	current := ln
	for true {
		res.WriteString(strconv.Itoa(current.Val))
		if current.Next == nil {
			break
		}
		current = current.Next
		res.WriteString("->")
	}
	return res.String()
}
