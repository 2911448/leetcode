package hot

import "leetcode/common"

func SwapPairs(head *common.ListNode) *common.ListNode {
	dummy := &common.ListNode{Next: head}
	current := dummy
	for current.Next != nil && current.Next.Next != nil {
		one := current.Next
		two := current.Next.Next
		current.Next = two
		one.Next = two.Next
		two.Next = one
		current = one
	}
	return dummy.Next
}
