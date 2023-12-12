package day

import "leetcode/common"

func AddTwoNumbers(l1 *common.ListNode, l2 *common.ListNode) *common.ListNode {
	head := &common.ListNode{}
	current := head
	carryBit := 0
	for l1 != nil || l2 != nil || carryBit != 0 {
		if l1 != nil {
			carryBit += l1.Val
		}
		if l2 != nil {
			carryBit += l2.Val
		}
		current.Next = &common.ListNode{Val: carryBit % 10}
		carryBit /= 10
		current = current.Next
		if l1 != nil {
			l1 = l1.Next
		}
		if l2 != nil {
			l2 = l2.Next
		}
	}
	return head.Next
}
