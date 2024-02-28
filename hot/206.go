package hot

import "leetcode/common"

// 1->2->3->4->5->nil
func ReverseList(head *common.ListNode) *common.ListNode {
	var res *common.ListNode
	current := head
	for current != nil {
		next := current.Next
		current.Next = res
		res = current
		current = next
	}
	return res
}
